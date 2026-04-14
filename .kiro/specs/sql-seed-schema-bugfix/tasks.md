# SQL Seed Schema Bugfix — Implementation Plan

## Overview

This task list implements the SQL seed schema bugfix using the exploratory bugfix workflow:
1. **Explore** - Write tests to demonstrate the bug exists (SQL submissions fail with schema error)
2. **Preserve** - Write tests to verify non-SQL problems continue working
3. **Implement** - Update the seed file to populate schema fields for all SQL problems
4. **Validate** - Verify the fix works and doesn't break anything

---

## Task 1: Write Bug Condition Exploration Test

- [x] 1. Write bug condition exploration test
  - **Property 1: Bug Condition** - SQL Schema Missing Error
  - **CRITICAL**: This test MUST FAIL on unfixed code - failure confirms the bug exists
  - **DO NOT attempt to fix the test or the code when it fails**
  - **NOTE**: This test encodes the expected behavior - it will validate the fix when it passes after implementation
  - **GOAL**: Surface counterexamples that demonstrate the bug exists
  - **Scoped PBT Approach**: For deterministic bugs, scope the property to the concrete failing case(s) to ensure reproducibility
  - Write property-based test that attempts to submit SQL solutions for various SQL problems
  - Test assertions should verify that submissions do NOT return the "Schema soal SQL tidak ditemukan" error
  - The test should check that the submission service can evaluate SQL queries without schema errors
  - Run test on UNFIXED code (with empty schema fields in seed file)
  - **EXPECTED OUTCOME**: Test FAILS (this is correct - it proves the bug exists)
  - Document counterexamples found to understand root cause (e.g., "SQL problem 00000000-0000-0000-0004-000000000001 returns schema error on submission")
  - Mark task complete when test is written, run, and failure is documented
  - _Requirements: 1.1, 1.2, 1.3_

---

## Task 2: Write Preservation Property Tests

- [x] 2. Write preservation property tests (BEFORE implementing fix)
  - **Property 2: Preservation** - Non-SQL Problems Continue Working
  - **IMPORTANT**: Follow observation-first methodology
  - Observe behavior on UNFIXED code for non-SQL problems (loop, string, array)
  - Write property-based tests capturing observed behavior patterns from Preservation Requirements
  - Test that loop problems can be submitted and evaluated correctly
  - Test that string problems can be submitted and evaluated correctly
  - Test that array problems can be submitted and evaluated correctly
  - Property-based testing generates many test cases for stronger guarantees
  - Run tests on UNFIXED code
  - **EXPECTED OUTCOME**: Tests PASS (this confirms baseline behavior to preserve)
  - Mark task complete when tests are written, run, and passing on unfixed code
  - _Requirements: 3.1, 3.2_

---

## Task 3: Fix SQL Seed Schema

- [x] 3. Update SQL seed file with schema fields

  - [x] 3.1 Populate schema for single-table SQL problems
    - Update `backend/internal/database/seeds/sql.json`
    - For each single-table SQL problem (problems 1-9), populate the `schema` field
    - Problem 1 (SELECT Semua Kolom): Add schema with karyawan table (id, nama, gaji)
    - Problem 2 (SELECT Kolom Tertentu): Add schema with karyawan table
    - Problem 3 (WHERE Kondisi Sederhana): Add schema with karyawan table
    - Problem 4 (ORDER BY): Add schema with karyawan table
    - Problem 5 (LIMIT): Add schema with produk table (id, nama, harga)
    - Problem 6 (COUNT): Add schema with siswa table (id, nama, kelas)
    - Problem 7 (DISTINCT): Add schema with penjualan table (id, produk, kota)
    - Problem 8 (GROUP BY): Add schema with penjualan table (id, produk, kota, jumlah)
    - Problem 9 (HAVING): Add schema with penjualan table (id, produk, kota, jumlah)
    - Each schema should contain CREATE TABLE statements with appropriate column types
    - Each schema should contain INSERT statements with sample data from first test case
    - _Bug_Condition: isBugCondition(problem) where problem.category = "sql" AND problem.schema = ""_
    - _Expected_Behavior: expectedBehavior(result) where result.schema contains valid SQLite CREATE TABLE and INSERT statements_
    - _Preservation: Non-SQL problems must not be modified_
    - _Requirements: 2.1, 2.2, 2.3_

  - [x] 3.2 Populate schema for multi-table SQL problems
    - Problem 10 (INNER JOIN): Add schema with pelanggan and pesanan tables
    - Problem 11 (LEFT JOIN): Add schema with departemen and karyawan tables
    - Each schema should contain CREATE TABLE statements for both tables
    - Each schema should contain INSERT statements with sample data from first test case
    - Ensure foreign key relationships are properly represented in the schema
    - _Bug_Condition: isBugCondition(problem) where problem.category = "sql" AND problem.schema = ""_
    - _Expected_Behavior: expectedBehavior(result) where result.schema contains valid SQLite CREATE TABLE and INSERT statements for all tables_
    - _Preservation: Non-SQL problems must not be modified_
    - _Requirements: 2.1, 2.2, 2.3_

  - [x] 3.3 Verify seed file structure and format
    - Ensure all SQL problems now have non-empty schema fields
    - Verify schema fields contain valid SQLite SQL syntax
    - Verify schema fields contain CREATE TABLE statements
    - Verify schema fields contain INSERT statements with correct data
    - Verify non-SQL problems (loop, string, array) are unchanged
    - Verify JSON structure is valid and parseable
    - _Requirements: 2.1, 2.2, 2.3_

---

## Task 4: Verify Bug Condition Exploration Test Now Passes

- [x] 4. Verify bug condition exploration test now passes
  - **Property 1: Expected Behavior** - SQL Schema Population
  - **IMPORTANT**: Re-run the SAME test from task 1 - do NOT write a new test
  - The test from task 1 encodes the expected behavior
  - When this test passes, it confirms the expected behavior is satisfied
  - Run bug condition exploration test from step 1
  - **EXPECTED OUTCOME**: Test PASSES (confirms bug is fixed)
  - Verify that SQL submissions no longer return "Schema soal SQL tidak ditemukan" error
  - Verify that SQL problems can now be evaluated without schema errors
  - _Requirements: 2.1, 2.2, 2.3_

---

## Task 5: Verify Preservation Tests Still Pass

- [x] 5. Verify preservation tests still pass
  - **Property 2: Preservation** - Non-SQL Problems Unchanged
  - **IMPORTANT**: Re-run the SAME tests from task 2 - do NOT write new tests
  - Run preservation property tests from step 2
  - **EXPECTED OUTCOME**: Tests PASS (confirms no regressions)
  - Confirm all non-SQL problem submissions continue to work correctly
  - Confirm loop, string, and array problems are unaffected by the seed file changes
  - _Requirements: 3.1, 3.2_

---

## Task 6: Checkpoint — Ensure All Tests Pass

- [x] 6. Checkpoint - Ensure all tests pass
  - Verify bug condition exploration test passes (from task 4)
  - Verify preservation tests pass (from task 5)
  - Verify seed file loads correctly with populated schemas
  - Verify no regressions in non-SQL problem evaluation
  - Confirm all SQL problems can now be submitted and evaluated
  - Ask the user if questions arise or if any issues are encountered
