package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flyingmutant/rapid"
	"github.com/gin-gonic/gin"
)

// TestAdminAuthMiddlewareValidCredentials tests that valid credentials pass through
func TestAdminAuthMiddlewareValidCredentials(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(AdminAuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Create valid auth header
	authHeader := EncodeBasicAuth("siful", "atmin162")

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", authHeader)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

// TestAdminAuthMiddlewareInvalidCredentials tests that invalid credentials return 401
func TestAdminAuthMiddlewareInvalidCredentials(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(AdminAuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Create invalid auth header
	authHeader := EncodeBasicAuth("siful", "wrongpassword")

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", authHeader)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

// TestAdminAuthMiddlewareMissingCredentials tests that missing credentials return 401
func TestAdminAuthMiddlewareMissingCredentials(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(AdminAuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	// No Authorization header

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

// TestAdminAuthSecurityProperty tests that any request without valid credentials returns 401
// Feature: evaluator-fix-and-solution-keys, Property 13: admin authentication security
func TestAdminAuthSecurityProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		gin.SetMode(gin.TestMode)
		router := gin.New()
		router.Use(AdminAuthMiddleware())
		router.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		// Generate random credentials
		username := rapid.String().Draw(t, "username")
		password := rapid.String().Draw(t, "password")

		// Skip if it happens to be the correct credentials
		if username == "siful" && password == "atmin162" {
			return
		}

		// Create auth header with random credentials
		credentials := fmt.Sprintf("%s:%s", username, password)
		encoded := base64.StdEncoding.EncodeToString([]byte(credentials))
		authHeader := fmt.Sprintf("Basic %s", encoded)

		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", authHeader)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("Expected status 401 for invalid credentials, got %d", w.Code)
		}
	})
}

// TestAdminAuthMiddlewareInvalidFormat tests that invalid auth format returns 401
func TestAdminAuthMiddlewareInvalidFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(AdminAuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	tests := []struct {
		name      string
		authValue string
	}{
		{
			name:      "invalid base64",
			authValue: "Basic !!!invalid!!!",
		},
		{
			name:      "missing Basic prefix",
			authValue: base64.StdEncoding.EncodeToString([]byte("siful:atmin162")),
		},
		{
			name:      "missing colon separator",
			authValue: "Basic " + base64.StdEncoding.EncodeToString([]byte("sifulatmin162")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/test", nil)
			req.Header.Set("Authorization", tt.authValue)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != http.StatusUnauthorized {
				t.Errorf("Expected status 401, got %d", w.Code)
			}
		})
	}
}
