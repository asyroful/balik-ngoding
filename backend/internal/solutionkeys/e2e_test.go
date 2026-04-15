package solutionkeys

import (
	"balik-ngoding-backend/internal/auth"
	"balik-ngoding-backend/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestSolutionKeyAPIE2E verifies complete API workflow for solution key management
// Feature: evaluator-fix-and-solution-keys, Property 8: solution key round-trip
func TestSolutionKeyAPIE2E(t *testing.T) {
	db := setupTestDB(t)
	handler := NewHandler(db)

	// Create test problem
	problem := &models.Problem{
		ID:       "test-api-problem",
		Title:    "API Test Problem",
		Category: "loop",
	}
	if err := db.Create(problem).Error; err != nil {
		t.Fatalf("failed to create test problem: %v", err)
	}

	// Test 1: Create solution key via API
	t.Run("create solution key", func(t *testing.T) {
		createReq := CreateSolutionKeyRequest{
			ProblemID: problem.ID,
			Code:      "function solution(n) { return n * 2; }",
			Language:  "javascript",
		}

		body, _ := json.Marshal(createReq)
		req := httptest.NewRequest("POST", "/api/solution-keys", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", createAuthHeader())

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		handler.CreateSolutionKey(c)

		if w.Code != http.StatusOK && w.Code != http.StatusCreated {
			t.Errorf("expected status 200 or 201, got %d", w.Code)
		}

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		data, ok := response["data"].(map[string]interface{})
		if !ok {
			t.Fatal("expected data in response")
		}

		if data["code"] != createReq.Code {
			t.Errorf("expected code %q, got %q", createReq.Code, data["code"])
		}
	})

	// Test 2: Retrieve solution key via API
	t.Run("retrieve solution key", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/solution-keys/"+problem.ID, nil)
		req.Header.Set("Authorization", createAuthHeader())

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Params = append(c.Params, gin.Param{Key: "problemId", Value: problem.ID})

		handler.GetSolutionKey(c)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		data, ok := response["data"].(map[string]interface{})
		if !ok {
			t.Fatal("expected data in response")
		}

		if data["problemId"] != problem.ID {
			t.Errorf("expected problemId %q, got %q", problem.ID, data["problemId"])
		}
	})

	// Test 3: Update solution key via API
	t.Run("update solution key", func(t *testing.T) {
		// First get the solution key ID
		var sk models.SolutionKey
		if err := db.Where("problem_id = ?", problem.ID).First(&sk).Error; err != nil {
			t.Fatalf("failed to get solution key: %v", err)
		}

		updateReq := UpdateSolutionKeyRequest{
			Code: "function solution(n) { return n * 3; }",
		}

		body, _ := json.Marshal(updateReq)
		req := httptest.NewRequest("PUT", "/api/solution-keys/"+sk.ID, bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", createAuthHeader())

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Params = append(c.Params, gin.Param{Key: "id", Value: sk.ID})

		handler.UpdateSolutionKey(c)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		data, ok := response["data"].(map[string]interface{})
		if !ok {
			t.Fatal("expected data in response")
		}

		if data["code"] != updateReq.Code {
			t.Errorf("expected updated code %q, got %q", updateReq.Code, data["code"])
		}
	})

	// Test 4: Delete solution key via API
	t.Run("delete solution key", func(t *testing.T) {
		// First get the solution key ID
		var sk models.SolutionKey
		if err := db.Where("problem_id = ?", problem.ID).First(&sk).Error; err != nil {
			t.Fatalf("failed to get solution key: %v", err)
		}

		req := httptest.NewRequest("DELETE", "/admin/solution-keys/"+sk.ID, nil)
		req.Header.Set("Authorization", createAuthHeader())

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Params = append(c.Params, gin.Param{Key: "id", Value: sk.ID})

		handler.DeleteSolutionKey(c)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		// Verify deletion
		var deleted models.SolutionKey
		result := db.First(&deleted, "id = ?", sk.ID)
		if result.Error != gorm.ErrRecordNotFound {
			t.Errorf("expected record not found, got error: %v", result.Error)
		}
	})
}

// TestUnauthorizedAccessE2E verifies that unauthenticated users cannot access solution keys
// Feature: evaluator-fix-and-solution-keys, Property 13: admin authentication security
func TestUnauthorizedAccessE2E(t *testing.T) {
	db := setupTestDB(t)
	handler := NewHandler(db)

	// Create test problem
	problem := &models.Problem{
		ID:       "test-unauth-problem",
		Title:    "Unauthorized Test",
		Category: "loop",
	}
	if err := db.Create(problem).Error; err != nil {
		t.Fatalf("failed to create test problem: %v", err)
	}

	// Create a solution key
	sk := &models.SolutionKey{
		ProblemID: problem.ID,
		Code:      "function solution() { return 1; }",
		Language:  "javascript",
	}
	if err := db.Create(sk).Error; err != nil {
		t.Fatalf("failed to create solution key: %v", err)
	}

	tests := []struct {
		name           string
		method         string
		path           string
		paramKey       string
		paramValue     string
		body           interface{}
		expectedStatus int
	}{
		{
			name:           "get without auth",
			method:         "GET",
			path:           "/api/solution-keys/" + problem.ID,
			paramKey:       "problemId",
			paramValue:     problem.ID,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "create without auth",
			method:         "POST",
			path:           "/api/solution-keys",
			expectedStatus: http.StatusUnauthorized,
			body: CreateSolutionKeyRequest{
				ProblemID: problem.ID,
				Code:      "function test() {}",
				Language:  "javascript",
			},
		},
		{
			name:           "update without auth",
			method:         "PUT",
			path:           "/api/solution-keys/" + sk.ID,
			paramKey:       "id",
			paramValue:     sk.ID,
			expectedStatus: http.StatusUnauthorized,
			body: UpdateSolutionKeyRequest{
				Code: "function updated() {}",
			},
		},
		{
			name:           "delete without auth",
			method:         "DELETE",
			path:           "/admin/solution-keys/" + sk.ID,
			paramKey:       "id",
			paramValue:     sk.ID,
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.body != nil {
				body, _ := json.Marshal(tt.body)
				req = httptest.NewRequest(tt.method, tt.path, bytes.NewReader(body))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req = httptest.NewRequest(tt.method, tt.path, nil)
			}

			// Don't set Authorization header - test unauthorized access
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			if tt.paramKey != "" && tt.paramValue != "" {
				c.Params = append(c.Params, gin.Param{Key: tt.paramKey, Value: tt.paramValue})
			}

			// Call the appropriate handler
			switch tt.method {
			case "GET":
				handler.GetSolutionKey(c)
			case "POST":
				handler.CreateSolutionKey(c)
			case "PUT":
				handler.UpdateSolutionKey(c)
			case "DELETE":
				handler.DeleteSolutionKey(c)
			}

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			// Verify solution key is not exposed in response
			if w.Code == http.StatusUnauthorized {
				var response map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &response)
				if _, hasData := response["data"]; hasData {
					t.Error("solution key should not be exposed in unauthorized response")
				}
			}
		})
	}
}

// TestInvalidCredentialsE2E verifies that invalid credentials are rejected
// Feature: evaluator-fix-and-solution-keys, Property 13: admin authentication security
func TestInvalidCredentialsE2E(t *testing.T) {
	db := setupTestDB(t)
	handler := NewHandler(db)

	// Create test problem
	problem := &models.Problem{
		ID:       "test-invalid-creds",
		Title:    "Invalid Credentials Test",
		Category: "loop",
	}
	if err := db.Create(problem).Error; err != nil {
		t.Fatalf("failed to create test problem: %v", err)
	}

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
	}{
		{
			name:           "invalid username",
			authHeader:     auth.EncodeBasicAuth("wronguser", "atmin162"),
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "invalid password",
			authHeader:     auth.EncodeBasicAuth("siful", "wrongpass"),
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "malformed auth header",
			authHeader:     "Bearer invalid",
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/solution-keys/"+problem.ID, nil)
			req.Header.Set("Authorization", tt.authHeader)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Params = append(c.Params, gin.Param{Key: "problemId", Value: problem.ID})

			handler.GetSolutionKey(c)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

// setupTestDB creates a test database with migrations
func setupTestDB(t *testing.T) *gorm.DB {
	db := database.SetupTestDB()
	if err := db.AutoMigrate(
		&models.Problem{},
		&models.SolutionKey{},
	).Error; err != nil {
		t.Fatalf("failed to migrate test database: %v", err)
	}
	return db
}

// createAuthHeader creates a valid Basic Auth header for testing
func createAuthHeader() string {
	return auth.EncodeBasicAuth("siful", "atmin162")
}
