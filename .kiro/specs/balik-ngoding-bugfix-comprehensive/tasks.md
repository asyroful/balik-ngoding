# Tasks — Balik Ngoding Comprehensive Bugfix

## Phase 1: Bug #1 - Strict Answer Validation

### 1.1 Implement Output Normalization Function
- [ ] Create `normalizeOutput()` function in `backend/internal/evaluator/evaluator.go`
- [ ] Function SHALL trim whitespace, normalize spaces, dan normalize newlines
- [ ] Add unit tests untuk normalization function
- [ ] Validate idempotence property

### 1.2 Update Evaluator Comparison Logic
- [ ] Modify `Evaluate()` function untuk use normalized output
- [ ] Apply normalization ke both actual dan expected output
- [ ] Ensure comparison uses normalized values
- [ ] Run existing tests untuk verify no regression

### 1.3 Add Property-Based Tests
- [ ] Create property test untuk whitespace tolerance
- [ ] Create property test untuk preserve wrong answers
- [ ] Create property test untuk idempotent normalization
- [ ] Use fast-check untuk generate random test cases

### 1.4 Integration Testing
- [ ] Test dengan various whitespace scenarios
- [ ] Test dengan multiple test cases
- [ ] Test dengan error cases
- [ ] Verify user experience improvement

---

## Phase 2: Bug #3 - SQL Session Error 500

### 2.1 Fix DSN Generation
- [ ] Update `sql_evaluator.go` untuk use UUID instead of timestamp
- [ ] Import `github.com/google/uuid` package
- [ ] Replace `time.Now().UnixNano()` dengan `uuid.New().String()`
- [ ] Ensure unique DSN per query execution

### 2.2 Improve Connection Management
- [ ] Set `cache=private` untuk avoid shared state
- [ ] Configure connection limits: `SetMaxOpenConns(1)`, `SetMaxIdleConns(1)`
- [ ] Ensure proper cleanup dengan defer statements
- [ ] Add connection timeout configuration

### 2.3 Enhance Error Handling
- [ ] Add logging untuk actual errors di handler
- [ ] Improve error messages untuk debugging
- [ ] Ensure HTTP 200 response dengan error status di body
- [ ] Add error context untuk troubleshooting

### 2.4 Add Concurrent Access Tests
- [ ] Create test untuk concurrent SQL submissions
- [ ] Verify each submission gets isolated database
- [ ] Test dengan high concurrency (10+ concurrent requests)
- [ ] Verify no race conditions atau shared state

### 2.5 Add Timeout Tests
- [ ] Test query timeout enforcement
- [ ] Verify timeout error message
- [ ] Test dengan long-running queries
- [ ] Ensure timeout doesn't affect other queries

---

## Phase 3: Bug #2 - Code Storage Best Practice

### 3.1 Create File Storage Service
- [ ] Create `backend/internal/storage/file_storage.go`
- [ ] Implement `SaveCode()` function untuk write code ke file system
- [ ] Implement `ReadCode()` function untuk read code dari file system
- [ ] Implement `DeleteCode()` function untuk cleanup old files
- [ ] Add proper error handling dan logging

### 3.2 Update Submission Model
- [ ] Modify `backend/internal/models/models.go`
- [ ] Replace `Code` field dengan `CodePath` field
- [ ] Update JSON tags untuk proper serialization
- [ ] Add migration notes untuk existing data

### 3.3 Update Submissions Service
- [ ] Modify `backend/internal/submissions/service.go`
- [ ] Update `Submit()` function untuk use file storage
- [ ] Save code ke file system sebelum persist ke database
- [ ] Handle file storage errors properly
- [ ] Add cleanup strategy untuk old submissions

### 3.4 Update Submissions Handler
- [ ] Modify `backend/internal/submissions/handler.go`
- [ ] Update request/response handling untuk new model
- [ ] Ensure backward compatibility jika diperlukan
- [ ] Add proper error handling untuk file operations

### 3.5 Add Storage Tests
- [ ] Create `backend/internal/storage/storage_test.go`
- [ ] Test file write/read operations
- [ ] Test concurrent file access
- [ ] Test cleanup strategy
- [ ] Test error scenarios

### 3.6 Add Integration Tests
- [ ] Test end-to-end submission flow
- [ ] Test dengan multiple concurrent submissions
- [ ] Test file system cleanup
- [ ] Verify database reference integrity

### 3.7 Performance Benchmarking
- [ ] Benchmark query performance sebelum/sesudah
- [ ] Measure database size reduction
- [ ] Measure file I/O performance
- [ ] Document performance improvements

---

## Phase 4: Automation Setup

### 4.1 Create Custom Agents

#### 4.1.1 Code Quality Checker Agent
- [ ] Create `.kiro/agents/code-quality-checker.md`
- [ ] Define linting rules untuk Go code
- [ ] Define linting rules untuk TypeScript code
- [ ] Add test coverage requirements
- [ ] Add documentation requirements

#### 4.1.2 Performance Tester Agent
- [ ] Create `.kiro/agents/performance-tester.md`
- [ ] Define benchmark scenarios
- [ ] Define performance thresholds
- [ ] Add monitoring metrics
- [ ] Add alerting rules

### 4.2 Setup Git Hooks

#### 4.2.1 Pre-commit Hook
- [ ] Create `.git/hooks/pre-commit`
- [ ] Run `golangci-lint` untuk Go code
- [ ] Run `eslint` untuk TypeScript code
- [ ] Run unit tests
- [ ] Check code formatting

#### 4.2.2 Pre-push Hook
- [ ] Create `.git/hooks/pre-push`
- [ ] Run full test suite
- [ ] Run integration tests
- [ ] Run performance benchmarks
- [ ] Check test coverage

### 4.3 Create Steering Files

#### 4.3.1 Go Coding Standards
- [ ] Create `.kiro/steering/go-standards.md`
- [ ] Define naming conventions
- [ ] Define error handling patterns
- [ ] Define testing patterns
- [ ] Define documentation requirements

#### 4.3.2 TypeScript Coding Standards
- [ ] Create `.kiro/steering/typescript-standards.md`
- [ ] Define naming conventions
- [ ] Define type safety rules
- [ ] Define testing patterns
- [ ] Define component patterns

#### 4.3.3 Testing Standards
- [ ] Create `.kiro/steering/testing-standards.md`
- [ ] Define unit test requirements
- [ ] Define integration test requirements
- [ ] Define property-based test requirements
- [ ] Define test coverage thresholds

### 4.4 Setup CI/CD Pipeline

#### 4.4.1 GitHub Actions Workflow
- [ ] Create `.github/workflows/test.yml`
- [ ] Define test job untuk backend
- [ ] Define test job untuk frontend
- [ ] Define linting job
- [ ] Define coverage reporting

#### 4.4.2 Deployment Workflow
- [ ] Create `.github/workflows/deploy.yml`
- [ ] Define build job
- [ ] Define Docker image build
- [ ] Define staging deployment
- [ ] Define smoke tests

---

## Acceptance Criteria Checklist

### Bug #1 - Strict Answer Validation
- [ ] Whitespace variations (leading, trailing, multiple spaces) are handled correctly
- [ ] Newline variations (\n, \r\n) are handled correctly
- [ ] Wrong answers are still marked as wrong
- [ ] All existing tests pass
- [ ] Property-based tests pass
- [ ] User experience improved (verified manually)

### Bug #2 - Code Storage Best Practice
- [ ] File storage service implemented dan tested
- [ ] Submission model updated dengan CodePath
- [ ] Submissions service uses file storage
- [ ] File write/read operations work correctly
- [ ] Concurrent access handled properly
- [ ] Cleanup strategy implemented
- [ ] Performance improved (verified dengan benchmarks)
- [ ] All existing tests pass

### Bug #3 - SQL Session Error 500
- [ ] DSN generation uses UUID untuk uniqueness
- [ ] Connection management properly configured
- [ ] Error handling improved dengan logging
- [ ] Concurrent SQL submissions work correctly
- [ ] Timeout enforcement verified
- [ ] No race conditions detected
- [ ] All existing tests pass
- [ ] Error 500 no longer occurs

### Automation Setup
- [ ] Code quality checker agent created
- [ ] Performance tester agent created
- [ ] Git hooks configured dan working
- [ ] Steering files created dengan standards
- [ ] CI/CD pipeline configured
- [ ] All automation tests pass

---

## Testing Checklist

### Unit Tests
- [ ] Bug #1: Normalization function tests
- [ ] Bug #2: File storage tests
- [ ] Bug #3: DSN generation tests
- [ ] All tests pass locally

### Integration Tests
- [ ] Bug #1: End-to-end submission flow
- [ ] Bug #2: File storage dengan database
- [ ] Bug #3: Concurrent SQL submissions
- [ ] All tests pass locally

### Property-Based Tests
- [ ] Bug #1: Whitespace tolerance property
- [ ] Bug #1: Wrong answer preservation property
- [ ] Bug #1: Normalization idempotence property
- [ ] Bug #3: Concurrent isolation property
- [ ] All properties verified

### Performance Tests
- [ ] Bug #2: Query performance improvement
- [ ] Bug #2: Database size reduction
- [ ] Bug #3: Concurrent request handling
- [ ] Benchmarks documented

---

## Documentation Checklist

- [ ] Update README.md dengan new features
- [ ] Document file storage strategy
- [ ] Document automation setup
- [ ] Document coding standards
- [ ] Document testing strategy
- [ ] Add inline code comments
- [ ] Update API documentation

