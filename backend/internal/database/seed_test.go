package database

import (
	"encoding/json"
	"regexp"
	"strings"
	"testing"
)

var funcDeclRegex = regexp.MustCompile(`function\s+[a-zA-Z_$][a-zA-Z0-9_$]*`)

// loadAllSeedProblems loads all problems from the embedded JSON seed files.
func loadAllSeedProblems(t *testing.T) []SeedProblem {
	t.Helper()
	var all []SeedProblem
	for _, filename := range seedFileNames {
		problems, err := loadProblemsFromFile(seedFiles, filename)
		if err != nil {
			t.Fatalf("failed to load %s: %v", filename, err)
		}
		all = append(all, problems...)
	}
	return all
}

// Feature: tambah-soal, Property 1: Semua soal JS memiliki field wajib yang valid
// Validates: Requirements 1.3, 2.3, 3.3, 5.2, 7.1
func TestSeedJSProblemsHaveRequiredFields(t *testing.T) {
	problems := loadAllSeedProblems(t)
	jsCategories := map[string]bool{"loop": true, "string": true, "array": true}

	for _, p := range problems {
		if !jsCategories[p.Category] {
			continue
		}
		if p.Title == "" {
			t.Errorf("JS problem has empty Title")
		}
		if p.Description == "" {
			t.Errorf("JS problem %q has empty Description", p.Title)
		}
		if p.StarterCode == "" {
			t.Errorf("JS problem %q has empty StarterCode", p.Title)
		}
		if p.Category != "loop" && p.Category != "string" && p.Category != "array" {
			t.Errorf("JS problem %q has invalid Category: %q", p.Title, p.Category)
		}
		if p.Difficulty != "easy" && p.Difficulty != "medium" && p.Difficulty != "hard" {
			t.Errorf("JS problem %q has invalid Difficulty: %q", p.Title, p.Difficulty)
		}
	}
}

// Feature: tambah-soal, Property 2: Semua soal SQL memiliki field Schema non-kosong
// Validates: Requirements 4.3
func TestSeedSQLProblemsHaveSchema(t *testing.T) {
	problems := loadAllSeedProblems(t)
	for _, p := range problems {
		if p.Category != "sql" {
			continue
		}
		if p.Schema == "" {
			t.Errorf("SQL problem %q has empty Schema", p.Title)
		}
		if !strings.Contains(strings.ToUpper(p.Schema), "CREATE TABLE") {
			t.Errorf("SQL problem %q Schema does not contain CREATE TABLE", p.Title)
		}
	}
}

// Feature: tambah-soal, Property 3: Setiap soal memiliki jumlah test case minimum sesuai difficulty
// Validates: Requirements 1.4, 2.4, 3.4, 4.4, 6.2
func TestSeedProblemsHaveMinimumTestCases(t *testing.T) {
	problems := loadAllSeedProblems(t)
	for _, p := range problems {
		visible := 0
		hidden := 0
		for _, tc := range p.TestCases {
			if tc.IsHidden {
				hidden++
			} else {
				visible++
			}
		}

		if visible < 3 {
			t.Errorf("Problem %q (%s/%s) has only %d visible test cases (need >= 3)", p.Title, p.Category, p.Difficulty, visible)
		}
		if hidden < 2 {
			t.Errorf("Problem %q (%s/%s) has only %d hidden test cases (need >= 2)", p.Title, p.Category, p.Difficulty, hidden)
		}
	}
}

// Feature: tambah-soal, Property 4: Tidak ada dua soal dengan judul yang sama
// Validates: Requirements 5.4
func TestSeedNoDuplicateTitles(t *testing.T) {
	problems := loadAllSeedProblems(t)
	titles := make(map[string]bool)
	for _, p := range problems {
		if titles[p.Title] {
			t.Errorf("Duplicate title found: %q", p.Title)
		}
		titles[p.Title] = true
	}
}

// Feature: tambah-soal, Property 5: Format expected output valid JSON
// Validates: Requirements 5.3, 7.5
func TestSeedExpectedOutputValidJSON(t *testing.T) {
	problems := loadAllSeedProblems(t)
	for _, p := range problems {
		for i, tc := range p.TestCases {
			var v interface{}
			if err := json.Unmarshal([]byte(tc.ExpectedOutput), &v); err != nil {
				t.Errorf("Problem %q test case %d has invalid JSON expected output %q: %v",
					p.Title, i, tc.ExpectedOutput, err)
			}
		}
	}
}

// Feature: tambah-soal, Property 9: StarterCode soal non-SQL mengandung deklarasi fungsi
// Validates: Requirements 5.2, 7.1, 7.3
func TestSeedNonSQLStarterCodeHasFunctionDecl(t *testing.T) {
	problems := loadAllSeedProblems(t)
	for _, p := range problems {
		if p.Category == "sql" {
			continue
		}
		if !funcDeclRegex.MatchString(p.StarterCode) {
			t.Errorf("Non-SQL problem %q StarterCode does not contain a function declaration: %q",
				p.Title, p.StarterCode)
		}
	}
}

// Feature: sql-seed-schema-bugfix, Property 1: Bug Condition - SQL Schema Missing Error
// Validates: Requirements 1.1, 1.2, 1.3
//
// This test demonstrates the bug exists on unfixed code by checking that SQL problems
// in the seed file have empty schema fields. The test MUST FAIL on unfixed code
// (when schema fields are empty) to prove the bug exists. When the bug is fixed
// (schema fields populated), this test will PASS.
//
// The test loads all SQL problems from the seed file and verifies that they have
// non-empty schema fields. On unfixed code, SQL problems will have empty schemas,
// causing this test to fail and proving the bug exists.
func TestSQLSubmissionBugConditionExplorationProperty(t *testing.T) {
	// Load all SQL problems from the seed file
	problems := loadAllSeedProblems(t)

	// Filter to only SQL problems
	var sqlProblems []SeedProblem
	for _, p := range problems {
		if p.Category == "sql" {
			sqlProblems = append(sqlProblems, p)
		}
	}

	if len(sqlProblems) == 0 {
		t.Fatal("no SQL problems found in seed file")
	}

	// Property: SQL problems MUST have non-empty schema fields
	// On unfixed code, this assertion will fail, proving the bug exists
	// The bug condition is: problem.Category == "sql" && problem.Schema == ""
	for _, problem := range sqlProblems {
		if problem.Schema == "" {
			t.Fatalf("BUG CONDITION DETECTED: SQL problem %q (ID: %s) has empty schema field. "+
				"This causes submission service to return error 'Schema soal SQL tidak ditemukan'",
				problem.Title, problem.ID)
		}

		// Verify schema contains CREATE TABLE statements
		if !strings.Contains(strings.ToUpper(problem.Schema), "CREATE TABLE") {
			t.Fatalf("SQL problem %q schema does not contain CREATE TABLE statement", problem.Title)
		}
	}
}
