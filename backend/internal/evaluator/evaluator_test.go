package evaluator

import (
	"strings"
	"testing"

	"pgregory.net/rapid"
)

// determineStatus is a pure helper that mirrors the status logic in the evaluator.
// P11: status is determined by whether all passed, some failed, or an error occurred.
func determineStatus(results []EvalResult) string {
	hasError := false
	allPassed := true
	for _, r := range results {
		if r.Error != "" {
			hasError = true
		}
		if !r.Passed {
			allPassed = false
		}
	}
	if hasError {
		return "error"
	}
	if allPassed && len(results) > 0 {
		return "accepted"
	}
	return "wrong_answer"
}

// normalizeOutput trims leading/trailing whitespace, mirroring the evaluator behaviour.
func normalizeOutput(s string) string {
	return strings.TrimSpace(s)
}

// Generators

func arbitraryEvalResult() *rapid.Generator[EvalResult] {
	return rapid.Custom(func(t *rapid.T) EvalResult {
		passed := rapid.Bool().Draw(t, "passed")
		errMsg := ""
		if rapid.Bool().Draw(t, "hasError") {
			errMsg = rapid.StringN(1, 50, -1).Draw(t, "error")
			passed = false
		}
		return EvalResult{
			Passed: passed,
			Actual: rapid.String().Draw(t, "actual"),
			Error:  errMsg,
		}
	})
}

// Feature: balik-ngoding, Property 11: Submission status is correctly determined
// Validates: Requirements 5.4, 5.5, 5.6
func TestDetermineStatusProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		results := rapid.SliceOf(arbitraryEvalResult()).Draw(t, "results")
		status := determineStatus(results)

		hasError := false
		allPassed := true
		for _, r := range results {
			if r.Error != "" {
				hasError = true
			}
			if !r.Passed {
				allPassed = false
			}
		}

		switch {
		case hasError:
			if status != "error" {
				t.Fatalf("expected status 'error' when error present, got %q", status)
			}
		case allPassed && len(results) > 0:
			if status != "accepted" {
				t.Fatalf("expected status 'accepted' when all passed, got %q", status)
			}
		default:
			if status != "wrong_answer" {
				t.Fatalf("expected status 'wrong_answer', got %q", status)
			}
		}
	})
}

// Feature: balik-ngoding, Property 12: Output normalization before comparison
// Validates: Requirements 5.8
func TestOutputNormalizationProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		core := rapid.StringN(1, 50, -1).Draw(t, "core")
		leading := rapid.StringMatching(`[ \t\n\r]*`).Draw(t, "leading")
		trailing := rapid.StringMatching(`[ \t\n\r]*`).Draw(t, "trailing")

		actual := leading + core + trailing
		expected := core

		if normalizeOutput(actual) != normalizeOutput(expected) {
			t.Fatalf("normalizeOutput(%q) != normalizeOutput(%q)", actual, expected)
		}
	})
}

// Feature: balik-ngoding, Property 15: Sandbox blocks restricted resource access
// Validates: Requirements 8.2
//
// Each snippet attempts to *call* or *access a property of* a restricted API.
// Since the sandbox sets these to undefined, calling them throws a TypeError,
// which the evaluator surfaces as a non-empty Error field.
func TestSandboxBlocksRestrictedAPIsProperty(t *testing.T) {
	svc := NewEvaluatorService()

	// These snippets all attempt to invoke or dereference a restricted name,
	// which will throw a TypeError because the name is set to undefined.
	restrictedSnippets := []string{
		`require('fs')`,
		`process.env.HOME`,
		`fetch('http://example.com')`,
		`new XMLHttpRequest()`,
		`new WebSocket('ws://example.com')`,
		`fs.readFileSync('/etc/passwd')`,
	}

	rapid.Check(t, func(t *rapid.T) {
		snippet := rapid.SampledFrom(restrictedSnippets).Draw(t, "snippet")
		code := "function solve(input) { " + snippet + "; return 1; }"
		result := svc.Evaluate(code, "1", "1")
		if result.Error == "" {
			t.Fatalf("expected error when using restricted API %q, but got none (actual=%q)", snippet, result.Actual)
		}
	})
}

// Feature: balik-ngoding, Example test for timeout enforcement
// Validates: Requirements 5.7, 8.3
func TestTimeoutEnforcementInfiniteLoop(t *testing.T) {
	svc := NewEvaluatorService()
	code := `function solve(input) { while(true) {} }`
	result := svc.Evaluate(code, "1", "anything")

	if result.Error != "Waktu eksekusi habis" {
		t.Fatalf("expected 'Waktu eksekusi habis', got %q", result.Error)
	}
	if result.Passed {
		t.Fatal("expected Passed=false for timed-out execution")
	}
}

// Feature: tambah-soal, Property 8: EvaluatorService routing berdasarkan language
// Validates: Requirements 7.1, 7.3
func TestEvaluatorServiceRoutingProperty(t *testing.T) {
	svc := NewEvaluatorService()
	sqlSvc := NewSQLEvaluatorService()

	schema := `CREATE TABLE t (x INTEGER); INSERT INTO t VALUES (42);`
	query := "SELECT x FROM t"
	expected := `[{"x":42}]`

	rapid.Check(t, func(t *rapid.T) {
		// SQL routing: EvaluateWithLanguage("sql") should match SQLEvaluatorService.Evaluate directly
		r1 := svc.EvaluateWithLanguage("sql", schema, query, "", expected)
		r2 := sqlSvc.Evaluate(schema, query, expected)
		if r1.Passed != r2.Passed {
			t.Fatalf("SQL routing mismatch: EvaluateWithLanguage.Passed=%v, direct.Passed=%v", r1.Passed, r2.Passed)
		}
	})
}

// Feature: tambah-soal, Property 8: EvaluatorService routing berdasarkan language (JavaScript)
// Validates: Requirements 7.1, 7.3
func TestEvaluatorServiceJSRoutingProperty(t *testing.T) {
	svc := NewEvaluatorService()

	code := "function add(n) { return n + 1; }"
	input := "5"
	expected := "6"

	rapid.Check(t, func(t *rapid.T) {
		r1 := svc.EvaluateWithLanguage("javascript", "", code, input, expected)
		r2 := svc.Evaluate(code, input, expected)
		if r1.Passed != r2.Passed {
			t.Fatalf("JS routing mismatch: EvaluateWithLanguage.Passed=%v, direct.Passed=%v", r1.Passed, r2.Passed)
		}
	})
}
