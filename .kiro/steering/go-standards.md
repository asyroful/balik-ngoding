# Go Coding Standards — Balik Ngoding

## Naming Conventions

### Packages
- Use lowercase, single-word package names
- Examples: `evaluator`, `submissions`, `database`, `models`
- Avoid generic names like `util`, `helper`, `common`

### Functions & Methods
- Use camelCase for exported functions
- Use camelCase for unexported functions
- Use descriptive names that indicate action/purpose
- Examples: `NewEvaluatorService()`, `Evaluate()`, `extractFunctionName()`

### Variables & Constants
- Use camelCase for variables
- Use UPPER_SNAKE_CASE for constants
- Use descriptive names
- Examples: `evaluator`, `testCases`, `MAX_TIMEOUT`

### Interfaces
- Use descriptive names ending with "er" or "Service"
- Examples: `EvaluatorService`, `SubmissionsService`

## Error Handling

### Error Checking
- Always check errors explicitly
- Use `if err != nil` pattern
- Never ignore errors silently
- Wrap errors dengan context when appropriate

```go
// Good
result := database.DB.First(&problem, "id = ?", req.ProblemID)
if result.Error != nil {
  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    return nil // caller will 404
  }
  return nil, result.Error
}

// Bad
result := database.DB.First(&problem, "id = ?", req.ProblemID)
// Ignoring error
```

### Error Messages
- Use clear, actionable error messages
- Include context about what failed
- Use Indonesian untuk user-facing messages
- Use English untuk internal/debug messages

```go
// Good
return EvalResult{Passed: false, Error: "Waktu eksekusi habis"}

// Bad
return EvalResult{Passed: false, Error: "Error"}
```

## Testing Patterns

### Unit Tests
- Test file naming: `{package}_test.go`
- Test function naming: `Test{FunctionName}(t *testing.T)`
- Use table-driven tests untuk multiple scenarios
- Test both success dan failure cases

```go
func TestEvaluate(t *testing.T) {
  tests := []struct {
    name     string
    code     string
    input    string
    expected string
    want     bool
  }{
    {
      name:     "simple function",
      code:     "function add(a, b) { return a + b; }",
      input:    "[1, 2]",
      expected: "3",
      want:     true,
    },
  }
  
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      // Test implementation
    })
  }
}
```

### Property-Based Tests
- Use `github.com/flyingmutant/rapid` untuk property-based testing
- Test invariants dan properties
- Generate random test cases

```go
func TestNormalizationProperty(t *testing.T) {
  rapid.Check(t, func(t *rapid.T) {
    output := rapid.String().Draw(t, "output")
    
    // Property: idempotence
    normalized1 := normalizeOutput(output)
    normalized2 := normalizeOutput(normalized1)
    
    if normalized1 != normalized2 {
      t.Fatalf("normalization not idempotent: %q != %q", normalized1, normalized2)
    }
  })
}
```

## Code Organization

### Package Structure
- Keep packages focused dan single-responsibility
- Avoid circular dependencies
- Use `internal/` untuk non-public packages
- Organize by feature/domain, not by type

```
backend/internal/
├── database/      # Database connection & operations
├── evaluator/     # Code evaluation logic
├── models/        # Data models
├── problems/      # Problem-related handlers & services
└── submissions/   # Submission-related handlers & services
```

### File Organization
- One main type per file (or closely related types)
- Keep files under 500 lines
- Group related functions together
- Use clear file names

## Documentation

### Comments
- Write comments untuk "why", not "what"
- Use clear, concise language
- Update comments when code changes
- Use package-level comments untuk exported packages

```go
// Good - explains why
// extractFunctionName uses regex to find the first function declaration
// because we need to call it with the parsed input

// Bad - explains what (obvious from code)
// Extract the function name
```

### Function Documentation
- Document exported functions dengan comment starting with function name
- Include parameter descriptions
- Include return value descriptions
- Include error conditions

```go
// Evaluate runs the user's code with the given input and returns an EvalResult.
// It enforces a 5-second timeout, blocks restricted APIs, and normalizes output.
func (e *EvaluatorService) Evaluate(code string, input string, expected string) EvalResult {
  // Implementation
}
```

## Performance Considerations

### Database Queries
- Use indexes untuk frequently queried columns
- Avoid N+1 queries
- Use connection pooling
- Batch operations when possible

### Memory Management
- Avoid unnecessary allocations
- Use buffers untuk string concatenation
- Clean up resources dengan defer
- Monitor goroutine leaks

### Concurrency
- Use channels untuk communication
- Avoid shared mutable state
- Use mutexes untuk protecting shared data
- Test concurrent scenarios

## Security Best Practices

### Input Validation
- Validate all user input
- Use parameterized queries untuk database
- Sanitize output untuk prevent injection
- Check file paths untuk prevent traversal

### Secrets Management
- Use environment variables untuk secrets
- Never commit secrets ke repository
- Use `.env` files locally (not in git)
- Rotate secrets regularly

### Error Messages
- Don't expose internal details di error messages
- Don't expose database structure
- Don't expose file paths
- Log detailed errors internally

