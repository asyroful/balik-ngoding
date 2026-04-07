package submissions

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// submitRequestBody is the JSON body for POST /submit.
type submitRequestBody struct {
	ProblemID string `json:"problemId" binding:"required"`
	Code      string `json:"code"`
	Language  string `json:"language"`
}

// Handler holds the service dependency for submission-related routes.
type Handler struct {
	service *SubmissionsService
}

// NewHandler creates a new submissions Handler.
func NewHandler() *Handler {
	return &Handler{service: NewSubmissionsService()}
}

// Submit handles POST /submit
func (h *Handler) Submit(c *gin.Context) {
	var body submitRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "problemId wajib diisi",
			"error":      "Bad Request",
		})
		return
	}

	// Validate code is non-empty
	if strings.TrimSpace(body.Code) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "code tidak boleh kosong",
			"error":      "Bad Request",
		})
		return
	}

	// Default language to "javascript"
	language := body.Language
	if strings.TrimSpace(language) == "" {
		language = "javascript"
	}

	req := SubmitRequest{
		ProblemID: body.ProblemID,
		Code:      body.Code,
		Language:  language,
	}

	result, err := h.service.Submit(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    "Gagal memproses submission",
			"error":      "Internal Server Error",
		})
		return
	}

	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "Soal tidak ditemukan",
			"error":      "Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
