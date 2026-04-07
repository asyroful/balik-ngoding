package evaluator

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/dop251/goja"
)

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
// For "sql", it delegates to the SQL evaluator (input is ignored).
// For "javascript" or empty string, it delegates to the JS evaluator.
// Unknown languages fall back to the JS evaluator with a warning log.
func (e *EvaluatorService) EvaluateWithLanguage(language, schema, code, input, expected string) EvalResult {
	switch language {
	case "sql":
		return e.sqlEvaluator.Evaluate(schema, code, expected)
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
		// Call the user's function with the parsed input
		script = fmt.Sprintf(`
%s
(function() {
  try {
    var __input = JSON.parse(%s);
    var __result = %s(__input);
    console.log(JSON.stringify(__result));
  } catch(e) {
    throw e;
  }
})();
`, code, jsonStringLiteral(input), funcName)
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

	// Compare with expected
	passed := actual == strings.TrimSpace(expected)

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

// jsonStringLiteral wraps a raw string value in a JSON string literal for safe embedding in JS.
// If the input is already a valid JSON value (object, array, number, bool, null), it is used as-is.
func jsonStringLiteral(input string) string {
	trimmed := strings.TrimSpace(input)
	if len(trimmed) == 0 {
		return `""`
	}
	// If it looks like a JSON value (starts with {, [, digit, -, t, f, n, or is a quoted string), use as-is
	first := trimmed[0]
	if first == '{' || first == '[' || first == '"' ||
		(first >= '0' && first <= '9') || first == '-' ||
		trimmed == "true" || trimmed == "false" || trimmed == "null" {
		return trimmed
	}
	// Otherwise wrap as a JSON string
	return fmt.Sprintf("%q", trimmed)
}
