# Design Document — Balik Ngoding Comprehensive Bugfix

## Overview

Dokumen ini mendefinisikan strategi teknis untuk mengatasi tiga bugs kritis di Balik Ngoding dengan fokus pada implementation details, testing strategy, dan automation setup.

---

## Bug #1: Strict Answer Validation — Design

### Technical Context

**Current Implementation:**
- File: `backend/internal/evaluator/evaluator.go`
- Fungsi: `Evaluate()` melakukan string comparison langsung: `actual == strings.TrimSpace(expected)`
- Hanya melakukan `TrimSpace()` pada expected output, tidak pada actual output
- Tidak menangani multiple spaces, newlines, atau whitespace variations

**Root Cause:**
```go
// Current (line ~130)
passed := actual == strings.TrimSpace(expected)
```

Perbandingan ini terlalu strict karena:
1. Hanya trim expected, tidak actual
2. Tidak normalize multiple spaces
3. Tidak handle newline variations

### Solution Design

**Normalization Strategy:**
```go
func normalizeOutput(s string) string {
  // 1. Trim leading/trailing whitespace
  s = strings.TrimSpace(s)
  
  // 2. Replace multiple spaces dengan single space
  s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
  
  // 3. Normalize newlines (convert \r\n to \n)
  s = strings.ReplaceAll(s, "\r\n", "\n")
  
  return s
}
```

**Implementation Changes:**
1. Update `evaluator.go` untuk menggunakan `normalizeOutput()` pada both actual dan expected
2. Ensure normalization terjadi sebelum comparison
3. Add property-based tests untuk validate normalization behavior

**Files to Modify:**
- `backend/internal/evaluator/evaluator.go` — Update comparison logic
- `backend/internal/evaluator/evaluator_test.go` — Add normalization tests

### Correctness Properties

**Property 1: Whitespace Tolerance**
```
FOR ALL (actual, expected) WHERE normalize(actual) == normalize(expected)
  ASSERT evaluator'(actual, expected).passed == true
```

**Property 2: Preserve Wrong Answers**
```
FOR ALL (actual, expected) WHERE normalize(actual) != normalize(expected)
  ASSERT evaluator'(actual, expected).passed == false
```

**Property 3: Idempotent Normalization**
```
FOR ALL output
  ASSERT normalize(normalize(output)) == normalize(output)
```

---

## Bug #2: Code Storage Best Practice — Design

### Technical Context

**Current Implementation:**
- Problem code stored in PostgreSQL: `description` (text), `starterCode` (text), `schema` (text)
- User submissions stored in PostgreSQL: `code` (text)
- Models: `backend/internal/models/models.go`

**Current Schema:**
```go
type Problem struct {
  Description string          // Large text blob
  StarterCode string          // Large text blob
  Schema      string          // Large text blob (for SQL)
}

type Submission struct {
  Code string                 // Large text blob
}
```

**Performance Concerns:**
1. Large text blobs dalam database → slower queries
2. Backup/restore operations lebih lambat
3. Indexing pada text columns tidak efisien
4. Replication overhead untuk large blobs

### Solution Design

**Recommendation: Hybrid Approach**

**For Problem Code (Description, StarterCode, Schema):**
- **Keep in Database** (PostgreSQL)
- Reasoning:
  - Problem code tidak sering berubah (write-once, read-many)
  - Perlu atomicity dengan problem metadata
  - Queries untuk problem detail perlu semua data sekaligus
  - Easier backup/restore strategy
- **Optimization:**
  - Add database indexing pada `problem_id`
  - Use connection pooling untuk reduce overhead
  - Consider caching layer (Redis) untuk frequently accessed problems

**For User Submission Code:**
- **Move to File System** (dengan database reference)
- Reasoning:
  - Submissions grow rapidly (write-heavy)
  - Tidak perlu atomicity dengan submission metadata
  - Can be archived/cleaned up independently
  - Reduces database size dan improves query performance
- **Implementation:**
  - Store code di file system: `/submissions/{submission_id}.{language}`
  - Store file path reference di database
  - Implement cleanup strategy untuk old submissions

**Storage Structure:**
```
submissions/
├── {submission_id}.js
├── {submission_id}.sql
└── ...
```

**Database Schema Update:**
```go
type Submission struct {
  ID           string          // UUID
  ProblemID    string          // UUID
  CodePath     string          // File path reference
  Language     string          // javascript, sql, etc
  Status       string          // accepted, wrong_answer, error
  // ... other fields
}
```

**Implementation Plan:**
1. Create file storage service: `backend/internal/storage/file_storage.go`
2. Update Submission model untuk use CodePath instead of Code
3. Update submissions handler untuk save code ke file system
4. Add migration untuk existing submissions (optional, can be done gradually)
5. Implement cleanup strategy untuk old submissions

**Files to Create/Modify:**
- Create: `backend/internal/storage/file_storage.go` — File storage service
- Modify: `backend/internal/models/models.go` — Update Submission model
- Modify: `backend/internal/submissions/service.go` — Use file storage
- Modify: `backend/internal/submissions/handler.go` — Handle file operations
- Create: `backend/internal/storage/storage_test.go` — Storage tests

### Correctness Properties

**Property 1: Data Integrity**
```
FOR ALL submission
  ASSERT read_from_storage(submission.code_path) == original_code
```

**Property 2: Atomic Operations**
```
FOR ALL submission
  ASSERT (submission exists in DB) IFF (code file exists in storage)
```

**Property 3: Performance Improvement**
```
FOR ALL queries
  ASSERT query_time_after_migration < query_time_before_migration
```

---

## Bug #3: SQL Session Error 500 — Design

### Technical Context

**Current Implementation:**
- File: `backend/internal/evaluator/sql_evaluator.go`
- Function: `Evaluate()` creates in-memory SQLite database per query
- Uses unique DSN: `file:memdb{timestamp}?mode=memory&cache=shared`

**Potential Issues:**
1. **DSN Collision**: Multiple queries dengan timestamp yang sama bisa share database
2. **Connection Pooling**: SQLite in-memory dengan `cache=shared` bisa cause race conditions
3. **Error Handling**: Generic 500 error di handler tidak expose actual error
4. **Concurrent Access**: Multiple goroutines accessing same in-memory database

**Current Code (sql_evaluator.go line ~25):**
```go
dsn := fmt.Sprintf("file:memdb%d?mode=memory&cache=shared", time.Now().UnixNano())
db, err := sql.Open("sqlite", dsn)
```

**Problem:**
- `time.Now().UnixNano()` bisa collision pada high concurrency
- `cache=shared` dengan in-memory database bisa cause issues
- Error handling di handler returns generic 500

### Solution Design

**Fix Strategy:**

**1. Unique DSN Generation:**
```go
import "github.com/google/uuid"

func (s *SQLEvaluatorService) Evaluate(schema, query, expected string) EvalResult {
  // Use UUID untuk guarantee uniqueness
  dsn := fmt.Sprintf("file:memdb_%s?mode=memory&cache=private", uuid.New().String())
  // ... rest of code
}
```

**2. Proper Connection Management:**
```go
// Use cache=private untuk avoid shared state
dsn := fmt.Sprintf("file:memdb_%s?mode=memory&cache=private", uuid.New().String())

// Set connection limits
db.SetMaxOpenConns(1)
db.SetMaxIdleConns(1)

// Ensure proper cleanup
defer db.Close()
```

**3. Better Error Handling:**
```go
// In handler.go
result, err := h.service.Submit(req)
if err != nil {
  // Log actual error untuk debugging
  log.Printf("Submission error: %v", err)
  
  // Return appropriate status code
  c.JSON(http.StatusInternalServerError, gin.H{
    "statusCode": http.StatusInternalServerError,
    "message":    "Gagal memproses submission",
    "error":      "Internal Server Error",
  })
  return
}
```

**4. Add Timeout Context:**
```go
// Ensure query timeout is properly enforced
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

rows, err := db.QueryContext(ctx, query)
```

**Files to Modify:**
- `backend/internal/evaluator/sql_evaluator.go` — Fix DSN generation dan connection management
- `backend/internal/submissions/handler.go` — Improve error logging
- `backend/internal/evaluator/sql_evaluator_test.go` — Add concurrent access tests

### Correctness Properties

**Property 1: Concurrent Isolation**
```
FOR ALL concurrent_submissions
  ASSERT each_submission_gets_unique_database_instance
```

**Property 2: Query Timeout**
```
FOR ALL long_running_query
  ASSERT query_execution_time <= 5_seconds
```

**Property 3: Error Handling**
```
FOR ALL invalid_query
  ASSERT http_status_code == 200 AND result.status == "error"
```

---

## Testing Strategy

### Unit Tests

**Bug #1 - Normalization Tests:**
- Test whitespace variations (leading, trailing, multiple spaces)
- Test newline variations (\n, \r\n)
- Test edge cases (empty string, only whitespace)
- Property-based tests dengan fast-check

**Bug #2 - Storage Tests:**
- Test file write/read operations
- Test concurrent file access
- Test cleanup strategy
- Test database reference integrity

**Bug #3 - SQL Concurrency Tests:**
- Test concurrent SQL submissions
- Test unique DSN generation
- Test timeout enforcement
- Test error handling

### Integration Tests

- End-to-end submission flow untuk setiap bug fix
- Test dengan multiple concurrent users
- Test dengan various input types

### Property-Based Tests

- Use fast-check untuk generate random test cases
- Validate normalization idempotence
- Validate storage atomicity
- Validate SQL isolation

---

## Automation Setup

### Custom Agents

**Agent 1: Code Quality Checker**
- Lint Go code dengan `golangci-lint`
- Check TypeScript dengan `eslint`
- Validate test coverage

**Agent 2: Performance Tester**
- Benchmark evaluator performance
- Monitor database query times
- Track file I/O performance

### Git Hooks

**Pre-commit Hook:**
- Run linters
- Run unit tests
- Check code formatting

**Pre-push Hook:**
- Run full test suite
- Run integration tests
- Check performance benchmarks

### CI/CD Pipeline

**On Pull Request:**
1. Run all tests
2. Check code coverage
3. Run linters
4. Run performance benchmarks

**On Merge to Main:**
1. Run full test suite
2. Build Docker images
3. Deploy to staging
4. Run smoke tests

---

## Implementation Timeline

### Phase 1: Bug #1 - Strict Answer Validation (Priority: High)
- Duration: 2-3 days
- Impact: Immediate improvement to user experience
- Risk: Low (isolated change)

### Phase 2: Bug #3 - SQL Session Error 500 (Priority: High)
- Duration: 2-3 days
- Impact: Fix critical user-blocking issue
- Risk: Medium (requires testing concurrent scenarios)

### Phase 3: Bug #2 - Code Storage Best Practice (Priority: Medium)
- Duration: 3-5 days
- Impact: Long-term performance improvement
- Risk: Medium (requires migration strategy)

### Phase 4: Automation Setup (Priority: Medium)
- Duration: 2-3 days
- Impact: Improve development workflow
- Risk: Low (non-blocking)

