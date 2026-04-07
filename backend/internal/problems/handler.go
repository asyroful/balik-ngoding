package problems

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler holds the service dependency for problem-related routes.
type Handler struct {
	service *ProblemsService
}

// NewHandler creates a new Handler with a ProblemsService.
func NewHandler() *Handler {
	return &Handler{service: &ProblemsService{}}
}

// GetProblems handles GET /problems
// Supports optional query params: category, difficulty
func (h *Handler) GetProblems(c *gin.Context) {
	category := c.Query("category")
	difficulty := c.Query("difficulty")

	problems, err := h.service.FindAll(category, difficulty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    "Gagal mengambil daftar soal",
			"error":      "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": problems})
}

// GetProblemByID handles GET /problems/:id
func (h *Handler) GetProblemByID(c *gin.Context) {
	id := c.Param("id")

	problem, err := h.service.FindOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    "Gagal mengambil detail soal",
			"error":      "Internal Server Error",
		})
		return
	}

	if problem == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "Soal tidak ditemukan",
			"error":      "Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": problem})
}
