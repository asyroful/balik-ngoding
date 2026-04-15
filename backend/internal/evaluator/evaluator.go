package evaluator

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dop251/goja"
)

// normalizeOutput normalizes output for comparison by:
// 1. Trimming leading/trailing whitespace
// 2. Replacing multiple spaces with single space
// 3. Normalizing newlines (\r\n to \n)
func normalizeOutput(s string) string {
	// Trim leading/trailing whitespace
	s = strings.TrimSpace(s)

	// Replace multiple spaces with single space
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")

	// Normalize newlines (convert \r\n to \n)
	s = strings.ReplaceAll(s, "\r\n", "\n")

	return s
}

// EvalResult holds the result of a single test case evaluation.
type EvalResult struct {
	Passed bool   `json:"passed"`
	Actual string `json:"actual"`
	Error  string `json:"error,omitempty"`
}

// EvaluatorService executes user-submitted JavaScript code in a sandboxed goja runtime.
type EvaluatorService struct {
	sqlEvaluator *SQLEvaluatorService
}

// NewEvaluatorService creates a new EvaluatorService.
func NewEvaluatorService() *EvaluatorService {
	return &EvaluatorService{
		sqlEvaluator: NewSQLEvaluatorService(),
	}
}

// EvaluateWithLanguage routes evaluation to the appropriate evaluator based on language.
// For "sql", it delegates to the SQL evaluator with input data.
// For "javascript" or empty string, it delegates to the JS evaluator.
// Unknown languages fall back to the JS evaluator with a warning log.
func (e *EvaluatorService) EvaluateWithLanguage(language, schema, code, input, expected string) EvalResult {
	switch language {
	case "sql":
		return e.sqlEvaluator.Evaluate(schema, code, input, expected)
	case "javascript", "":
		return e.Evaluate(code, input, expected)
	default:
		log.Printf("Warning: unknown language %q, falling back to JavaScript evaluator", language)
		return e.Evaluate(code, input, expected)
	}
}

// Evaluate runs the user's code with the given input and returns an EvalResult.
// It enforces a 5-second timeout, blocks restricted APIs, and normalizes output.
func (e *EvaluatorService) Evaluate(code string, input string, expected string) EvalResult {
	vm := goja.New()

	// Block restricted APIs (task 4.3)
	restricted := []string{
		"require", "process", "fs", "http", "https",
		"fetch", "XMLHttpRequest", "WebSocket",
	}
	for _, name := range restricted {
		vm.Set(name, goja.Undefined())
	}

	// Capture console.log output (task 4.1)
	var outputBuf strings.Builder
	console := vm.NewObject()
	console.Set("log", func(call goja.FunctionCall) goja.Value {
		parts := make([]string, len(call.Arguments))
		for i, arg := range call.Arguments {
			parts[i] = fmt.Sprintf("%v", arg.Export())
		}
		outputBuf.WriteString(strings.Join(parts, " "))
		outputBuf.WriteString("\n")
		return goja.Undefined()
	})
	vm.Set("console", console)

	// Extract function name from user code using regex (task 4.1)
	funcName := extractFunctionName(code)

	// Build the full script to execute
	var script string
	if funcName != "" {
		// Prepare input - try JSON parse, fallback to comma-separated
		inputValue := prepareInput(input)

		script = fmt.Sprintf(`
%s
(function() {
  try {
    var __result = %s(%s);
    console.log(JSON.stringify(__result));
  } catch(e) {
    throw e;
  }
})();
`, code, funcName, inputValue)
	} else {
		// Fallback: inject input as a variable and run code as-is
		script = fmt.Sprintf(`
var input = JSON.parse(%s);
%s
`, jsonStringLiteral(input), code)
	}

	// Run with timeout using goroutine + interrupt (task 4.2)
	type result struct {
		val goja.Value
		err error
	}
	ch := make(chan result, 1)

	timer := time.AfterFunc(5*time.Second, func() {
		vm.Interrupt("Waktu eksekusi habis")
	})

	go func() {
		val, err := vm.RunString(script)
		ch <- result{val, err}
	}()

	res := <-ch
	timer.Stop()

	// Handle errors (task 4.5)
	if res.err != nil {
		errMsg := res.err.Error()
		// Check if it's a timeout interrupt
		if strings.Contains(errMsg, "Waktu eksekusi habis") {
			return EvalResult{Passed: false, Error: "Waktu eksekusi habis"}
		}
		// Unwrap goja exception for a cleaner message
		if jsErr, ok := res.err.(*goja.Exception); ok {
			errMsg = jsErr.Error()
		}
		return EvalResult{Passed: false, Error: fmt.Sprintf("Runtime error: %s", errMsg)}
	}

	// Normalize actual output (task 4.4)
	actual := strings.TrimSpace(outputBuf.String())

	// Compare with expected using normalized output
	passed := normalizeOutput(actual) == normalizeOutput(expected)

	return EvalResult{
		Passed: passed,
		Actual: actual,
	}
}

// extractFunctionName uses a regex to find the first function declaration name in the code.
func extractFunctionName(code string) string {
	re := regexp.MustCompile(`(?m)^\s*function\s+([a-zA-Z_$][a-zA-Z0-9_$]*)`)
	matches := re.FindStringSubmatch(code)
	if len(matches) >= 2 {
		return matches[1]
	}
	return ""
}

// prepareInput converts input string to JavaScript code that can be passed to a function.
// If input is valid JSON, it's used as-is. Otherwise, it's treated as comma-separated arguments.
// For each argument:
// - If it's a valid JSON string (quoted), use as-is
// - If it's a number, use as-is
// - Otherwise, wrap as a string literal
func prepareInput(input string) string {
	trimmed := strings.TrimSpace(input)
	if len(trimmed) == 0 {
		return `""`
	}

	// Try to parse as JSON first
	var jsonValue interface{}
	if err := json.Unmarshal([]byte(trimmed), &jsonValue); err == nil {
		// Valid JSON, use as-is
		return trimmed
	}

	// Not valid JSON, treat as comma-separated arguments
	// Split by comma and process each part
	parts := strings.Split(trimmed, ",")
	var args []string
	for _, part := range parts {
		trimmedPart := strings.TrimSpace(part)

		// Check if it's already a valid JSON string (starts and ends with quotes)
		if len(trimmedPart) >= 2 && trimmedPart[0] == '"' && trimmedPart[len(trimmedPart)-1] == '"' {
			// Try to parse as JSON string to validate it
			var strValue string
			if err := json.Unmarshal([]byte(trimmedPart), &strValue); err == nil {
				// Valid JSON string, use as-is
				args = append(args, trimmedPart)
				continue
			}
		}

		// Try to parse as number
		if _, err := strconv.ParseFloat(trimmedPart, 64); err == nil {
			args = append(args, trimmedPart)
		} else {
			// Not a number or JSON string, wrap as string literal
			args = append(args, fmt.Sprintf("%q", trimmedPart))
		}
	}
	return strings.Join(args, ", ")
}

// jsonStringLiteral wraps a raw string value in a JSON string literal for safe embedding in JS.
// If the input is valid JSON, it's used as-is. Otherwise, it's wrapped as a JSON string.
func jsonStringLiteral(input string) string {
	trimmed := strings.TrimSpace(input)
	if len(trimmed) == 0 {
		return `""`
	}

	// Try to parse as JSON first
	var jsonValue interface{}
	if err := json.Unmarshal([]byte(trimmed), &jsonValue); err == nil {
		// Valid JSON, use as-is
		return trimmed
	}

	// Not valid JSON, wrap as string
	return fmt.Sprintf("%q", trimmed)
}
