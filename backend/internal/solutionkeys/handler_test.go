package solutionkeys

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"balik-ngoding-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestRouter creates a test router with solution key handler
func setupTestRouter(t *testing.T) (*gin.Engine, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	if err := db.AutoMigrate(&models.Problem{}, &models.SolutionKey{}); err != nil {
		t.Fatalf("Failed to run migrations: %v", err)
	}

	gin.SetMode(gin.TestMode)
	router := gin.New()

	service := NewSolutionKeyService(db)
	handler := NewHandler(service)

	router.POST("/api/solution-keys", handler.CreateSolutionKey)
	router.PUT("/api/solution-keys/:id", handler.UpdateSolutionKey)
	router.GET("/api/solution-keys/:problemId", handler.GetSolutionKey)
	router.DELETE("/admin/solution-keys/:id", handler.DeleteSolutionKey)

	return router, db
}

// TestCreateSolutionKeyEndpoint tests POST endpoint
func TestCreateSolutionKeyEndpoint(t *testing.T) {
	router, db := setupTestRouter(t)

	// Create a test problem
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

	body, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest("POST", "/api/solution-keys", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	var sk models.SolutionKey
	json.Unmarshal(w.Body.Bytes(), &sk)

	if sk.ProblemID != "test-problem-1" {
		t.Errorf("Expected problemId test-problem-1, got %s", sk.ProblemID)
	}
}

// TestUpdateSolutionKeyEndpoint tests PUT endpoint
func TestUpdateSolutionKeyEndpoint(t *testing.T) {
	router, db := setupTestRouter(t)

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
	}
	db.Create(sk)

	// Update
	updateReq := UpdateSolutionKeyRequest{
		Code:     "function test() { return 2; }",
		Language: "javascript",
	}

	body, _ := json.Marshal(updateReq)
	httpReq, _ := http.NewRequest("PUT", "/api/solution-keys/test-sk-1", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var updated models.SolutionKey
	json.Unmarshal(w.Body.Bytes(), &updated)

	if updated.Code != "function test() { return 2; }" {
		t.Errorf("Expected updated code, got %s", updated.Code)
	}
}

// TestGetSolutionKeyEndpoint tests GET endpoint
func TestGetSolutionKeyEndpoint(t *testing.T) {
	router, db := setupTestRouter(t)

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

	httpReq, _ := http.NewRequest("GET", "/api/solution-keys/test-problem-1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var retrieved models.SolutionKey
	json.Unmarshal(w.Body.Bytes(), &retrieved)

	if retrieved.ID != "test-sk-1" {
		t.Errorf("Expected ID test-sk-1, got %s", retrieved.ID)
	}
}

// TestDeleteSolutionKeyEndpoint tests DELETE endpoint
func TestDeleteSolutionKeyEndpoint(t *testing.T) {
	router, db := setupTestRouter(t)

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

	httpReq, _ := http.NewRequest("DELETE", "/admin/solution-keys/test-sk-1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Verify deletion
	var count int64
	db.Model(&models.SolutionKey{}).Where("id = ?", "test-sk-1").Count(&count)
	if count != 0 {
		t.Errorf("Expected solution key to be deleted, but found %d records", count)
	}
}

// TestLanguageValidationEndpoint tests language validation
// Feature: evaluator-fix-and-solution-keys, Property 14: language validation
func TestLanguageValidationEndpoint(t *testing.T) {
	router, db := setupTestRouter(t)

	// Create a test problem
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

	body, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest("POST", "/api/solution-keys", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}
