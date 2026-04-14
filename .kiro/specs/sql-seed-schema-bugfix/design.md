# SQL Seed Schema Bugfix Design

## Overview

The SQL submission endpoint fails with a 500 error because SQL problems in the seed file lack the `schema` field. The SQL evaluator requires each SQL problem to have a `schema` field containing SQLite CREATE TABLE and INSERT statements that set up the database before executing user queries. This design specifies the exact format and structure needed to populate the schema field for all SQL problems in the seed file.

## Glossary

- **Bug_Condition (C)**: SQL problems have an empty or missing `schema` field in the seed file
- **Property (P)**: Each SQL problem must have a properly formatted `schema` field with CREATE TABLE and INSERT statements
- **Preservation**: Non-SQL problems (loop, string, array) must continue to work without schema fields
- **Schema Field**: A text field in the Problem model containing SQLite SQL statements (CREATE TABLE + INSERT)
- **Test Case Data**: The input/expectedOutput pairs in testCases that define what data should be in the database
- **SQLite**: The in-memory database engine used by the SQL evaluator to sandbox user queries

## Bug Details

### Bug Condition

The bug manifests when a user submits a SQL solution for any problem. The submission service checks if `problem.Category == "sql" && problem.Schema == ""` and returns the error "Schema soal SQL tidak ditemukan" (SQL problem schema not found). This occurs because the seed file (`backend/internal/database/seeds/sql.json`) does not populate the `schema` field for any SQL problems.

**Formal Specification:**
```
FUNCTION isBugCondition(problem)
  INPUT: problem of type Problem
  OUTPUT: boolean
  
  RETURN problem.Category = "sql" 
         AND problem.Schema = ""
         AND userSubmitsCode(problem)
END FUNCTION
```

### Examples

**Example 1: SELECT All Columns Problem**
- Problem ID: `00000000-0000-0000-0004-000000000001`
- Title: "SELECT Semua Kolom"
- Current State: `schema: ""`
- Expected State: `schema: "CREATE TABLE karyawan (id INTEGER PRIMARY KEY, nama TEXT, gaji INTEGER); INSERT INTO karyawan VALUES (1, 'Alice', 5000); INSERT INTO karyawan VALUES (2, 'Bob', 3000);"`
- When user submits: Error "Schema soal SQL tidak ditemukan"

**Example 2: WHERE Condition Problem**
- Problem ID: `00000000-0000-0000-0004-000000000003`
- Title: "WHERE Kondisi Sederhana"
- Current State: `schema: ""`
- Expected State: Schema with karyawan table and sample data matching test cases
- When user submits: Error "Schema soal SQL tidak ditemukan"

**Example 3: INNER JOIN Problem**
- Problem ID: `00000000-0000-0000-0004-000000000010`
- Title: "INNER JOIN — Gabungkan Dua Tabel"
- Current State: `schema: ""`
- Expected State: Schema with both pelanggan and pesanan tables with sample data
- When user submits: Error "Schema soal SQL tidak ditemukan"

## Expected Behavior

### Preservation Requirements

**Unchanged Behaviors:**
- Non-SQL problems (loop, string, array) must continue to work without schema fields
- The submission service must continue to validate non-SQL problems correctly
- The seed loading logic must continue to work for all problem categories
- Test case evaluation for non-SQL problems must remain unchanged
- The database schema and models must not be modified

**Scope:**
All inputs that do NOT involve SQL problems should be completely unaffected by this fix. This includes:
- Loop problems (array iteration, nested loops)
- String problems (string manipulation, parsing)
- Array problems (array operations, searching)
- All non-SQL submission flows

## Hypothesized Root Cause

Based on the bug description, the root cause is:

1. **Missing Schema Population in Seed File**: The `backend/internal/database/seeds/sql.json` file was created without populating the `schema` field for any SQL problems. The field exists in the Problem model but is empty in the seed data.

2. **Seed File Structure**: The JSON seed file contains problem definitions but lacks the CREATE TABLE and INSERT statements needed to initialize the database for each problem.

3. **Data Mapping Gap**: There is no mechanism to automatically derive the schema from test case data. The schema must be manually created for each problem based on the table structure described in the problem description and the data in the test cases.

## Correctness Properties

Property 1: Bug Condition - SQL Schema Population

_For any_ SQL problem where the schema field is currently empty, the fixed seed file SHALL populate the schema field with valid SQLite CREATE TABLE and INSERT statements that match the table structure described in the problem description and the data in the test cases.

**Validates: Requirements 2.1, 2.2, 2.3**

Property 2: Preservation - Non-SQL Problems Unchanged

_For any_ non-SQL problem (loop, string, array), the fixed seed file SHALL NOT modify the problem definition, and the submission service SHALL continue to evaluate these problems correctly without requiring a schema field.

**Validates: Requirements 3.1, 3.2**

## Fix Implementation

### Changes Required

The fix is primarily a data update to the seed file, not a code change. The seed loading logic does not need modification.

**File**: `backend/internal/database/seeds/sql.json`

**Approach**: For each SQL problem in the seed file, populate the `schema` field with:
1. CREATE TABLE statements that define the table structure
2. INSERT statements that populate the tables with sample data from the test cases

**Schema Format for Each Problem**:

The schema field should contain valid SQLite SQL statements. Format:
```sql
CREATE TABLE table_name (
  column1 TYPE,
  column2 TYPE,
  ...
);
INSERT INTO table_name VALUES (...);
INSERT INTO table_name VALUES (...);
...
```

**Specific Changes**:

1. **Problem 00000000-0000-0000-0004-000000000001 (SELECT Semua Kolom)**
   - Table: `karyawan` with columns: id (INTEGER PRIMARY KEY), nama (TEXT), gaji (INTEGER)
   - Sample data from first test case: (1, 'Alice', 5000), (2, 'Bob', 3000)
   - Schema: CREATE TABLE + INSERT statements for karyawan

2. **Problem 00000000-0000-0000-0004-000000000002 (SELECT Kolom Tertentu)**
   - Table: `karyawan` with same structure as problem 1
   - Sample data: (1, 'Alice', 5000), (2, 'Bob', 3000)
   - Schema: CREATE TABLE + INSERT statements for karyawan

3. **Problem 00000000-0000-0000-0004-000000000003 (WHERE Kondisi Sederhana)**
   - Table: `karyawan` with same structure
   - Sample data: (1, 'Alice', 5000), (2, 'Bob', 3000), (3, 'Citra', 4500)
   - Schema: CREATE TABLE + INSERT statements for karyawan

4. **Problem 00000000-0000-0000-0004-000000000004 (ORDER BY)**
   - Table: `karyawan` with same structure
   - Sample data: (1, 'Alice', 5000), (2, 'Bob', 3000), (3, 'Eka', 7000)
   - Schema: CREATE TABLE + INSERT statements for karyawan

5. **Problem 00000000-0000-0000-0004-000000000005 (LIMIT)**
   - Table: `produk` with columns: id (INTEGER PRIMARY KEY), nama (TEXT), harga (INTEGER)
   - Sample data: (1, 'Laptop', 15000000), (2, 'HP', 8000000), (3, 'Tablet', 5000000), (4, 'Charger', 200000), (5, 'Kabel', 50000), (6, 'Mouse', 300000)
   - Schema: CREATE TABLE + INSERT statements for produk

6. **Problem 00000000-0000-0000-0004-000000000006 (COUNT)**
   - Table: `siswa` with columns: id (INTEGER PRIMARY KEY), nama (TEXT), kelas (TEXT)
   - Sample data: (1, 'Andi', '10A'), (2, 'Budi', '10B'), (3, 'Cici', '10A'), (4, 'Deni', '10A')
   - Schema: CREATE TABLE + INSERT statements for siswa

7. **Problem 00000000-0000-0000-0004-000000000007 (DISTINCT)**
   - Table: `penjualan` with columns: id (INTEGER PRIMARY KEY), produk (TEXT), kota (TEXT)
   - Sample data: (1, 'A', 'Jakarta'), (2, 'B', 'Bandung'), (3, 'C', 'Jakarta'), (4, 'D', 'Surabaya')
   - Schema: CREATE TABLE + INSERT statements for penjualan

8. **Problem 00000000-0000-0000-0004-000000000008 (GROUP BY)**
   - Table: `penjualan` with columns: id (INTEGER PRIMARY KEY), produk (TEXT), kota (TEXT), jumlah (INTEGER)
   - Sample data: (1, 'A', 'Jakarta', 10), (2, 'B', 'Bandung', 5), (3, 'C', 'Jakarta', 20), (4, 'D', 'Surabaya', 10), (5, 'E', 'Bandung', 10)
   - Schema: CREATE TABLE + INSERT statements for penjualan

9. **Problem 00000000-0000-0000-0004-000000000009 (HAVING)**
   - Table: `penjualan` with same structure as problem 8
   - Sample data: (1, 'A', 'Jakarta', 10), (2, 'B', 'Bandung', 15), (3, 'C', 'Jakarta', 20), (4, 'D', 'Surabaya', 5), (5, 'E', 'Bandung', 10)
   - Schema: CREATE TABLE + INSERT statements for penjualan

10. **Problem 00000000-0000-0000-0004-000000000010 (INNER JOIN)**
    - Tables: `pelanggan` (id, nama, kota) and `pesanan` (id, pelanggan_id, total)
    - Sample data: pelanggan: (1, 'Alice', 'Jakarta'), (2, 'Bob', 'Bandung'); pesanan: (1, 1, 150000), (2, 2, 75000)
    - Schema: CREATE TABLE + INSERT statements for both tables

11. **Problem 00000000-0000-0000-0004-000000000011 (LEFT JOIN)**
    - Tables: `departemen` (id, nama_dept) and `karyawan` (id, nama, departemen_id)
    - Sample data: departemen: (1, 'Engineering'), (2, 'Marketing'); karyawan: (1, 'Alice', 1), (2, 'Bob', NULL), (3, 'Citra', 2)
    - Schema: CREATE TABLE + INSERT statements for both tables

### Data Mapping Strategy

For each problem:
1. Extract the first test case's input data (which contains the table structure and sample data)
2. Parse the table structure from the input JSON
3. Generate CREATE TABLE statements with appropriate column types
4. Generate INSERT statements from the sample data
5. Combine into a single schema string

### Example Schema Generation

For problem "SELECT Semua Kolom" with test case input:
```json
{"tables":{"karyawan":[{"id":1,"nama":"Alice","gaji":5000},{"id":2,"nama":"Bob","gaji":3000}]}}
```

Generated schema:
```sql
CREATE TABLE karyawan (id INTEGER PRIMARY KEY, nama TEXT, gaji INTEGER);
INSERT INTO karyawan VALUES (1, 'Alice', 5000);
INSERT INTO karyawan VALUES (2, 'Bob', 3000);
```

## Testing Strategy

### Validation Approach

The testing strategy follows a two-phase approach: first, surface counterexamples that demonstrate the bug on unfixed code, then verify the fix works correctly and preserves existing behavior.

### Exploratory Bug Condition Checking

**Goal**: Surface counterexamples that demonstrate the bug BEFORE implementing the fix. Confirm that SQL problems fail with the schema error.

**Test Plan**: Write tests that attempt to submit SQL solutions for various problems and assert that the error "Schema soal SQL tidak ditemukan" is returned. Run these tests on the UNFIXED seed file to observe failures.

**Test Cases**:
1. **SELECT All Columns Submission**: Submit a correct query for problem 1 (will fail on unfixed code)
2. **WHERE Condition Submission**: Submit a correct query for problem 3 (will fail on unfixed code)
3. **INNER JOIN Submission**: Submit a correct query for problem 10 (will fail on unfixed code)
4. **Multiple Table JOIN**: Submit a query for problem 11 (will fail on unfixed code)

**Expected Counterexamples**:
- Submission returns 500 error with message "Schema soal SQL tidak ditemukan"
- No schema field is populated in the seed data

### Fix Checking

**Goal**: Verify that for all SQL problems, the schema field is now populated and submissions can be evaluated.

**Pseudocode:**
```
FOR ALL sqlProblem IN seedFile WHERE sqlProblem.category = "sql" DO
  ASSERT sqlProblem.schema != ""
  ASSERT sqlProblem.schema CONTAINS "CREATE TABLE"
  ASSERT sqlProblem.schema CONTAINS "INSERT INTO"
  
  result := submitCode(sqlProblem, correctQuery)
  ASSERT result.error != "Schema soal SQL tidak ditemukan"
  ASSERT result.status IN ["accepted", "wrong_answer", "error"]
END FOR
```

### Preservation Checking

**Goal**: Verify that non-SQL problems continue to work correctly and are not affected by the schema field addition.

**Pseudocode:**
```
FOR ALL nonSqlProblem IN seedFile WHERE nonSqlProblem.category IN ["loop", "string", "array"] DO
  ASSERT nonSqlProblem.schema = "" OR nonSqlProblem.schema NOT PRESENT
  
  result := submitCode(nonSqlProblem, correctCode)
  ASSERT result.status IN ["accepted", "wrong_answer", "error"]
  ASSERT result.error != "Schema soal SQL tidak ditemukan"
END FOR
```

**Testing Approach**: Property-based testing is recommended for preservation checking because:
- It generates many test cases automatically across the input domain
- It catches edge cases that manual unit tests might miss
- It provides strong guarantees that behavior is unchanged for all non-SQL problems

**Test Plan**: Verify that loop, string, and array problems continue to work correctly after the seed file is updated.

**Test Cases**:
1. **Loop Problem Submission**: Submit code for a loop problem (should work as before)
2. **String Problem Submission**: Submit code for a string problem (should work as before)
3. **Array Problem Submission**: Submit code for an array problem (should work as before)
4. **Multiple Problem Types**: Submit solutions for different problem types in sequence

### Unit Tests

- Test that each SQL problem has a non-empty schema field
- Test that schema field contains valid SQLite SQL statements
- Test that schema field contains CREATE TABLE statements
- Test that schema field contains INSERT statements with correct data
- Test that non-SQL problems do not have schema fields populated

### Property-Based Tests

- Generate random SQL problems and verify schema field is populated
- Generate random non-SQL problems and verify schema field is empty
- Test that SQL evaluator can execute queries against populated schemas
- Test that submission service accepts SQL submissions without schema errors

### Integration Tests

- Test full submission flow for SQL problems with populated schemas
- Test that correct SQL queries pass evaluation
- Test that incorrect SQL queries fail evaluation appropriately
- Test that non-SQL problems continue to work correctly
- Test that seed file loads correctly with populated schemas
