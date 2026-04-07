# Tasks — Balik Ngoding

## Phase 1: Project Setup & Infrastructure

- [x] 1.1 Initialize Go backend project with Go modules, Gin framework, GORM, PostgreSQL
- [x] 1.2 Initialize Next.js 14 frontend project with TailwindCSS and TypeScript
- [x] 1.3 Configure environment variables (.env) for both frontend and backend
- [x] 1.4 Set up PostgreSQL database and run initial schema migrations
- [x] 1.5 Configure CORS on Gin to allow frontend origin
- [x] 1.6 Install and configure `fast-check` for frontend property-based testing and `pgregory.net/rapid` for Go backend property-based testing

---

## Phase 2: Backend — Data Layer

- [x] 2.1 Create `Problem` Go struct with GORM tags for all required fields (Title, Description, Category, Difficulty, StarterCode, IsActive)
- [x] 2.2 Create `TestCase` Go struct with GORM tags (Input, ExpectedOutput, IsHidden, ProblemID FK)
- [x] 2.3 Create `Submission` Go struct with GORM tags (ProblemID, Code, Language, Status, Score, Total, ResultDetail as JSONB)
- [x] 2.4 Create database indexes on `test_cases.problem_id` and `submissions.problem_id`
- [x] 2.5 Seed database with at least 5 sample problems across all categories with test cases

---

## Phase 3: Backend — Problems API

- [x] 3.1 Implement `GET /problems` Gin handler returning only active problems (isActive = true)
- [x] 3.2 Add optional `category` and `difficulty` query param filtering to `GET /problems`
- [x] 3.3 Implement `GET /problems/:id` Gin handler returning problem detail with non-hidden test cases as examples
- [x] 3.4 Return HTTP 404 with clear message when problem ID is not found
- [x] 3.5 Apply middleware for consistent `{ data: ... }` response format
- [x] 3.6 Apply error handling middleware for consistent error response format

---

## Phase 4: Backend — Evaluator

- [x] 4.1 Implement `EvaluatorService` in Go using `goja` (Go JS runtime) for sandboxed JS execution, or `os/exec` subprocess as MVP fallback
- [x] 4.2 Enforce per-test-case execution timeout (max 5 seconds) using Go context with deadline
- [x] 4.3 Block access to `fs`, `http`, `https`, `process.env`, and other restricted APIs inside the JS sandbox
- [x] 4.4 Normalize actual output (trim whitespace) before comparing with expected output
- [x] 4.5 Catch runtime errors and syntax errors, returning descriptive error messages

---

## Phase 5: Backend — Submissions API

- [x] 5.1 Implement `POST /submit` Gin handler with request binding validation (problemId UUID, code non-empty, language)
- [x] 5.2 Fetch all test cases (including hidden) for the given problemId
- [x] 5.3 Run evaluator against each test case and collect results
- [x] 5.4 Determine submission status: `accepted` / `wrong_answer` / `error` based on results
- [x] 5.5 Calculate score (passed count) and total (test case count)
- [x] 5.6 Persist submission to database with all required fields
- [x] 5.7 Return `SubmissionResult` response with status, score, total, and per-test-case results

---

## Phase 6: Frontend — Shared Infrastructure

- [x] 6.1 Define TypeScript types/interfaces in `lib/types.ts` (Problem, TestCase, SubmissionResult, etc.)
- [x] 6.2 Implement API client functions in `lib/api.ts` (getProblems, getProblemById, submitSolution)
- [x] 6.3 Implement Zustand store in `store/submissionStore.ts` (code, isLoading, result, setters)
- [x] 6.4 Create root layout with consistent Navbar component (logo + link to /problems)
- [x] 6.5 Configure dynamic import for Monaco Editor to prevent SSR crash

---

## Phase 7: Frontend — Problem List Page

- [x] 7.1 Implement `/problems` page that fetches and displays all active problems
- [x] 7.2 Implement `ProblemTable` component showing title, category, difficulty badge per row
- [x] 7.3 Implement `CategoryFilter` tab component (Semua / Loop / String / Array / SQL)
- [x] 7.4 Wire category filter to client-side filtering of the problem list
- [x] 7.5 Limit displayed problems to max 30 items
- [x] 7.6 Navigate to `/problems/:id` when a row is clicked
- [x] 7.7 Display informative error banner when backend is unreachable

---

## Phase 8: Frontend — Problem Detail Page

- [x] 8.1 Implement `/problems/[id]` page that fetches problem detail by ID
- [x] 8.2 Implement `ProblemDescription` component rendering title, description, constraints
- [x] 8.3 Implement `ExampleBlock` component rendering non-hidden test case examples
- [x] 8.4 Redirect to `/problems` with error message when problem ID is not found (404)
- [x] 8.5 Initialize Code Editor with `starterCode` from the fetched problem

---

## Phase 9: Frontend — Code Editor & Submit

- [x] 9.1 Implement `CodeEditor` component wrapping Monaco Editor with `vs-dark` theme, line numbers, min height 400px
- [x] 9.2 Implement fallback `<textarea>` when Monaco Editor fails to load
- [x] 9.3 Implement `SubmitButton` component that is disabled when editor is empty or whitespace-only
- [x] 9.4 Show validation message "Kode tidak boleh kosong" when submit is attempted with empty editor
- [x] 9.5 Disable submit button and show loading indicator while submission is in-flight
- [x] 9.6 Prevent duplicate submissions by ignoring clicks while `isLoading` is true
- [x] 9.7 On successful submission, update Zustand store with result and display Result Panel
- [x] 9.8 On network error, display actionable error message in Result Panel and re-enable submit button

---

## Phase 10: Frontend — Result Panel

- [x] 10.1 Implement `ResultPanel` component showing per-test-case results (passed/failed, input, expected, actual)
- [x] 10.2 Display score in format "X dari Y test case passed"
- [x] 10.3 Show "Coba Lagi" button that resets Code Editor to original `starterCode`
- [x] 10.4 Show "Soal Berikutnya" button that navigates to the next problem

---

## Phase 11: Frontend — Navigation & Error Pages

- [x] 11.1 Implement landing page (`/`) with hero section and CTA linking to `/problems`
- [x] 11.2 Implement 404 page with clear message and link back to `/problems`
- [x] 11.3 Verify user can reach problem detail page from landing page in max 2 clicks

---

## Phase 12: Property-Based Tests — Backend

- [x] 12.1 Write property test for P1: GET /problems returns only active problems (Go + `pgregory.net/rapid`)
- [x] 12.2 Write property test for P2: Category filter returns only matching problems (Go + `pgregory.net/rapid`)
- [x] 12.3 Write property test for P3: Problem list never exceeds 30 items (Go + `pgregory.net/rapid`)
- [x] 12.4 Write property test for P4: Problem detail contains all required fields (Go + `pgregory.net/rapid`)
- [x] 12.5 Write property test for P8: Evaluation covers all test cases (including hidden) (Go + `pgregory.net/rapid`)
- [x] 12.6 Write property test for P11: Submission status is correctly determined (accepted/wrong_answer/error) (Go + `pgregory.net/rapid`)
- [x] 12.7 Write property test for P12: Output normalization before comparison (Go + `pgregory.net/rapid`)
- [x] 12.8 Write property test for P13: Problem detail only exposes non-hidden test cases (Go + `pgregory.net/rapid`)
- [x] 12.9 Write property test for P14: Submission is persisted with all required fields (Go + `pgregory.net/rapid`)
- [x] 12.10 Write property test for P15: Sandbox blocks restricted resource access (Go + `pgregory.net/rapid`)
- [x] 12.11 Write example test for timeout enforcement (infinite loop → error + "Waktu eksekusi habis")

---

## Phase 13: Property-Based Tests — Frontend

- [x] 13.1 Write property test for P5: Editor initializes with starter code
- [x] 13.2 Write property test for P6: Empty/whitespace code submission is rejected
- [x] 13.3 Write property test for P7: Duplicate submit is idempotent (only one request sent)
- [x] 13.4 Write property test for P9: Result panel contains required fields per test case
- [x] 13.5 Write property test for P10: Score display matches actual counts
- [x] 13.6 Write property test for P16: "Coba Lagi" resets editor to starter code
- [x] 13.7 Write example test for problem not found redirect (Req 2.3)
- [x] 13.8 Write example test for 404 page (Req 6.4)
