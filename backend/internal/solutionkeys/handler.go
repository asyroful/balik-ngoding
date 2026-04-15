package solutionkeys

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for solution keys
type Handler struct {
	service *SolutionKeyService
}

// NewHandler creates a new Handler
func NewHandler(service *SolutionKeyService) *Handler {
	return &Handler{service: service}
}

// CreateSolutionKey handles POST /api/solution-keys
func (h *Handler) CreateSolutionKey(c *gin.Context) {
	var req CreateSolutionKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	sk, err := h.service.CreateSolutionKey(req)
	if err != nil {
		if err.Error() == "problem not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Problem not found"})
			return
		}
		if err.Error() == "solution key already exists for this problem and language" {
			c.JSON(http.StatusConflict, gin.H{"error": "Solution key already exists for this problem and language"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sk)
}

// UpdateSolutionKey handles PUT /api/solution-keys/:id
func (h *Handler) UpdateSolutionKey(c *gin.Context) {
	id := c.Param("id")

	var req UpdateSolutionKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	sk, err := h.service.UpdateSolutionKey(id, req)
	if err != nil {
		if err.Error() == "solution key not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Solution key not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sk)
}

// GetSolutionKey handles GET /api/solution-keys/:problemId
func (h *Handler) GetSolutionKey(c *gin.Context) {
	problemID := c.Param("problemId")

	sk, err := h.service.GetSolutionKeyByProblemID(problemID)
	if err != nil {
		if err.Error() == "solution key not found" {
			c.JSON(http.StatusOK, gin.H{"data": nil})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sk})
}

// DeleteSolutionKey handles DELETE /admin/solution-keys/:id
func (h *Handler) DeleteSolutionKey(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteSolutionKey(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Solution key deleted"})
}
