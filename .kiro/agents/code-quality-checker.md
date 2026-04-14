# Code Quality Checker Agent — Balik Ngoding

## Purpose

Automated code quality verification untuk memastikan semua code mengikuti standards dan best practices sebelum merge.

## Responsibilities

1. **Linting**: Check code style dan potential issues
2. **Type Safety**: Verify type annotations dan type safety
3. **Test Coverage**: Ensure adequate test coverage
4. **Documentation**: Verify code documentation
5. **Security**: Check untuk security vulnerabilities

## Execution Triggers

- On Pull Request creation
- On commit push to feature branch
- Manual trigger via CLI

## Go Code Quality Checks

### 1. Linting dengan golangci-lint

**Configuration:**
```yaml
# .golangci.yml
linters:
  enable:
    - errcheck          # Check unchecked errors
    - govet             # Vet issues
    - ineffassign       # Ineffectual assignments
    - staticcheck       # Static analysis
    - unused            # Unused code
    - misspell          # Misspelled words
    - gofmt             # Code formatting
    - goimports         # Import organization

issues:
  exclude-rules:
    - path: _test\.go$
      linters:
        - errcheck
```

**Execution:**
```bash
golangci-lint run ./...
```

**Checks:**
- [ ] No unused variables or imports
- [ ] All errors are checked
- [ ] No ineffectual assignments
- [ ] Code is properly formatted
- [ ] Imports are organized

### 2. Type Safety

**Checks:**
- [ ] All function parameters have type annotations
- [ ] All function return types are specified
- [ ] No use of `interface{}`
- [ ] Proper error handling with typed errors

**Validation:**
```bash
go vet ./...
```

### 3. Test Coverage

**Minimum Coverage:**
- Backend: 80% overall, 100% for critical paths
- Critical paths: evaluator, submissions, storage

**Execution:**
```bash
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

**Report:**
- [ ] Coverage >= 80%
- [ ] Critical paths >= 100%
- [ ] No coverage regressions

### 4. Documentation

**Checks:**
- [ ] All exported functions have comments
- [ ] Comments explain "why", not "what"
- [ ] Package-level comments present
- [ ] Complex logic has inline comments

**Validation:**
```bash
go doc ./...
```

### 5. Security Checks

**Checks:**
- [ ] No hardcoded secrets
- [ ] Proper input validation
- [ ] SQL injection prevention (parameterized queries)
- [ ] No unsafe operations

**Tools:**
```bash
gosec ./...
```

## TypeScript Code Quality Checks

### 1. Linting dengan ESLint

**Configuration:**
```json
{
  "extends": ["next/core-web-vitals"],
  "rules": {
    "@typescript-eslint/no-explicit-any": "error",
    "@typescript-eslint/explicit-function-return-types": "error",
    "@typescript-eslint/no-unused-vars": "error",
    "no-console": ["warn", { "allow": ["warn", "error"] }]
  }
}
```

**Execution:**
```bash
eslint . --ext .ts,.tsx
```

**Checks:**
- [ ] No `any` types
- [ ] All functions have return types
- [ ] No unused variables
- [ ] No console.log in production code

### 2. Type Safety dengan TypeScript

**Checks:**
- [ ] No type errors
- [ ] Strict mode enabled
- [ ] All types properly defined
- [ ] No implicit `any`

**Execution:**
```bash
tsc --noEmit
```

### 3. Test Coverage

**Minimum Coverage:**
- Frontend: 75% overall, 100% for critical paths
- Critical paths: CodeEditor, ResultPanel, API client

**Execution:**
```bash
vitest run --coverage
```

**Report:**
- [ ] Coverage >= 75%
- [ ] Critical paths >= 100%
- [ ] No coverage regressions

### 4. Documentation

**Checks:**
- [ ] All exported functions have JSDoc
- [ ] Components have prop documentation
- [ ] Complex logic has inline comments
- [ ] README updated if needed

**Validation:**
```bash
# Manual check for JSDoc comments
grep -r "^/\*\*" src/
```

### 5. Security Checks

**Checks:**
- [ ] No hardcoded secrets
- [ ] Proper input sanitization
- [ ] XSS prevention
- [ ] CSRF protection

**Tools:**
```bash
npm audit
```

## Accessibility Checks

### WCAG Compliance

**Checks:**
- [ ] All images have alt text
- [ ] All buttons have accessible labels
- [ ] Color contrast meets WCAG AA
- [ ] Keyboard navigation works
- [ ] Screen reader compatible

**Tools:**
```bash
axe-core
lighthouse
```

## Performance Checks

### Frontend Performance

**Metrics:**
- [ ] Largest Contentful Paint (LCP) < 2.5s
- [ ] First Input Delay (FID) < 100ms
- [ ] Cumulative Layout Shift (CLS) < 0.1
- [ ] Bundle size < 500KB

**Tools:**
```bash
lighthouse
next/bundle-analyzer
```

### Backend Performance

**Metrics:**
- [ ] API response time < 200ms
- [ ] Database query time < 100ms
- [ ] Memory usage < 100MB
- [ ] CPU usage < 50%

**Tools:**
```bash
pprof
benchstat
```

## Report Generation

### Quality Report

**Output Format:**
```
Code Quality Report
===================

Go Code:
  ✓ Linting: PASS
  ✓ Type Safety: PASS
  ✓ Test Coverage: 85% (target: 80%)
  ✓ Documentation: PASS
  ✓ Security: PASS

TypeScript Code:
  ✓ Linting: PASS
  ✓ Type Safety: PASS
  ✓ Test Coverage: 78% (target: 75%)
  ✓ Documentation: PASS
  ✓ Security: PASS

Accessibility:
  ✓ WCAG AA: PASS

Performance:
  ✓ Frontend: PASS (LCP: 1.8s)
  ✓ Backend: PASS (API: 150ms)

Overall: PASS ✓
```

### Failure Handling

**If Checks Fail:**
1. Generate detailed report
2. List specific issues
3. Provide remediation steps
4. Block merge until fixed

**Example:**
```
FAILED: Code Quality Check

Issues Found:
1. Go Linting (3 issues):
   - backend/internal/evaluator/evaluator.go:45: unused variable 'x'
   - backend/internal/submissions/service.go:12: error not checked

2. TypeScript Linting (2 issues):
   - frontend/components/Editor/CodeEditor.tsx:8: no-explicit-any
   - frontend/lib/api.ts:15: missing return type

3. Test Coverage (1 issue):
   - Coverage dropped from 85% to 82%

Remediation:
- Run: go fmt ./...
- Run: eslint --fix .
- Add missing tests
```

## Integration with CI/CD

### GitHub Actions Workflow

```yaml
name: Code Quality Check

on: [pull_request, push]

jobs:
  quality:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.25'
      
      - name: Setup Node
        uses: actions/setup-node@v3
        with:
          node-version: '18'
      
      - name: Go Linting
        run: |
          go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
          golangci-lint run ./...
      
      - name: Go Tests
        run: go test -cover ./...
      
      - name: TypeScript Linting
        run: |
          cd frontend
          npm install
          npm run lint
      
      - name: TypeScript Tests
        run: |
          cd frontend
          npm run test:coverage
      
      - name: Security Scan
        run: |
          go install github.com/securego/gosec/v2/cmd/gosec@latest
          gosec ./...
          cd frontend && npm audit
```

## Manual Execution

### Run All Checks

```bash
# Backend
cd backend
go fmt ./...
golangci-lint run ./...
go test -cover ./...
go vet ./...
gosec ./...

# Frontend
cd frontend
npm run lint
npm run lint:fix
npm run test:coverage
npm audit
```

### Run Specific Check

```bash
# Go linting only
golangci-lint run ./...

# TypeScript linting only
cd frontend && npm run lint

# Test coverage only
go test -cover ./...
cd frontend && npm run test:coverage
```

## Configuration Files

### .golangci.yml
- Linting rules untuk Go
- Excluded patterns
- Severity levels

### .eslintrc.json
- Linting rules untuk TypeScript
- Plugin configuration
- Override rules

### .prettierrc
- Code formatting rules
- Line length
- Indentation

### tsconfig.json
- TypeScript strict mode
- Target version
- Module resolution

## Success Criteria

- [ ] All linting checks pass
- [ ] All type checks pass
- [ ] Test coverage meets minimum
- [ ] No security vulnerabilities
- [ ] Documentation complete
- [ ] Performance metrics acceptable
- [ ] Accessibility standards met

