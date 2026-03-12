package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/username/multi-inventory-manager/internal/domain"
)

type DashboardHandler struct {
	usecase domain.ReportUsecase
}

func NewDashboardHandler(u domain.ReportUsecase) *DashboardHandler {
	return &DashboardHandler{usecase: u}
}

func (h *DashboardHandler) GetSummary(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	summary, err := h.usecase.GetDashboardSummary(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summary)
}
