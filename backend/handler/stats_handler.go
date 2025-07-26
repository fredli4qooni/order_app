package handler

import (
	"github.com/gin-gonic/gin"
	"it-order-app/backend/repository"
	"net/http"
)

type StatsHandler struct {
	statsRepository repository.StatsRepository
}

func NewStatsHandler(statsRepo repository.StatsRepository) *StatsHandler {
	return &StatsHandler{statsRepo}
}

func (h *StatsHandler) GetStats(c *gin.Context) {
	stats, err := h.statsRepository.GetStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch statistics"})
		return
	}
	c.JSON(http.StatusOK, stats)
}
