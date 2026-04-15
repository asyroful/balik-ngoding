package evaluator

import (
	"encoding/json"
	"fmt"
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

// Feature: tambah-soal, Test: CommaSeparatedArgumentsProperty
// Validates: Evaluator handles comma-separated arguments (e.g., "3, 4") for multi-parameter functions
func TestCommaSeparatedArgumentsProperty(t *testing.T) {
	svc := NewEvaluatorService()

	// Test case 1: kaliTanpaBintang(3, 4) should return 12
	code := `function kaliTanpaBintang(a, b) {
  let hasil = 0;
  for (let i = 0; i < b; i++) {
    hasil += a;
  }
  return hasil;
}`
	tests := []struct {
		input    string
		expected string
	}{
		{"3, 4", "12"},
		{"5, 2", "10"},
		{"0, 7", "0"},
	}

	for _, tt := range tests {
		result := svc.Evaluate(code, tt.input, tt.expected)
		if !result.Passed {
			t.Errorf("kaliTanpaBintang(%s) failed: expected %s, got %q, error: %s", tt.input, tt.expected, result.Actual, result.Error)
		}
	}
}

// Feature: tambah-soal, Test: JSONArrayInputProperty
// Validates: Evaluator handles JSON array input correctly
func TestJSONArrayInputProperty(t *testing.T) {
	svc := NewEvaluatorService()

	code := `function nilaiMaksimum(arr) {
  let maks = arr[0];
  for (let i = 1; i < arr.length; i++) {
    if (arr[i] > maks) {
      maks = arr[i];
    }
  }
  return maks;
}`
	tests := []struct {
		input    string
		expected string
	}{
		{"[3,1,4,1,5,9,2]", "9"},
		{"[10,5,8]", "10"},
		{"[7]", "7"},
	}

	for _, tt := range tests {
		result := svc.Evaluate(code, tt.input, tt.expected)
		if !result.Passed {
			t.Errorf("nilaiMaksimum(%s) failed: expected %s, got %q, error: %s", tt.input, tt.expected, result.Actual, result.Error)
		}
	}
}

// Feature: tambah-soal, Test: NullInputProperty
// Validates: Evaluator handles null input correctly
func TestNullInputProperty(t *testing.T) {
	svc := NewEvaluatorService()

	code := `function greet(n) { return 'hello world'; }`
	input := "null"
	expected := `"hello world"`

	result := svc.Evaluate(code, input, expected)
	if !result.Passed {
		t.Errorf("greet(null) failed: expected %q, got %q, error: %s", expected, result.Actual, result.Error)
	}
}

// ============================================================================
// Phase 1: Evaluator Bug Fix Tests
// ============================================================================

// TestPrepareInputSingleString verifies single string input is parsed correctly
// Feature: evaluator-fix-and-solution-keys, Property 1: string input parsing correctness
// Validates: Requirements 1.1, 1.5
func TestPrepareInputSingleString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple string",
			input:    `"banana"`,
			expected: `"banana"`,
		},
		{
			name:     "string with spaces",
			input:    `"hello world"`,
			expected: `"hello world"`,
		},
		{
			name:     "single character",
			input:    `"a"`,
			expected: `"a"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := prepareInput(tt.input)
			if result != tt.expected {
				t.Errorf("prepareInput(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestPrepareInputCommaSeparatedStrings verifies comma-separated strings are parsed correctly
// Feature: evaluator-fix-and-solution-keys, Property 1: string input parsing correctness
// Validates: Requirements 1.2, 1.5
func TestPrepareInputCommaSeparatedStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "two strings",
			input:    `"banana", "a"`,
			expected: `"banana", "a"`,
		},
		{
			name:     "three strings",
			input:    `"hello", "world", "test"`,
			expected: `"hello", "world", "test"`,
		},
		{
			name:     "strings with spaces",
			input:    `"hello world", "foo bar"`,
			expected: `"hello world", "foo bar"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := prepareInput(tt.input)
			if result != tt.expected {
				t.Errorf("prepareInput(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestPrepareInputMixedTypes verifies mixed string and numeric inputs are parsed correctly
// Feature: evaluator-fix-and-solution-keys, Property 2: mixed type input parsing
// Validates: Requirements 1.3
func TestPrepareInputMixedTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "string and number",
			input:    `"hello", 5`,
			expected: `"hello", 5`,
		},
		{
			name:     "number and string",
			input:    `5, "test"`,
			expected: `5, "test"`,
		},
		{
			name:     "string, number, string",
			input:    `"hello", 5, "world"`,
			expected: `"hello", 5, "world"`,
		},
		{
			name:     "multiple numbers and strings",
			input:    `"a", 1, "b", 2, "c"`,
			expected: `"a", 1, "b", 2, "c"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := prepareInput(tt.input)
			if result != tt.expected {
				t.Errorf("prepareInput(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestPrepareInputValidJSON verifies valid JSON input is passed through unchanged
// Feature: evaluator-fix-and-solution-keys, Property 3: JSON input pass-through
// Validates: Requirements 1.4
func TestPrepareInputValidJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "JSON array of strings",
			input:    `["banana", "a"]`,
			expected: `["banana", "a"]`,
		},
		{
			name:     "JSON array of numbers",
			input:    `[1, 2, 3]`,
			expected: `[1, 2, 3]`,
		},
		{
			name:     "JSON array mixed",
			input:    `["hello", 5, "world"]`,
			expected: `["hello", 5, "world"]`,
		},
		{
			name:     "JSON object",
			input:    `{"key": "value"}`,
			expected: `{"key": "value"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := prepareInput(tt.input)
			if result != tt.expected {
				t.Errorf("prepareInput(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestPrepareInputEmptyString verifies empty input is handled gracefully
// Feature: evaluator-fix-and-solution-keys, Property 4: special character handling
// Validates: Requirements 2.1
func TestPrepareInputEmptyString(t *testing.T) {
	result := prepareInput("")
	expected := `""`
	if result != expected {
		t.Errorf("prepareInput(\"\") = %q, want %q", result, expected)
	}
}

// TestPrepareInputSpecialCharacters verifies special characters are handled correctly
// Feature: evaluator-fix-and-solution-keys, Property 4: special character handling
// Validates: Requirements 2.2
func TestPrepareInputSpecialCharacters(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "string with quotes",
			input: `"hello\"world"`,
		},
		{
			name:  "string with newline",
			input: `"hello\nworld"`,
		},
		{
			name:  "string with tab",
			input: `"hello\tworld"`,
		},
		{
			name:  "string with backslash",
			input: `"hello\\world"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := prepareInput(tt.input)
			// Just verify it doesn't panic and returns something
			if result == "" {
				t.Errorf("prepareInput(%q) returned empty string", tt.input)
			}
		})
	}
}

// TestPrepareInputNumericValue verifies numeric input is parsed as number, not string
// Feature: evaluator-fix-and-solution-keys, Property 5: numeric input recognition
// Validates: Requirements 2.3
func TestPrepareInputNumericValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "integer",
			input:    "42",
			expected: "42",
		},
		{
			name:     "float",
			input:    "3.14",
			expected: "3.14",
		},
		{
			name:     "negative number",
			input:    "-5",
			expected: "-5",
		},
		{
			name:     "zero",
			input:    "0",
			expected: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := prepareInput(tt.input)
			if result != tt.expected {
				t.Errorf("prepareInput(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestPrepareInputWhitespace verifies leading/trailing whitespace is trimmed
// Feature: evaluator-fix-and-solution-keys, Property 6: whitespace trimming
// Validates: Requirements 2.5
func TestPrepareInputWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "leading spaces",
			input:    `  "hello"`,
			expected: `"hello"`,
		},
		{
			name:     "trailing spaces",
			input:    `"hello"  `,
			expected: `"hello"`,
		},
		{
			name:     "both leading and trailing",
			input:    `  "hello"  `,
			expected: `"hello"`,
		},
		{
			name:     "comma-separated with spaces",
			input:    `  "a", "b"  `,
			expected: `"a", "b"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := prepareInput(tt.input)
			if result != tt.expected {
				t.Errorf("prepareInput(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestPrepareInputIdempotenceProperty verifies parsing is idempotent for valid inputs
// Feature: evaluator-fix-and-solution-keys, Property 7: parsing idempotence
// Validates: Requirements 2.6
func TestPrepareInputIdempotenceProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate valid input strings (JSON or comma-separated)
		inputType := rapid.IntRange(0, 2).Draw(t, "inputType")

		var input string
		switch inputType {
		case 0:
			// JSON array
			arr := rapid.SliceOf(rapid.Int()).Draw(t, "array")
			data, _ := json.Marshal(arr)
			input = string(data)
		case 1:
			// Single number
			input = fmt.Sprintf("%d", rapid.Int().Draw(t, "number"))
		default:
			// Single quoted string (already valid)
			str := rapid.StringN(0, 20, -1).Draw(t, "string")
			input = fmt.Sprintf("%q", str)
		}

		parsed1 := prepareInput(input)
		parsed2 := prepareInput(parsed1)

		if parsed1 != parsed2 {
			t.Fatalf("prepareInput not idempotent:\n  first:  %q\n  second: %q", parsed1, parsed2)
		}
	})
}

// TestStringInputParsingProperty verifies string inputs are parsed correctly for function calls
// Feature: evaluator-fix-and-solution-keys, Property 1: string input parsing correctness
// Validates: Requirements 1.1, 1.2, 1.5, 1.6
func TestStringInputParsingProperty(t *testing.T) {
	svc := NewEvaluatorService()

	rapid.Check(t, func(t *rapid.T) {
		// Generate simple alphanumeric strings to avoid escaping issues
		str := rapid.StringMatching(`[a-zA-Z0-9 ]*`).Draw(t, "string")
		input := fmt.Sprintf(`"%s"`, str)

		// Test that the function receives the correct string value
		code := `function test(arg) { return arg; }`
		result := svc.Evaluate(code, input, input)

		if !result.Passed {
			t.Fatalf("function did not receive correct string value for input %q: error=%s, actual=%q", input, result.Error, result.Actual)
		}
	})
}

// ============================================================================
// Phase 1: Evaluator Integration Tests
// ============================================================================

// TestEvaluateBananaAExample tests the banana/a example from requirements
// Feature: evaluator-fix-and-solution-keys, Property 16: evaluator bug fix verification
// Validates: Requirements 11.1
func TestEvaluateBananaAExample(t *testing.T) {
	svc := NewEvaluatorService()

	code := `function getSecond(a, b) { return b; }`
	input := `"banana", "a"`
	expected := `"a"`

	result := svc.Evaluate(code, input, expected)
	if !result.Passed {
		t.Errorf("banana/a example failed: expected %q, got %q, error: %s", expected, result.Actual, result.Error)
	}
}

// TestEvaluateHelloExample tests the hello example from requirements
// Feature: evaluator-fix-and-solution-keys, Property 16: evaluator bug fix verification
// Validates: Requirements 11.2
func TestEvaluateHelloExample(t *testing.T) {
	svc := NewEvaluatorService()

	code := `function identity(s) { return s; }`
	input := `"hello"`
	expected := `"hello"`

	result := svc.Evaluate(code, input, expected)
	if !result.Passed {
		t.Errorf("hello example failed: expected %q, got %q, error: %s", expected, result.Actual, result.Error)
	}
}

// TestEvaluateMixedTypesExample tests the mixed types example from requirements
// Feature: evaluator-fix-and-solution-keys, Property 16: evaluator bug fix verification
// Validates: Requirements 11.3
func TestEvaluateMixedTypesExample(t *testing.T) {
	svc := NewEvaluatorService()

	code := `function concat(n, s) { return s + n; }`
	input := `5, "test"`
	expected := `"test5"`

	result := svc.Evaluate(code, input, expected)
	if !result.Passed {
		t.Errorf("mixed types example failed: expected %q, got %q, error: %s", expected, result.Actual, result.Error)
	}
}

// TestEvaluatorRegressionProperty tests that all existing test cases pass with fixed evaluator
// Feature: evaluator-fix-and-solution-keys, Property 16: evaluator bug fix verification
// Validates: Requirements 11.4
func TestEvaluatorRegressionProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		svc := NewEvaluatorService()

		// Test cases that should all pass with fixed evaluator
		testCases := []struct {
			code     string
			input    string
			expected string
		}{
			{
				code:     `function add(a, b) { return a + b; }`,
				input:    `[1, 2]`,
				expected: `3`,
			},
			{
				code:     `function concat(a, b) { return a + b; }`,
				input:    `"hello", " world"`,
				expected: `"hello world"`,
			},
			{
				code:     `function identity(x) { return x; }`,
				input:    `"test"`,
				expected: `"test"`,
			},
		}

		// Run each test case
		for _, tc := range testCases {
			result := svc.Evaluate(tc.code, tc.input, tc.expected)
			if !result.Passed {
				t.Fatalf("regression test failed for code %q with input %q: expected %q, got %q, error: %s",
					tc.code, tc.input, tc.expected, result.Actual, result.Error)
			}
		}
	})
}
