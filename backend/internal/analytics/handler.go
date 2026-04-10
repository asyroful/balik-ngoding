package analytics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler holds the service dependency for analytics routes.
type Handler struct {
	service *AnalyticsService
}

// NewHandler creates a new analytics Handler.
func NewHandler() *Handler {
	return &Handler{service: NewAnalyticsService()}
}

// GetStats handles GET /analytics/stats
func (h *Handler) GetStats(c *gin.Context) {
	stats, err := h.service.GetStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil statistik"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": stats})
}
