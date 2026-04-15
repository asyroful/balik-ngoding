package submissions

import (
	"balik-ngoding-backend/internal/models"
	"balik-ngoding-backend/internal/storage"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestEvaluatorBugFixE2E verifies the evaluator correctly handles string inputs
// Feature: evaluator-fix-and-solution-keys, Property 16: evaluator bug fix verification
func TestEvaluatorBugFixE2E(t *testing.T) {
	// Setup test database
	db := setupTestDB(t)

	// Create temporary directory for file storage
	tmpDir, err := os.MkdirTemp("", "test-storage-*")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	fileStorage, err := storage.NewFileStorageService(tmpDir)
	if err != nil {
		t.Fatalf("failed to create file storage: %v", err)
	}

	service := NewSubmissionsService(fileStorage)

	// Create test problem with string input test cases
	problem := &models.Problem{
		ID:       "test-string-problem",
		Title:    "String Manipulation",
		Category: "string",
	}
	if err := db.Create(problem).Error; err != nil {
		t.Fatalf("failed to create test problem: %v", err)
	}

	// Create test cases with string inputs
	testCases := []models.TestCase{
		{
			ID:             "tc-1",
			ProblemID:      problem.ID,
			Input:          `"banana", "a"`,
			ExpectedOutput: `"a"`,
		},
		{
			ID:             "tc-2",
			ProblemID:      problem.ID,
			Input:          `"hello"`,
			ExpectedOutput: `"hello"`,
		},
		{
			ID:             "tc-3",
			ProblemID:      problem.ID,
			Input:          `5, "test"`,
			ExpectedOutput: `"test5"`,
		},
	}

	for _, tc := range testCases {
		if err := db.Create(&tc).Error; err != nil {
			t.Fatalf("failed to create test case: %v", err)
		}
	}

	tests := []struct {
		name       string
		code       string
		problemID  string
		wantPassed int
		wantFailed int
	}{
		{
			name:       "return second argument",
			code:       "function solution(a, b) { return b; }",
			problemID:  problem.ID,
			wantPassed: 1,
			wantFailed: 0,
		},
		{
			name:       "return first argument",
			code:       "function solution(a) { return a; }",
			problemID:  problem.ID,
			wantPassed: 1,
			wantFailed: 0,
		},
		{
			name:       "concatenate number and string",
			code:       "function solution(num, str) { return str + num; }",
			problemID:  problem.ID,
			wantPassed: 1,
			wantFailed: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := SubmitRequest{
				ProblemID: tt.problemID,
				Code:      tt.code,
				Language:  "javascript",
			}

			result, err := service.Submit(req)
			if err != nil {
				t.Fatalf("submission failed: %v", err)
			}

			if result == nil {
				t.Fatal("expected result, got nil")
			}

			// Count passed and failed test cases
			passed := 0
			failed := 0
			for _, r := range result.Results {
				if r.Passed {
					passed++
				} else {
					failed++
				}
			}

			if passed != tt.wantPassed {
				t.Errorf("expected %d passed, got %d", tt.wantPassed, passed)
			}
			if failed != tt.wantFailed {
				t.Errorf("expected %d failed, got %d", tt.wantFailed, failed)
			}
		})
	}
}

// TestSolutionKeyManagementE2E verifies complete solution key CRUD operations
// Feature: evaluator-fix-and-solution-keys, Property 8: solution key round-trip
func TestSolutionKeyManagementE2E(t *testing.T) {
	db := setupTestDB(t)

	// Create test problem
	problem := &models.Problem{
		ID:       "test-solution-problem",
		Title:    "Test Problem",
		Category: "loop",
	}
	if err := db.Create(problem).Error; err != nil {
		t.Fatalf("failed to create test problem: %v", err)
	}

	// Test data
	solutionCode := "function solution(n) { return n * 2; }"
	language := "javascript"

	// Test 1: Create solution key
	sk := &models.SolutionKey{
		ProblemID: problem.ID,
		Code:      solutionCode,
		Language:  language,
	}

	if err := db.Create(sk).Error; err != nil {
		t.Fatalf("failed to create solution key: %v", err)
	}

	if sk.ID == "" {
		t.Fatal("solution key ID should be generated")
	}

	// Test 2: Retrieve solution key by problem ID
	var retrieved models.SolutionKey
	if err := db.Where("problem_id = ?", problem.ID).First(&retrieved).Error; err != nil {
		t.Fatalf("failed to retrieve solution key: %v", err)
	}

	if retrieved.Code != solutionCode {
		t.Errorf("expected code %q, got %q", solutionCode, retrieved.Code)
	}

	if retrieved.Language != language {
		t.Errorf("expected language %q, got %q", language, retrieved.Language)
	}

	// Test 3: Update solution key
	updatedCode := "function solution(n) { return n * 3; }"
	if err := db.Model(&retrieved).Update("code", updatedCode).Error; err != nil {
		t.Fatalf("failed to update solution key: %v", err)
	}

	// Verify update
	var updated models.SolutionKey
	if err := db.First(&updated, "id = ?", retrieved.ID).Error; err != nil {
		t.Fatalf("failed to fetch updated solution key: %v", err)
	}

	if updated.Code != updatedCode {
		t.Errorf("expected updated code %q, got %q", updatedCode, updated.Code)
	}

	// Verify createdAt is preserved
	if !retrieved.CreatedAt.Equal(updated.CreatedAt) {
		t.Error("createdAt should be preserved on update")
	}

	// Verify updatedAt is changed
	if !updated.UpdatedAt.After(retrieved.UpdatedAt) {
		t.Error("updatedAt should be updated")
	}

	// Test 4: Delete solution key
	if err := db.Delete(&updated).Error; err != nil {
		t.Fatalf("failed to delete solution key: %v", err)
	}

	// Verify deletion
	var deleted models.SolutionKey
	result := db.First(&deleted, "id = ?", updated.ID)
	if result.Error != gorm.ErrRecordNotFound {
		t.Errorf("expected record not found, got error: %v", result.Error)
	}

	// Test 5: Verify problem still exists after solution key deletion
	var problemAfterDelete models.Problem
	if err := db.First(&problemAfterDelete, "id = ?", problem.ID).Error; err != nil {
		t.Fatalf("problem should still exist after solution key deletion: %v", err)
	}
}

// TestCascadeDeleteE2E verifies that deleting a problem cascade-deletes solution keys
// Feature: evaluator-fix-and-solution-keys, Property 11: cascade delete behavior
func TestCascadeDeleteE2E(t *testing.T) {
	db := setupTestDB(t)

	// Create test problem
	problem := &models.Problem{
		ID:       "test-cascade-problem",
		Title:    "Cascade Test",
		Category: "loop",
	}
	if err := db.Create(problem).Error; err != nil {
		t.Fatalf("failed to create test problem: %v", err)
	}

	// Create multiple solution keys for different languages
	sk1 := &models.SolutionKey{
		ProblemID: problem.ID,
		Code:      "function solution() { return 1; }",
		Language:  "javascript",
	}
	sk2 := &models.SolutionKey{
		ProblemID: problem.ID,
		Code:      "SELECT 1;",
		Language:  "sql",
	}

	if err := db.Create(sk1).Error; err != nil {
		t.Fatalf("failed to create solution key 1: %v", err)
	}
	if err := db.Create(sk2).Error; err != nil {
		t.Fatalf("failed to create solution key 2: %v", err)
	}

	// Delete problem
	if err := db.Delete(problem).Error; err != nil {
		t.Fatalf("failed to delete problem: %v", err)
	}

	// Verify solution keys are cascade-deleted
	var count int64
	if err := db.Model(&models.SolutionKey{}).Where("problem_id = ?", problem.ID).Count(&count).Error; err != nil {
		t.Fatalf("failed to count solution keys: %v", err)
	}

	if count != 0 {
		t.Errorf("expected 0 solution keys after cascade delete, got %d", count)
	}
}

// setupTestDB creates a test database with migrations
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	if err := db.AutoMigrate(
		&models.Problem{},
		&models.TestCase{},
		&models.Submission{},
		&models.SolutionKey{},
	).Error; err != nil {
		t.Fatalf("failed to migrate test database: %v", err)
	}

	return db
}
