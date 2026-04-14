# SQL Seed Schema Bugfix

## Introduction

The SQL submission endpoint is returning 500 errors when users try to submit SQL solutions. The root cause is that the SQL seed data in `backend/internal/database/seeds/sql.json` is missing the `schema` field that the SQL evaluator requires. The SQL evaluator expects each SQL problem to have a `schema` field containing CREATE TABLE and INSERT statements to set up the database. When a user submits a SQL solution, the backend checks if `problem.Schema == ""` and returns an error: "Schema soal SQL tidak ditemukan" (SQL problem schema not found).

## Bug Analysis

### Current Behavior (Defect)

1.1 WHEN a user submits a SQL solution for any problem (e.g., problem ID `00000000-0000-0000-0004-000000000001`) THEN the system returns a 500 error with message "Schema soal SQL tidak ditemukan" because the problem's schema field is empty

1.2 WHEN the backend loads SQL problems from the seed file THEN the system does not populate the `schema` field with CREATE TABLE and INSERT statements, leaving it as an empty string

1.3 WHEN the SQL evaluator attempts to execute a user's query THEN it fails because there is no schema to set up the database tables before running the query

### Expected Behavior (Correct)

2.1 WHEN a user submits a SQL solution for a problem THEN the system successfully evaluates the query against a properly initialized database with the correct schema

2.2 WHEN the backend loads SQL problems from the seed file THEN each problem contains a `schema` field with CREATE TABLE statements that match the table structure described in the problem description

2.3 WHEN the backend loads SQL problems from the seed file THEN each problem's schema includes sample INSERT statements to populate test data that matches the test cases

2.4 WHEN the SQL evaluator executes a user's query THEN it uses the schema to set up the database tables and data before running the query, allowing the query to execute successfully

### Unchanged Behavior (Regression Prevention)

3.1 WHEN a user submits a solution for non-SQL problems (loop, string, array) THEN the system SHALL CONTINUE TO evaluate the code correctly without requiring a schema field

3.2 WHEN the backend loads non-SQL problems from seed files THEN the system SHALL CONTINUE TO work correctly without schema fields

3.3 WHEN the SQL evaluator receives a valid schema and query THEN the system SHALL CONTINUE TO properly validate the query output against expected results

3.4 WHEN a user submits a SQL solution with an invalid query THEN the system SHALL CONTINUE TO return appropriate error messages about the query syntax or execution
