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
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}

	var table domain.Table
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	table.TenantID = tenantID
	table.BranchID = branchID.(uint)
	if err := h.usecase.CreateTable(&table); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, table)
}

func (h *OrderHandler) ListTables(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}

	tables, err := h.usecase.ListTables(branchID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tables)
}

func (h *OrderHandler) UpdateTable(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))

	var table domain.Table
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	table.ID = uint(id)
	table.BranchID = branchID.(uint)
	if err := h.usecase.UpdateTable(&table); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, table)
}

func (h *OrderHandler) DeleteTable(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.usecase.DeleteTable(uint(id), branchID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Table deleted"})
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}

	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.TenantID = tenantID
	order.BranchID = branchID.(uint)
	if err := h.usecase.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}

	orders, err := h.usecase.ListOrders(branchID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) CompleteOrder(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}
	orderIDStr := c.Param("id")
	orderID, _ := strconv.ParseUint(orderIDStr, 10, 32)

	if err := h.usecase.CompleteOrder(uint(orderID), branchID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order completed and inventory updated"})
}
