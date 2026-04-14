package evaluator

import (
	"strings"
	"testing"

	"pgregory.net/rapid"
)

// Feature: tambah-soal, Property 6: SQL Evaluator menolak operasi berbahaya
// Validates: Requirements 4.4
func TestSQLEvaluatorBlocksDangerousOpsProperty(t *testing.T) {
	svc := NewSQLEvaluatorService()
	dangerousOps := []string{"DROP", "DELETE", "UPDATE", "INSERT", "CREATE", "ALTER", "TRUNCATE", "ATTACH", "DETACH"}
	rapid.Check(t, func(t *rapid.T) {
		op := rapid.SampledFrom(dangerousOps).Draw(t, "op")
		query := op + " TABLE users"
		result := svc.Evaluate("CREATE TABLE users (id INTEGER);", query, "", "[]")
		if result.Passed {
			t.Fatalf("expected blocked op %s to fail, but passed", op)
		}
		if result.Error == "" {
			t.Fatalf("expected error message for blocked op %s", op)
		}
	})
}

// Feature: tambah-soal, Property 7: SQL Evaluator menghasilkan output deterministik
// Validates: Requirements 4.1
func TestSQLEvaluatorDeterministicProperty(t *testing.T) {
	svc := NewSQLEvaluatorService()
	schema := `CREATE TABLE t (id INTEGER, val TEXT);`
	input := `{"tables":{"t":[{"id":1,"val":"a"},{"id":2,"val":"b"}]}}`
	rapid.Check(t, func(t *rapid.T) {
		query := "SELECT * FROM t ORDER BY id"
		expected := `[{"id":1,"val":"a"},{"id":2,"val":"b"}]`
		r1 := svc.Evaluate(schema, query, input, expected)
		r2 := svc.Evaluate(schema, query, input, expected)
		if r1.Actual != r2.Actual {
			t.Fatalf("non-deterministic: %q vs %q", r1.Actual, r2.Actual)
		}
	})
}

// Validates: Requirements 4.1
func TestSQLEvaluatorSelectAll(t *testing.T) {
	svc := NewSQLEvaluatorService()
	schema := `CREATE TABLE users (id INTEGER, name TEXT);`
	input := `{"tables":{"users":[{"id":1,"name":"Alice"}]}}`
	result := svc.Evaluate(schema, "SELECT * FROM users", input, `[{"id":1,"name":"Alice"}]`)
	if !result.Passed {
		t.Fatalf("expected passed=true, got error: %s, actual: %s", result.Error, result.Actual)
	}
}

// Validates: Requirements 4.1 - Empty result set should return [] not null
func TestSQLEvaluatorEmptyResultSet(t *testing.T) {
	svc := NewSQLEvaluatorService()
	schema := `CREATE TABLE users (id INTEGER, name TEXT);`
	input := `{"tables":{"users":[]}}`
	result := svc.Evaluate(schema, "SELECT * FROM users", input, `[]`)
	if !result.Passed {
		t.Fatalf("expected passed=true for empty result, got error: %s, actual: %s", result.Error, result.Actual)
	}
	if result.Actual != "[]" {
		t.Fatalf("expected actual=[], got: %s", result.Actual)
	}
}

// Validates: Requirements 4.4
func TestSQLEvaluatorTimeout(t *testing.T) {
	svc := NewSQLEvaluatorService()
	result := svc.Evaluate("", "WITH RECURSIVE r(n) AS (SELECT 1 UNION ALL SELECT n+1 FROM r) SELECT * FROM r LIMIT 999999999", "", "[]")
	if result.Passed {
		t.Fatal("expected passed=false for timeout")
	}
	if !strings.Contains(result.Error, "Waktu eksekusi habis") {
		t.Fatalf("expected timeout error, got: %s", result.Error)
	}
}

// Validates: Requirements 4.1
func TestSQLEvaluatorSchemaError(t *testing.T) {
	svc := NewSQLEvaluatorService()
	result := svc.Evaluate("THIS IS NOT VALID SQL", "SELECT * FROM users", "", "[]")
	if result.Passed {
		t.Fatal("expected passed=false for schema error")
	}
	if result.Error == "" {
		t.Fatal("expected non-empty error for schema error")
	}
}

// Validates: Requirements 4.1
func TestSQLEvaluatorWrongAnswer(t *testing.T) {
	svc := NewSQLEvaluatorService()
	schema := `CREATE TABLE t (x INTEGER);`
	input := `{"tables":{"t":[{"x":1}]}}`
	result := svc.Evaluate(schema, "SELECT * FROM t", input, `[{"x":999}]`)
	if result.Passed {
		t.Fatal("expected passed=false for wrong answer")
	}
}

// TestSQLEvaluatorConcurrentAccess verifies that concurrent SQL submissions work correctly
// and each gets its own isolated database instance
func TestSQLEvaluatorConcurrentAccess(t *testing.T) {
	svc := NewSQLEvaluatorService()
	schema := `CREATE TABLE t (id INTEGER, val TEXT);`
	input := `{"tables":{"t":[{"id":1,"val":"a"},{"id":2,"val":"b"}]}}`

	// Run multiple concurrent evaluations
	numGoroutines := 10
	results := make(chan EvalResult, numGoroutines)
	errMessages := make(chan string, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			result := svc.Evaluate(schema, "SELECT * FROM t ORDER BY id", input, `[{"id":1,"val":"a"},{"id":2,"val":"b"}]`)
			if !result.Passed {
				errMessages <- result.Error
				return
			}
			results <- result
		}()
	}

	// Collect results
	passCount := 0
	for i := 0; i < numGoroutines; i++ {
		select {
		case errMsg := <-errMessages:
			t.Errorf("concurrent evaluation failed: %s", errMsg)
		case result := <-results:
			if result.Passed {
				passCount++
			}
		}
	}

	if passCount != numGoroutines {
		t.Errorf("expected all %d concurrent evaluations to pass, got %d", numGoroutines, passCount)
	}
}
