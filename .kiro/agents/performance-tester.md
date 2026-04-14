# Performance Tester Agent — Balik Ngoding

## Purpose

Automated performance testing dan benchmarking untuk memastikan aplikasi memenuhi performance targets dan mendeteksi regressions.

## Responsibilities

1. **Benchmarking**: Run performance benchmarks untuk critical paths
2. **Load Testing**: Test aplikasi dengan concurrent load
3. **Memory Profiling**: Monitor memory usage dan detect leaks
4. **Database Performance**: Monitor query performance
5. **Frontend Performance**: Measure Core Web Vitals

## Execution Triggers

- On Pull Request creation
- On merge to main
- Scheduled daily
- Manual trigger via CLI

## Backend Performance Testing

### 1. Benchmarking dengan Go

**Critical Paths untuk Benchmark:**
- `normalizeOutput()` — Output normalization
- `Evaluate()` — Code evaluation
- `SQLEvaluator.Evaluate()` — SQL evaluation
- `Submit()` — Submission processing

**Benchmark Implementation:**
```go
// backend/internal/evaluator/evaluator_bench_test.go

func BenchmarkNormalizeOutput(b *testing.B) {
  inputs := []string{
    "  hello world  ",
    "hello    world",
    "hello\r\nworld",
    strings.Repeat("x", 10000),
  }
  
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    normalizeOutput(inputs[i%len(inputs)])
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

func BenchmarkSQLEvaluate(b *testing.B) {
  schema := `
    CREATE TABLE users (id INT, name TEXT);
    INSERT INTO users VALUES (1, 'Alice'), (2, 'Bob');
  `
  query := "SELECT * FROM users WHERE id = 1"
  expected := `[{"id":1,"name":"Alice"}]`
  
  evaluator := NewSQLEvaluatorService()
  
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    evaluator.Evaluate(schema, query, expected)
  }
}

func BenchmarkSubmit(b *testing.B) {
  // Setup test database
  db := setupTestDB()
  service := NewSubmissionsService()
  
  req := SubmitRequest{
    ProblemID: "test-1",
    Code:      "function add(a, b) { return a + b; }",
    Language:  "javascript",
  }
  
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    service.Submit(req)
  }
}
```

**Execution:**
```bash
go test -bench=. -benchmem ./...
go test -bench=. -benchmem -benchtime=10s ./...
```

**Output Analysis:**
```
BenchmarkNormalizeOutput-8        1000000    1234 ns/op    256 B/op    2 allocs/op
BenchmarkEvaluate-8                 10000   123456 ns/op   5120 B/op   45 allocs/op
BenchmarkSQLEvaluate-8               1000  1234567 ns/op  51200 B/op  234 allocs/op
BenchmarkSubmit-8                     100 12345678 ns/op 512000 B/op 1234 allocs/op
```

**Performance Targets:**
- `normalizeOutput`: < 2µs per operation
- `Evaluate`: < 200ms per operation
- `SQLEvaluate`: < 500ms per operation
- `Submit`: < 1s per operation

### 2. Load Testing dengan Apache Bench

**Setup:**
```bash
# Install Apache Bench
apt-get install apache2-utils

# Or use wrk for more advanced testing
curl -L https://github.com/wg/wrk/releases/download/4.2.0/wrk-linux-x86_64.tar.gz | tar xz
```

**Load Test Scenarios:**

**Scenario 1: Normal Load**
```bash
ab -n 1000 -c 10 http://localhost:8080/problems
```

**Scenario 2: High Concurrency**
```bash
ab -n 10000 -c 100 http://localhost:8080/problems
```

**Scenario 3: Submission Load**
```bash
# Create test script
cat > submit_load.lua << 'EOF'
request = function()
  wrk.method = "POST"
  wrk.body = '{"problemId":"1","code":"function add(a,b){return a+b;}","language":"javascript"}'
  wrk.headers["Content-Type"] = "application/json"
  return wrk.format(nil, "/submit")
end
EOF

wrk -t4 -c100 -d30s -s submit_load.lua http://localhost:8080
```

**Performance Targets:**
- Requests per second: > 1000 RPS
- Average latency: < 100ms
- P95 latency: < 500ms
- P99 latency: < 1s
- Error rate: < 0.1%

### 3. Memory Profiling

**CPU Profiling:**
```go
// backend/main.go
import _ "net/http/pprof"

func init() {
  go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
  }()
}
```

**Collect Profile:**
```bash
# CPU profile (30 seconds)
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# Memory profile
go tool pprof http://localhost:6060/debug/pprof/heap

# Goroutine profile
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

**Analysis:**
```bash
# Top functions by CPU time
(pprof) top

# Memory allocations
(pprof) alloc_space

# Goroutine leaks
(pprof) goroutine
```

**Memory Targets:**
- Baseline memory: < 50MB
- Memory per request: < 1MB
- No goroutine leaks
- GC pause time: < 100ms

### 4. Database Query Performance

**Query Profiling:**
```sql
-- Enable query logging
SET log_statement = 'all';
SET log_duration = on;

-- Analyze slow queries
EXPLAIN ANALYZE SELECT * FROM problems WHERE category = 'sql';
```

**Performance Targets:**
- Simple queries: < 10ms
- Complex queries: < 100ms
- Index usage: 100% for indexed columns

## Frontend Performance Testing

### 1. Core Web Vitals

**Metrics:**
- **LCP (Largest Contentful Paint)**: < 2.5s
- **FID (First Input Delay)**: < 100ms
- **CLS (Cumulative Layout Shift)**: < 0.1

**Measurement dengan Lighthouse:**
```bash
npm install -g lighthouse

# Run Lighthouse
lighthouse http://localhost:3000 --view

# Headless mode
lighthouse http://localhost:3000 --output=json > report.json
```

**Measurement dengan Web Vitals:**
```typescript
// frontend/lib/web-vitals.ts
import { getCLS, getFID, getFCP, getLCP, getTTFB } from 'web-vitals';

export function reportWebVitals() {
  getCLS(console.log);
  getFID(console.log);
  getFCP(console.log);
  getLCP(console.log);
  getTTFB(console.log);
}
```

### 2. Bundle Size Analysis

**Analyze Bundle:**
```bash
npm install -g webpack-bundle-analyzer

# Build dan analyze
npm run build
npx webpack-bundle-analyzer .next/static/chunks/main-*.js
```

**Bundle Size Targets:**
- Main bundle: < 200KB
- Total JS: < 500KB
- CSS: < 100KB

### 3. Component Render Performance

**Benchmark Components:**
```typescript
// frontend/__tests__/performance.test.tsx
import { render } from '@testing-library/react';
import { CodeEditor } from '@/components/Editor/CodeEditor';

test('CodeEditor renders within acceptable time', () => {
  const startTime = performance.now();
  
  render(<CodeEditor problemId="1" />);
  
  const endTime = performance.now();
  const renderTime = endTime - startTime;
  
  // Should render in less than 100ms
  expect(renderTime).toBeLessThan(100);
});

test('ResultPanel renders within acceptable time', () => {
  const startTime = performance.now();
  
  render(<ResultPanel result={mockResult} />);
  
  const endTime = performance.now();
  const renderTime = endTime - startTime;
  
  expect(renderTime).toBeLessThan(50);
});
```

**Render Performance Targets:**
- Initial render: < 100ms
- Re-render: < 50ms
- Large lists: < 200ms

### 4. Network Performance

**Measure API Response Time:**
```typescript
// frontend/lib/api.ts
async function measureAPICall(url: string, options: RequestInit) {
  const startTime = performance.now();
  
  const response = await fetch(url, options);
  
  const endTime = performance.now();
  const duration = endTime - startTime;
  
  console.log(`API call to ${url} took ${duration}ms`);
  
  return response;
}
```

**API Response Time Targets:**
- GET /problems: < 100ms
- GET /problems/:id: < 100ms
- POST /submit: < 500ms

## Performance Report

### Report Format

```
Performance Test Report
=======================

Backend Benchmarks:
  normalizeOutput:  1234 ns/op (target: 2000 ns/op) ✓
  Evaluate:        123456 ns/op (target: 200ms) ✓
  SQLEvaluate:    1234567 ns/op (target: 500ms) ✓
  Submit:        12345678 ns/op (target: 1s) ✓

Load Testing:
  Requests/sec:     1500 RPS (target: 1000) ✓
  Avg Latency:       75ms (target: 100ms) ✓
  P95 Latency:      350ms (target: 500ms) ✓
  P99 Latency:      800ms (target: 1s) ✓
  Error Rate:      0.05% (target: 0.1%) ✓

Memory Profiling:
  Baseline:         45MB (target: 50MB) ✓
  Per Request:     0.8MB (target: 1MB) ✓
  Goroutines:        12 (no leaks) ✓
  GC Pause:         45ms (target: 100ms) ✓

Frontend Performance:
  LCP:             1.8s (target: 2.5s) ✓
  FID:              45ms (target: 100ms) ✓
  CLS:             0.08 (target: 0.1) ✓
  Bundle Size:     420KB (target: 500KB) ✓

Overall: PASS ✓
```

### Regression Detection

**If Performance Degrades:**
1. Identify which metric regressed
2. Compare with previous benchmark
3. Identify code changes that caused regression
4. Require performance improvement before merge

**Example:**
```
PERFORMANCE REGRESSION DETECTED

Metric: Evaluate() benchmark
Previous: 123456 ns/op
Current:  234567 ns/op
Change:   +90% (FAIL)

Likely cause: Commit abc123 - Added extra validation

Action Required:
- Optimize the validation logic
- Or revert the change
- Re-run benchmark to verify fix
```

## Integration with CI/CD

### GitHub Actions Workflow

```yaml
name: Performance Test

on:
  pull_request:
  push:
    branches: [main]
  schedule:
    - cron: '0 2 * * *'  # Daily at 2 AM

jobs:
  performance:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.25'
      
      - name: Run Benchmarks
        run: |
          cd backend
          go test -bench=. -benchmem -benchtime=10s ./... | tee benchmark.txt
      
      - name: Compare Benchmarks
        uses: benchmark-action/github-action-benchmark@v1
        with:
          tool: 'go'
          output-file-path: backend/benchmark.txt
          github-token: ${{ secrets.GITHUB_TOKEN }}
          auto-push: true
      
      - name: Load Testing
        run: |
          # Start backend
          cd backend && go run main.go &
          sleep 2
          
          # Run load test
          ab -n 1000 -c 10 http://localhost:8080/problems
      
      - name: Frontend Performance
        run: |
          cd frontend
          npm install
          npm run build
          npm run lighthouse
```

## Manual Execution

### Run All Performance Tests

```bash
# Backend benchmarks
cd backend
go test -bench=. -benchmem ./...

# Load testing
ab -n 1000 -c 10 http://localhost:8080/problems

# Memory profiling
go tool pprof http://localhost:6060/debug/pprof/heap

# Frontend performance
cd frontend
npm run lighthouse
npm run bundle-analyze
```

### Run Specific Test

```bash
# Specific benchmark
go test -bench=BenchmarkEvaluate -benchmem ./...

# Load test specific endpoint
ab -n 100 -c 5 http://localhost:8080/problems/1

# Memory profile
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

## Success Criteria

- [ ] All benchmarks meet targets
- [ ] Load test passes with acceptable latency
- [ ] No memory leaks detected
- [ ] No goroutine leaks
- [ ] Core Web Vitals meet targets
- [ ] Bundle size within limits
- [ ] No performance regressions
- [ ] Database queries optimized

