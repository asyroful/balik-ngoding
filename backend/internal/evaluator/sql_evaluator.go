package evaluator

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"regexp"
	"time"
)

var blockedOps = regexp.MustCompile(`(?i)\b(DROP|DELETE|UPDATE|INSERT|CREATE|ALTER|TRUNCATE|ATTACH|DETACH)\b`)

// SQLEvaluatorService evaluates SQL SELECT queries against an in-memory SQLite database.
type SQLEvaluatorService struct{}

// NewSQLEvaluatorService creates a new SQLEvaluatorService.
func NewSQLEvaluatorService() *SQLEvaluatorService {
	return &SQLEvaluatorService{}
}

// Evaluate runs the user's SQL query against the given schema and compares the result to expected JSON.
func (s *SQLEvaluatorService) Evaluate(schema, query, expected string) EvalResult {
	// Reject blocked operations
	if blockedOps.MatchString(query) {
		return EvalResult{Passed: false, Error: "Operasi tidak diizinkan: hanya SELECT yang diperbolehkan"}
	}

	// Unique DSN per call to avoid shared state
	dsn := fmt.Sprintf("file:memdb%d?mode=memory&cache=shared", time.Now().UnixNano())
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return EvalResult{Passed: false, Error: fmt.Sprintf("Gagal membuka database: %s", err.Error())}
	}
	defer db.Close()

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
	var result []map[string]interface{}
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
