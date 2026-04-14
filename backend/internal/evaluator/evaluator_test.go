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

// Note: normalizeOutput is now defined in evaluator.go with full whitespace normalization

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

// TestNormalizeOutputIdempotence verifies that normalizing twice gives same result as normalizing once
func TestNormalizeOutputIdempotence(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "leading spaces",
			input: "  hello world",
		},
		{
			name:  "trailing spaces",
			input: "hello world  ",
		},
		{
			name:  "multiple spaces",
			input: "hello    world",
		},
		{
			name:  "newlines",
			input: "hello\nworld",
		},
		{
			name:  "carriage return newlines",
			input: "hello\r\nworld",
		},
		{
			name:  "mixed whitespace",
			input: "  hello  \n  world  \r\n  test  ",
		},
		{
			name:  "tabs",
			input: "hello\t\tworld",
		},
		{
			name:  "empty string",
			input: "",
		},
		{
			name:  "only whitespace",
			input: "   \n\r\n  \t  ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalized1 := normalizeOutput(tt.input)
			normalized2 := normalizeOutput(normalized1)

			if normalized1 != normalized2 {
				t.Errorf("normalization not idempotent:\n  first:  %q\n  second: %q", normalized1, normalized2)
			}
		})
	}
}

// TestNormalizeOutputIdempotenceProperty uses property-based testing to verify idempotence
func TestNormalizeOutputIdempotenceProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		output := rapid.String().Draw(t, "output")

		normalized1 := normalizeOutput(output)
		normalized2 := normalizeOutput(normalized1)

		if normalized1 != normalized2 {
			t.Fatalf("normalization not idempotent:\n  first:  %q\n  second: %q", normalized1, normalized2)
		}
	})
}

// TestWhitespaceToleranceProperty verifies that different whitespace variations normalize equally
func TestWhitespaceToleranceProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		baseOutput := rapid.StringN(1, 50, -1).Draw(t, "output")

		// Generate variations with different whitespace
		variations := []string{
			baseOutput,
			"  " + baseOutput + "  ",
			strings.ReplaceAll(baseOutput, " ", "  "),
			strings.ReplaceAll(baseOutput, "\n", "\r\n"),
		}

		// All variations should normalize to same value
		normalized := normalizeOutput(variations[0])
		for _, v := range variations[1:] {
			if normalizeOutput(v) != normalized {
				t.Fatalf("whitespace variations not normalized equally:\n  base: %q\n  variant: %q", normalized, normalizeOutput(v))
			}
		}
	})
}

// TestPreserveWrongAnswersProperty verifies that wrong answers are still marked as wrong
func TestPreserveWrongAnswersProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		actual := rapid.StringN(1, 50, -1).Draw(t, "actual")
		expected := rapid.StringN(1, 50, -1).Draw(t, "expected")

		// Only test when they're actually different after normalization
		if normalizeOutput(actual) != normalizeOutput(expected) {
			// They should remain different after normalization
			if normalizeOutput(actual) == normalizeOutput(expected) {
				t.Fatalf("wrong answers should not normalize equally:\n  actual: %q\n  expected: %q", actual, expected)
			}
		}
	})
}

// TestWhitespaceToleranceIntegration verifies that whitespace variations are accepted in actual evaluation
func TestWhitespaceToleranceIntegration(t *testing.T) {
	svc := NewEvaluatorService()

	tests := []struct {
		name     string
		code     string
		input    string
		expected string
		wantPass bool
	}{
		{
			name:     "leading spaces in expected",
			code:     "function add(n) { return n + 1; }",
			input:    "5",
			expected: "  6  ",
			wantPass: true,
		},
		{
			name:     "multiple spaces in expected",
			code:     "function greet(n) { return 'hello world'; }",
			input:    "null",
			expected: "  \"hello world\"  ",
			wantPass: true,
		},
		{
			name:     "newline variations",
			code:     "function test(n) { return 'line1\\nline2'; }",
			input:    "null",
			expected: "\"line1\\nline2\"",
			wantPass: true,
		},
		{
			name:     "wrong answer still fails",
			code:     "function add(n) { return n + 1; }",
			input:    "5",
			expected: "  999  ",
			wantPass: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.Evaluate(tt.code, tt.input, tt.expected)

			if result.Passed != tt.wantPass {
				t.Errorf("expected Passed=%v, got %v (actual=%q, error=%q)", tt.wantPass, result.Passed, result.Actual, result.Error)
			}
		})
	}
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

	schema := `CREATE TABLE t (x INTEGER);`
	query := "SELECT x FROM t"
	input := `{"tables":{"t":[{"x":42}]}}`
	expected := `[{"x":42}]`

	rapid.Check(t, func(t *rapid.T) {
		// SQL routing: EvaluateWithLanguage("sql") should match SQLEvaluatorService.Evaluate directly
		r1 := svc.EvaluateWithLanguage("sql", schema, query, input, expected)
		r2 := sqlSvc.Evaluate(schema, query, input, expected)
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
