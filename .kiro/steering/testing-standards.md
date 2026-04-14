# Testing Standards — Balik Ngoding

## Testing Philosophy

- **Test-Driven Development**: Write tests before implementation when possible
- **Property-Based Testing**: Use fast-check dan rapid untuk generate test cases
- **Comprehensive Coverage**: Aim for >80% code coverage
- **Real-World Scenarios**: Test actual user workflows, not just isolated functions
- **Performance Testing**: Benchmark critical paths

## Unit Testing

### Backend (Go)

**Test File Organization:**
- Test files in same package as code being tested
- File naming: `{package}_test.go`
- Function naming: `Test{FunctionName}(t *testing.T)`

**Table-Driven Tests:**
```go
func TestNormalizeOutput(t *testing.T) {
  tests := []struct {
    name     string
    input    string
    expected string
  }{
    {
      name:     "trim leading/trailing spaces",
      input:    "  hello world  ",
      expected: "hello world",
    },
    {
      name:     "normalize multiple spaces",
      input:    "hello    world",
      expected: "hello world",
    },
    {
      name:     "normalize newlines",
      input:    "hello\r\nworld",
      expected: "hello\nworld",
    },
  }
  
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      result := normalizeOutput(tt.input)
      if result != tt.expected {
        t.Errorf("got %q, want %q", result, tt.expected)
      }
    })
  }
}
```

**Error Testing:**
```go
func TestEvaluateWithError(t *testing.T) {
  tests := []struct {
    name        string
    code        string
    expectError bool
    errorMsg    string
  }{
    {
      name:        "syntax error",
      code:        "function broken {",
      expectError: true,
      errorMsg:    "SyntaxError",
    },
  }
  
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      result := evaluator.Evaluate(tt.code, "[]", "")
      if (result.Error != "") != tt.expectError {
        t.Errorf("expectError %v, got error %q", tt.expectError, result.Error)
      }
    })
  }
}
```

### Frontend (TypeScript)

**Test File Organization:**
- Test files in `__tests__/` directory
- File naming: `{component}.test.tsx` atau `{function}.test.ts`
- Test function naming: `test('description', () => {})`

**Component Testing:**
```typescript
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { CodeEditor } from '@/components/Editor/CodeEditor';

describe('CodeEditor', () => {
  test('renders code input field', () => {
    render(<CodeEditor problemId="1" />);
    expect(screen.getByRole('textbox')).toBeInTheDocument();
  });
  
  test('calls onChange when code changes', async () => {
    const onChange = vi.fn();
    render(<CodeEditor problemId="1" onChange={onChange} />);
    
    const input = screen.getByRole('textbox');
    await userEvent.type(input, 'console.log("test")');
    
    expect(onChange).toHaveBeenCalled();
  });
  
  test('displays error message on submission failure', async () => {
    vi.mock('@/lib/api', () => ({
      submitCode: vi.fn().mockRejectedValue(new Error('Network error')),
    }));
    
    render(<CodeEditor problemId="1" />);
    const submitButton = screen.getByRole('button', { name: /submit/i });
    
    await userEvent.click(submitButton);
    
    expect(screen.getByText(/network error/i)).toBeInTheDocument();
  });
});
```

## Property-Based Testing

### Backend (Go with rapid)

**Normalization Idempotence:**
```go
import "github.com/flyingmutant/rapid"

func TestNormalizationIdempotenceProperty(t *testing.T) {
  rapid.Check(t, func(t *rapid.T) {
    output := rapid.String().Draw(t, "output")
    
    // Property: normalizing twice should give same result as normalizing once
    normalized1 := normalizeOutput(output)
    normalized2 := normalizeOutput(normalized1)
    
    if normalized1 != normalized2 {
      t.Fatalf("normalization not idempotent:\n  first:  %q\n  second: %q", normalized1, normalized2)
    }
  })
}
```

**Whitespace Tolerance:**
```go
func TestWhitespaceToleranceProperty(t *testing.T) {
  rapid.Check(t, func(t *rapid.T) {
    baseOutput := rapid.String().Draw(t, "output")
    
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
        t.Fatalf("whitespace variations not normalized equally")
      }
    }
  })
}
```

### Frontend (TypeScript with fast-check)

**Normalization Idempotence:**
```typescript
import fc from 'fast-check';

test('normalization is idempotent', () => {
  fc.assert(
    fc.property(fc.string(), (output) => {
      const normalized1 = normalizeOutput(output);
      const normalized2 = normalizeOutput(normalized1);
      return normalized1 === normalized2;
    })
  );
});
```

**Whitespace Tolerance:**
```typescript
test('whitespace variations normalize equally', () => {
  fc.assert(
    fc.property(fc.string(), (baseOutput) => {
      const normalized = normalizeOutput(baseOutput);
      
      // Test variations
      const withLeadingSpace = normalizeOutput('  ' + baseOutput);
      const withTrailingSpace = normalizeOutput(baseOutput + '  ');
      const withMultipleSpaces = normalizeOutput(baseOutput.replace(/ /g, '  '));
      
      return (
        withLeadingSpace === normalized &&
        withTrailingSpace === normalized &&
        withMultipleSpaces === normalized
      );
    })
  );
});
```

## Integration Testing

### Backend

**End-to-End Submission Flow:**
```go
func TestSubmissionFlow(t *testing.T) {
  // Setup
  evaluator := NewEvaluatorService()
  
  // Test case 1: Correct answer with whitespace variation
  result := evaluator.Evaluate(
    "function add(a, b) { return a + b; }",
    "[1, 2]",
    "  3  ", // Expected with extra spaces
  )
  
  if !result.Passed {
    t.Errorf("expected passed, got error: %s", result.Error)
  }
  
  // Test case 2: Wrong answer should still fail
  result = evaluator.Evaluate(
    "function add(a, b) { return a + b; }",
    "[1, 2]",
    "5", // Wrong expected value
  )
  
  if result.Passed {
    t.Errorf("expected failed, but passed")
  }
}
```

### Frontend

**End-to-End Problem Solving:**
```typescript
test('user can solve a problem end-to-end', async () => {
  // Mock API
  vi.mock('@/lib/api', () => ({
    fetchProblem: vi.fn().mockResolvedValue({
      id: '1',
      title: 'Add Numbers',
      description: 'Write a function that adds two numbers',
      starterCode: 'function add(a, b) {\n  // Your code here\n}',
    }),
    submitCode: vi.fn().mockResolvedValue({
      status: 'accepted',
      score: 1,
      total: 1,
      results: [{ passed: true, input: '[1, 2]', expected: '3', actual: '3' }],
    }),
  }));
  
  // Render page
  render(<ProblemPage params={{ id: '1' }} />);
  
  // Wait for problem to load
  await screen.findByText('Add Numbers');
  
  // Edit code
  const editor = screen.getByRole('textbox');
  await userEvent.type(editor, 'return a + b;');
  
  // Submit
  const submitButton = screen.getByRole('button', { name: /submit/i });
  await userEvent.click(submitButton);
  
  // Verify result
  await screen.findByText(/accepted/i);
  expect(screen.getByText('1 / 1')).toBeInTheDocument();
});
```

## Performance Testing

### Benchmarking

**Backend Benchmarks:**
```go
func BenchmarkNormalizeOutput(b *testing.B) {
  input := "  hello    world  \n\n"
  
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    normalizeOutput(input)
  }
}

func BenchmarkEvaluate(b *testing.B) {
  code := "function add(a, b) { return a + b; }"
  input := "[1, 2]"
  expected := "3"
  
  evaluator := NewEvaluatorService()
  
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    evaluator.Evaluate(code, input, expected)
  }
}
```

**Frontend Performance:**
```typescript
test('CodeEditor renders within acceptable time', () => {
  const startTime = performance.now();
  
  render(<CodeEditor problemId="1" />);
  
  const endTime = performance.now();
  const renderTime = endTime - startTime;
  
  // Should render in less than 100ms
  expect(renderTime).toBeLessThan(100);
});
```

## Test Coverage

### Coverage Goals

- **Backend**: Minimum 80% coverage
- **Frontend**: Minimum 75% coverage
- **Critical Paths**: 100% coverage (evaluator, submissions, storage)

### Coverage Reporting

**Backend:**
```bash
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

**Frontend:**
```bash
vitest run --coverage
```

## Continuous Integration

### Test Execution

**On Pull Request:**
1. Run all unit tests
2. Run all integration tests
3. Run property-based tests
4. Generate coverage report
5. Check coverage thresholds

**On Merge to Main:**
1. Run full test suite
2. Run performance benchmarks
3. Deploy to staging
4. Run smoke tests

### Test Failure Handling

- Failing tests block merge
- Coverage drops block merge
- Performance regressions require approval
- Flaky tests must be fixed or skipped with explanation

## Test Data Management

### Fixtures

**Backend:**
```go
func setupTestDB(t *testing.T) *gorm.DB {
  db := setupTestDatabase()
  
  // Create test problems
  problems := []models.Problem{
    {
      ID:       "test-1",
      Title:    "Add Numbers",
      Category: "loop",
    },
  }
  
  for _, p := range problems {
    db.Create(&p)
  }
  
  t.Cleanup(func() {
    db.Migrator().DropTable(&models.Problem{})
  })
  
  return db
}
```

**Frontend:**
```typescript
const mockProblems = [
  {
    id: '1',
    title: 'Add Numbers',
    description: 'Write a function that adds two numbers',
    category: 'loop',
    difficulty: 'easy',
  },
];

const mockSubmissionResult = {
  status: 'accepted',
  score: 1,
  total: 1,
  results: [
    {
      passed: true,
      input: '[1, 2]',
      expected: '3',
      actual: '3',
    },
  ],
};
```

## Debugging Tests

### Logging

**Backend:**
```go
t.Logf("Input: %q", input)
t.Logf("Expected: %q", expected)
t.Logf("Actual: %q", actual)
```

**Frontend:**
```typescript
console.log('Props:', props);
console.log('State:', state);
screen.debug(); // Print DOM
```

### Running Specific Tests

**Backend:**
```bash
go test -run TestNormalizeOutput ./...
go test -run TestNormalizeOutput/trim ./...
```

**Frontend:**
```bash
vitest run CodeEditor.test.tsx
vitest run -t "renders code input field"
```

