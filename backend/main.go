package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"balik-ngoding-backend/internal/database"
	"balik-ngoding-backend/internal/problems"
	"balik-ngoding-backend/internal/submissions"

	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		frontendOrigin := os.Getenv("FRONTEND_ORIGIN")
		if frontendOrigin == "" {
			frontendOrigin = "http://localhost:3000"
		}

		origin := c.Request.Header.Get("Origin")
		// Allow both http and https versions of the frontend origin
		httpsOrigin := strings.Replace(frontendOrigin, "http://", "https://", 1)
		httpOrigin := strings.Replace(frontendOrigin, "https://", "http://", 1)

		allowedOrigin := ""
		if origin == frontendOrigin || origin == httpsOrigin || origin == httpOrigin {
			allowedOrigin = origin
		}

		if allowedOrigin != "" {
			c.Header("Access-Control-Allow-Origin", allowedOrigin)
		}
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(corsMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Problems routes
	problemsHandler := problems.NewHandler()
	r.GET("/problems/summary", problemsHandler.GetSummary)
	r.GET("/problems", problemsHandler.GetProblems)
	r.GET("/problems/:id", problemsHandler.GetProblemByID)

	// Submissions routes
	submissionsHandler := submissions.NewHandler()
	r.POST("/submit", submissionsHandler.Submit)

	return r
}

func main() {
	database.Init()
	if err := database.SeedFromJSON(database.DB); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	r := setupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
