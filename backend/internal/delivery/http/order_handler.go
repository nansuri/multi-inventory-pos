package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/username/multi-inventory-manager/internal/domain"
)

type OrderHandler struct {
	usecase domain.OrderUsecase
}

func NewOrderHandler(u domain.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase: u}
}

func (h *OrderHandler) CreateTable(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	var table domain.Table
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	table.TenantID = tenantID
	if err := h.usecase.CreateTable(&table); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, table)
}

func (h *OrderHandler) ListTables(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	tables, err := h.usecase.ListTables(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tables)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.TenantID = tenantID
	// In a real POS, you'd calculate price on backend too
	if err := h.usecase.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	orders, err := h.usecase.ListOrders(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) CompleteOrder(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)
	orderIDStr := c.Param("id")
	orderID, _ := strconv.ParseUint(orderIDStr, 10, 32)

	if err := h.usecase.CompleteOrder(uint(orderID), tenantID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order completed and inventory updated"})
}
