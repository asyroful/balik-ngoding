package solutionkeys

import (
	"testing"
	"time"

	"balik-ngoding-backend/internal/models"

	"github.com/flyingmutant/rapid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB creates an in-memory SQLite database for testing
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	// Run migrations
	if err := db.AutoMigrate(&models.Problem{}, &models.SolutionKey{}); err != nil {
		t.Fatalf("Failed to run migrations: %v", err)
	}

	return db
}

// TestCreateSolutionKeyValid tests creating a solution key with valid data
func TestCreateSolutionKeyValid(t *testing.T) {
	db := setupTestDB(t)
	service := NewSolutionKeyService(db)

	// Create a test problem first
	problem := &models.Problem{
		ID:    "test-problem-1",
		Title: "Test Problem",
	}
	db.Create(problem)

	req := CreateSolutionKeyRequest{
		ProblemID: "test-problem-1",
		Code:      "function test() { return 42; }",
		Language:  "javascript",
	}

	sk, err := service.CreateSolutionKey(req)
	if err != nil {
		t.Errorf("CreateSolutionKey failed: %v", err)
	}

	if sk == nil {
		t.Errorf("Expected solution key, got nil")
	}

	if sk.ProblemID != "test-problem-1" {
		t.Errorf("Expected problemId test-problem-1, got %s", sk.ProblemID)
	}

	if sk.Code != "function test() { return 42; }" {
		t.Errorf("Expected code to match, got %s", sk.Code)
	}

	if sk.Language != "javascript" {
		t.Errorf("Expected language javascript, got %s", sk.Language)
	}
}

// TestCreateSolutionKeyMissingFields tests creating with missing fields
func TestCreateSolutionKeyMissingFields(t *testing.T) {
	db := setupTestDB(t)
	service := NewSolutionKeyService(db)

	tests := []struct {
		name    string
		req     CreateSolutionKeyRequest
		wantErr bool
	}{
		{
			name: "missing problemId",
			req: CreateSolutionKeyRequest{
				Code:     "function test() {}",
				Language: "javascript",
			},
			wantErr: true,
		},
		{
			name: "missing code",
			req: CreateSolutionKeyRequest{
				ProblemID: "test-1",
				Language:  "javascript",
			},
			wantErr: true,
		},
		{
			name: "missing language",
			req: CreateSolutionKeyRequest{
				ProblemID: "test-1",
				Code:      "function test() {}",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.CreateSolutionKey(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSolutionKey error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestCreateSolutionKeyInvalidLanguage tests creating with invalid language
func TestCreateSolutionKeyInvalidLanguage(t *testing.T) {
	db := setupTestDB(t)
	service := NewSolutionKeyService(db)

	problem := &models.Problem{
		ID:    "test-problem-1",
		Title: "Test Problem",
	}
	db.Create(problem)

	req := CreateSolutionKeyRequest{
		ProblemID: "test-problem-1",
		Code:      "function test() {}",
		Language:  "python",
	}

	_, err := service.CreateSolutionKey(req)
	if err == nil {
		t.Errorf("Expected error for invalid language, got nil")
	}
}

// TestCreateSolutionKeyNonExistentProblem tests creating with non-existent problem
func TestCreateSolutionKeyNonExistentProblem(t *testing.T) {
	db := setupTestDB(t)
	service := NewSolutionKeyService(db)

	req := CreateSolutionKeyRequest{
		ProblemID: "non-existent",
		Code:      "function test() {}",
		Language:  "javascript",
	}

	_, err := service.CreateSolutionKey(req)
	if err == nil {
		t.Errorf("Expected error for non-existent problem, got nil")
	}

	if err.Error() != "problem not found" {
		t.Errorf("Expected 'problem not found', got %v", err)
	}
}

// TestUpdateSolutionKeyPreservesCreatedAt tests that update preserves createdAt
func TestUpdateSolutionKeyPreservesCreatedAt(t *testing.T) {
	db := setupTestDB(t)
	service := NewSolutionKeyService(db)

	// Create problem and solution key
	problem := &models.Problem{
		ID:    "test-problem-1",
		Title: "Test Problem",
	}
	db.Create(problem)

	originalTime := time.Now().Add(-1 * time.Hour)
	sk := &models.SolutionKey{
		ID:        "test-sk-1",
		ProblemID: "test-problem-1",
		Code:      "function test() { return 1; }",
		Language:  "javascript",
		CreatedAt: originalTime,
	}
	db.Create(sk)

	// Update solution key
	updateReq := UpdateSolutionKeyRequest{
		Code:     "function test() { return 2; }",
		Language: "javascript",
	}

	updated, err := service.UpdateSolutionKey("test-sk-1", updateReq)
	if err != nil {
		t.Errorf("UpdateSolutionKey failed: %v", err)
	}

	if updated.CreatedAt != originalTime {
		t.Errorf("CreatedAt changed after update: %v != %v", updated.CreatedAt, originalTime)
	}
}

// TestUpdateSolutionKeyModifiesUpdatedAt tests that update modifies updatedAt
func TestUpdateSolutionKeyModifiesUpdatedAt(t *testing.T) {
	db := setupTestDB(t)
	service := NewSolutionKeyService(db)

	// Create problem and solution key
	problem := &models.Problem{
		ID:    "test-problem-1",
		Title: "Test Problem",
	}
	db.Create(problem)

	sk := &models.SolutionKey{
		ID:        "test-sk-1",
		ProblemID: "test-problem-1",
		Code:      "function test() { return 1; }",
		Language:  "javascript",
		UpdatedAt: time.Now().Add(-1 * time.Hour),
	}
	db.Create(sk)

	originalUpdatedAt := sk.UpdatedAt

	// Update solution key
	updateReq := UpdateSolutionKeyRequest{
		Code:     "function test() { return 2; }",
		Language: "javascript",
	}

	updated, err := service.UpdateSolutionKey("test-sk-1", updateReq)
	if err != nil {
		t.Errorf("UpdateSolutionKey failed: %v", err)
	}

	if updated.UpdatedAt.Before(originalUpdatedAt) {
		t.Errorf("UpdatedAt not modified: %v <= %v", updated.UpdatedAt, originalUpdatedAt)
	}
}

// TestGetSolutionKeyByProblemID tests retrieving by problemId
func TestGetSolutionKeyByProblemID(t *testing.T) {
	db := setupTestDB(t)
	service := NewSolutionKeyService(db)

	// Create problem and solution key
	problem := &models.Problem{
		ID:    "test-problem-1",
		Title: "Test Problem",
	}
	db.Create(problem)

	sk := &models.SolutionKey{
		ID:        "test-sk-1",
		ProblemID: "test-problem-1",
		Code:      "function test() { return 42; }",
		Language:  "javascript",
	}
	db.Create(sk)

	retrieved, err := service.GetSolutionKeyByProblemID("test-problem-1")
	if err != nil {
		t.Errorf("GetSolutionKeyByProblemID failed: %v", err)
	}

	if retrieved.ID != "test-sk-1" {
		t.Errorf("Expected ID test-sk-1, got %s", retrieved.ID)
	}

	if retrieved.Code != "function test() { return 42; }" {
		t.Errorf("Expected code to match, got %s", retrieved.Code)
	}
}

// TestDeleteSolutionKey tests deleting a solution key
func TestDeleteSolutionKey(t *testing.T) {
	db := setupTestDB(t)
	service := NewSolutionKeyService(db)

	// Create problem and solution key
	problem := &models.Problem{
		ID:    "test-problem-1",
		Title: "Test Problem",
	}
	db.Create(problem)

	sk := &models.SolutionKey{
		ID:        "test-sk-1",
		ProblemID: "test-problem-1",
		Code:      "function test() {}",
		Language:  "javascript",
	}
	db.Create(sk)

	// Delete
	err := service.DeleteSolutionKey("test-sk-1")
	if err != nil {
		t.Errorf("DeleteSolutionKey failed: %v", err)
	}

	// Verify deletion
	_, err = service.GetSolutionKeyByProblemID("test-problem-1")
	if err == nil {
		t.Errorf("Expected error after deletion, got nil")
	}
}

// TestSolutionKeyRoundTripProperty tests round-trip storage and retrieval
// Feature: evaluator-fix-and-solution-keys, Property 8: solution key round-trip
func TestSolutionKeyRoundTripProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		db := setupTestDB(t)
		service := NewSolutionKeyService(db)

		code := rapid.String().Draw(t, "code")
		language := rapid.SampledFrom([]string{"javascript", "sql"}).Draw(t, "language")
		problemID := "test-problem-" + rapid.String().Draw(t, "problemID")

		// Create problem
		problem := &models.Problem{
			ID:    problemID,
			Title: "Test Problem",
		}
		db.Create(problem)

		// Create solution key
		req := CreateSolutionKeyRequest{
			ProblemID: problemID,
			Code:      code,
			Language:  language,
		}

		sk, err := service.CreateSolutionKey(req)
		if err != nil {
			t.Fatalf("CreateSolutionKey failed: %v", err)
		}

		// Retrieve and verify
		retrieved, err := service.GetSolutionKeyByProblemID(problemID)
		if err != nil {
			t.Fatalf("GetSolutionKeyByProblemID failed: %v", err)
		}

		if retrieved.Code != code {
			t.Fatalf("Code mismatch: %q != %q", retrieved.Code, code)
		}

		if retrieved.Language != language {
			t.Fatalf("Language mismatch: %q != %q", retrieved.Language, language)
		}

		if retrieved.ID != sk.ID {
			t.Fatalf("ID mismatch: %q != %q", retrieved.ID, sk.ID)
		}
	})
}

// TestSolutionKeyIsolationProperty tests that multiple keys are isolated
// Feature: evaluator-fix-and-solution-keys, Property 9: solution key isolation
func TestSolutionKeyIsolationProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		db := setupTestDB(t)
		service := NewSolutionKeyService(db)

		// Create multiple problems and solution keys
		for i := 0; i < 3; i++ {
			problemID := "problem-" + rapid.String().Draw(t, "id")
			code := rapid.String().Draw(t, "code")

			problem := &models.Problem{
				ID:    problemID,
				Title: "Problem " + problemID,
			}
			db.Create(problem)

			req := CreateSolutionKeyRequest{
				ProblemID: problemID,
				Code:      code,
				Language:  "javascript",
			}

			service.CreateSolutionKey(req)
		}

		// Verify each retrieval returns correct key
		var problems []models.Problem
		db.Find(&problems)

		for _, p := range problems {
			retrieved, err := service.GetSolutionKeyByProblemID(p.ID)
			if err != nil {
				t.Fatalf("Failed to retrieve solution key for problem %s: %v", p.ID, err)
			}

			if retrieved.ProblemID != p.ID {
				t.Fatalf("Retrieved wrong solution key: expected problemId %s, got %s", p.ID, retrieved.ProblemID)
			}
		}
	})
}

// TestSolutionKeyUniquenessProperty tests uniqueness constraint
// Feature: evaluator-fix-and-solution-keys, Property 10: solution key uniqueness
func TestSolutionKeyUniquenessProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		db := setupTestDB(t)
		service := NewSolutionKeyService(db)

		problemID := "test-problem-" + rapid.String().Draw(t, "problemID")
		language := rapid.SampledFrom([]string{"javascript", "sql"}).Draw(t, "language")

		// Create problem
		problem := &models.Problem{
			ID:    problemID,
			Title: "Test Problem",
		}
		db.Create(problem)

		// Create first solution key
		req1 := CreateSolutionKeyRequest{
			ProblemID: problemID,
			Code:      "function test1() {}",
			Language:  language,
		}

		_, err := service.CreateSolutionKey(req1)
		if err != nil {
			t.Fatalf("First CreateSolutionKey failed: %v", err)
		}

		// Try to create duplicate - should fail
		req2 := CreateSolutionKeyRequest{
			ProblemID: problemID,
			Code:      "function test2() {}",
			Language:  language,
		}

		_, err = service.CreateSolutionKey(req2)
		if err == nil {
			t.Fatalf("Expected error for duplicate solution key, got nil")
		}
	})
}

// TestCascadeDeleteProperty tests cascade delete behavior
// Feature: evaluator-fix-and-solution-keys, Property 11: cascade delete behavior
func TestCascadeDeleteProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		db := setupTestDB(t)
		service := NewSolutionKeyService(db)

		problemID := "test-problem-" + rapid.String().Draw(t, "problemID")

		// Create problem
		problem := &models.Problem{
			ID:    problemID,
			Title: "Test Problem",
		}
		db.Create(problem)

		// Create solution key
		req := CreateSolutionKeyRequest{
			ProblemID: problemID,
			Code:      "function test() {}",
			Language:  "javascript",
		}

		_, err := service.CreateSolutionKey(req)
		if err != nil {
			t.Fatalf("CreateSolutionKey failed: %v", err)
		}

		// Delete problem
		db.Delete(&models.Problem{}, "id = ?", problemID)

		// Verify solution key is also deleted
		_, err = service.GetSolutionKeyByProblemID(problemID)
		if err == nil {
			t.Fatalf("Expected error after cascade delete, got nil")
		}
	})
}

// TestTimestampPreservationProperty tests timestamp preservation on update
// Feature: evaluator-fix-and-solution-keys, Property 12: timestamp preservation on update
func TestTimestampPreservationProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		db := setupTestDB(t)
		service := NewSolutionKeyService(db)

		problemID := "test-problem-" + rapid.String().Draw(t, "problemID")

		// Create problem
		problem := &models.Problem{
			ID:    problemID,
			Title: "Test Problem",
		}
		db.Create(problem)

		// Create solution key
		req := CreateSolutionKeyRequest{
			ProblemID: problemID,
			Code:      "function test() { return 1; }",
			Language:  "javascript",
		}

		sk, err := service.CreateSolutionKey(req)
		if err != nil {
			t.Fatalf("CreateSolutionKey failed: %v", err)
		}

		originalCreatedAt := sk.CreatedAt

		// Update solution key
		updateReq := UpdateSolutionKeyRequest{
			Code:     "function test() { return 2; }",
			Language: "javascript",
		}

		updated, err := service.UpdateSolutionKey(sk.ID, updateReq)
		if err != nil {
			t.Fatalf("UpdateSolutionKey failed: %v", err)
		}

		// Verify createdAt is preserved
		if updated.CreatedAt != originalCreatedAt {
			t.Fatalf("CreatedAt changed: %v != %v", updated.CreatedAt, originalCreatedAt)
		}

		// Verify updatedAt is modified
		if updated.UpdatedAt.Before(originalCreatedAt) {
			t.Fatalf("UpdatedAt not modified properly")
		}
	})
}
