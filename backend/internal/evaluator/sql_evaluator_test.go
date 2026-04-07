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
		result := svc.Evaluate("CREATE TABLE users (id INTEGER);", query, "[]")
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
	schema := `CREATE TABLE t (id INTEGER, val TEXT);
INSERT INTO t VALUES (1, 'a');
INSERT INTO t VALUES (2, 'b');`
	rapid.Check(t, func(t *rapid.T) {
		query := "SELECT * FROM t ORDER BY id"
		expected := `[{"id":1,"val":"a"},{"id":2,"val":"b"}]`
		r1 := svc.Evaluate(schema, query, expected)
		r2 := svc.Evaluate(schema, query, expected)
		if r1.Actual != r2.Actual {
			t.Fatalf("non-deterministic: %q vs %q", r1.Actual, r2.Actual)
		}
	})
}

// Validates: Requirements 4.1
func TestSQLEvaluatorSelectAll(t *testing.T) {
	svc := NewSQLEvaluatorService()
	schema := `CREATE TABLE users (id INTEGER, name TEXT);
INSERT INTO users VALUES (1, 'Alice');`
	result := svc.Evaluate(schema, "SELECT * FROM users", `[{"id":1,"name":"Alice"}]`)
	if !result.Passed {
		t.Fatalf("expected passed=true, got error: %s, actual: %s", result.Error, result.Actual)
	}
}

// Validates: Requirements 4.4
func TestSQLEvaluatorTimeout(t *testing.T) {
	svc := NewSQLEvaluatorService()
	result := svc.Evaluate("", "WITH RECURSIVE r(n) AS (SELECT 1 UNION ALL SELECT n+1 FROM r) SELECT * FROM r LIMIT 999999999", "[]")
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
	result := svc.Evaluate("THIS IS NOT VALID SQL", "SELECT * FROM users", "[]")
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
	schema := `CREATE TABLE t (x INTEGER); INSERT INTO t VALUES (1);`
	result := svc.Evaluate(schema, "SELECT * FROM t", `[{"x":999}]`)
	if result.Passed {
		t.Fatal("expected passed=false for wrong answer")
	}
}
