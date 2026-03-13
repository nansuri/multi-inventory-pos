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
	
	var branchID uint
	if bid, exists := c.Get("branch_id"); exists {
		branchID = bid.(uint)
	}

	summary, err := h.usecase.GetDashboardSummary(tenantID, branchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summary)
}

func (h *DashboardHandler) GetOrderHistory(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)
	
	var branchID uint
	if bid, exists := c.Get("branch_id"); exists {
		branchID = bid.(uint)
	}
	
	period := c.DefaultQuery("period", "day")

	orders, totalRevenue, orderCount, err := h.usecase.GetOrderHistory(tenantID, branchID, period)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders":        orders,
		"total_revenue": totalRevenue,
		"order_count":   orderCount,
	})
}
