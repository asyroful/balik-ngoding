package evaluator

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

var blockedOps = regexp.MustCompile(`(?i)\b(DROP|DELETE|UPDATE|INSERT|CREATE|ALTER|TRUNCATE|ATTACH|DETACH)\b`)

// SQLEvaluatorService evaluates SQL SELECT queries against an in-memory SQLite database.
type SQLEvaluatorService struct{}

// NewSQLEvaluatorService creates a new SQLEvaluatorService.
func NewSQLEvaluatorService() *SQLEvaluatorService {
	return &SQLEvaluatorService{}
}

// populateTablesFromInput parses the input JSON to extract table data
// and populates the database with INSERT statements.
// The input format is: {"tables": {"tableName": [{"col1": val1, ...}, ...]}}
// This allows each test case to have different data.
func (s *SQLEvaluatorService) populateTablesFromInput(tx *sql.Tx, input string) error {
	// Parse input to get table data
	var inputData struct {
		Tables map[string][]map[string]interface{} `json:"tables"`
	}
	if err := json.Unmarshal([]byte(input), &inputData); err != nil {
		// If input is not in the expected format, skip population
		return nil
	}

	// For each table in the input, delete existing data and insert new data
	for tableName, rows := range inputData.Tables {
		// Delete existing data from the table
		if _, err := tx.Exec(fmt.Sprintf("DELETE FROM %s", tableName)); err != nil {
			return fmt.Errorf("failed to delete from %s: %w", tableName, err)
		}

		// Insert new data
		for _, row := range rows {
			var columns []string
			var values []interface{}
			for col, val := range row {
				columns = append(columns, col)
				values = append(values, val)
			}

			// Build INSERT statement
			placeholders := make([]string, len(columns))
			for i := range columns {
				placeholders[i] = "?"
			}
			insertSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
				tableName,
				strings.Join(columns, ", "),
				strings.Join(placeholders, ", "),
			)

			if _, err := tx.Exec(insertSQL, values...); err != nil {
				return fmt.Errorf("failed to insert into %s: %w", tableName, err)
			}
		}
	}

	return nil
}

// Evaluate runs the user's SQL query against the given schema and compares the result to expected JSON.
// The input parameter contains the test data in format: {"tables": {"tableName": [...]}}
func (s *SQLEvaluatorService) Evaluate(schema, query, input, expected string) EvalResult {
	// Reject blocked operations
	if blockedOps.MatchString(query) {
		return EvalResult{Passed: false, Error: "Operasi tidak diizinkan: hanya SELECT yang diperbolehkan"}
	}

	// Unique DSN per call using UUID to guarantee uniqueness and avoid shared state
	dsn := fmt.Sprintf("file:memdb_%s?mode=memory&cache=private", uuid.New().String())
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return EvalResult{Passed: false, Error: fmt.Sprintf("Gagal membuka database: %s", err.Error())}
	}
	defer db.Close()

	// Configure connection limits for in-memory database
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	// Run schema in a transaction
	tx, err := db.Begin()
	if err != nil {
		return EvalResult{Passed: false, Error: fmt.Sprintf("Gagal memulai transaksi: %s", err.Error())}
	}
	if _, err := tx.Exec(schema); err != nil {
		tx.Rollback()
		return EvalResult{Passed: false, Error: fmt.Sprintf("Error pada schema: %s", err.Error())}
	}

	if err := tx.Commit(); err != nil {
		return EvalResult{Passed: false, Error: fmt.Sprintf("Gagal commit schema: %s", err.Error())}
	}

	// Populate tables from input data
	tx2, err := db.Begin()
	if err != nil {
		return EvalResult{Passed: false, Error: fmt.Sprintf("Gagal memulai transaksi: %s", err.Error())}
	}
	if err := s.populateTablesFromInput(tx2, input); err != nil {
		tx2.Rollback()
		return EvalResult{Passed: false, Error: fmt.Sprintf("Error pada input: %s", err.Error())}
	}
	if err := tx2.Commit(); err != nil {
		return EvalResult{Passed: false, Error: fmt.Sprintf("Gagal commit data: %s", err.Error())}
	}

	// Run user query with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return EvalResult{Passed: false, Error: "Waktu eksekusi habis"}
		}
		return EvalResult{Passed: false, Error: fmt.Sprintf("Error pada query: %s", err.Error())}
	}
	defer rows.Close()

	// Read column names
	cols, err := rows.Columns()
	if err != nil {
		return EvalResult{Passed: false, Error: fmt.Sprintf("Gagal membaca kolom: %s", err.Error())}
	}

	// Serialize rows to []map[string]interface{}
	// Initialize as empty slice (not nil) so json.Marshal returns [] instead of null
	result := make([]map[string]interface{}, 0)
	for rows.Next() {
		values := make([]interface{}, len(cols))
		valuePtrs := make([]interface{}, len(cols))
		for i := range values {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			if ctx.Err() == context.DeadlineExceeded {
				return EvalResult{Passed: false, Error: "Waktu eksekusi habis"}
			}
			return EvalResult{Passed: false, Error: fmt.Sprintf("Gagal membaca baris: %s", err.Error())}
		}
		row := make(map[string]interface{})
		for i, col := range cols {
			row[col] = values[i]
		}
		result = append(result, row)
	}
	if err := rows.Err(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return EvalResult{Passed: false, Error: "Waktu eksekusi habis"}
		}
		return EvalResult{Passed: false, Error: fmt.Sprintf("Error iterasi baris: %s", err.Error())}
	}

	// Marshal actual result to JSON
	actualBytes, err := json.Marshal(result)
	if err != nil {
		return EvalResult{Passed: false, Error: fmt.Sprintf("Gagal serialisasi hasil: %s", err.Error())}
	}
	actualStr := string(actualBytes)

	// Normalize expected JSON
	var expectedParsed interface{}
	if err := json.Unmarshal([]byte(expected), &expectedParsed); err != nil {
		return EvalResult{Passed: false, Actual: actualStr, Error: fmt.Sprintf("Expected output bukan JSON valid: %s", err.Error())}
	}
	normalizedExpected, err := json.Marshal(expectedParsed)
	if err != nil {
		return EvalResult{Passed: false, Actual: actualStr, Error: fmt.Sprintf("Gagal normalisasi expected: %s", err.Error())}
	}

	// Normalize actual JSON (unmarshal → marshal)
	var actualParsed interface{}
	if err := json.Unmarshal(actualBytes, &actualParsed); err != nil {
		return EvalResult{Passed: false, Actual: actualStr, Error: fmt.Sprintf("Gagal normalisasi actual: %s", err.Error())}
	}
	normalizedActual, err := json.Marshal(actualParsed)
	if err != nil {
		return EvalResult{Passed: false, Actual: actualStr, Error: fmt.Sprintf("Gagal normalisasi actual: %s", err.Error())}
	}

	passed := string(normalizedActual) == string(normalizedExpected)
	return EvalResult{
		Passed: passed,
		Actual: actualStr,
	}
}
