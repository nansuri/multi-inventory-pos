package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/username/multi-inventory-manager/internal/domain"
)

type InventoryHandler struct {
	usecase domain.InventoryUsecase
}

func NewInventoryHandler(u domain.InventoryUsecase) *InventoryHandler {
	return &InventoryHandler{usecase: u}
}

func (h *InventoryHandler) CreateItem(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}

	var item domain.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.TenantID = tenantID
	item.BranchID = branchID.(uint)
	if err := h.usecase.CreateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *InventoryHandler) ListItems(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}

	items, err := h.usecase.ListItems(branchID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *InventoryHandler) GetByBarcode(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}
	barcode := c.Param("barcode")

	item, err := h.usecase.GetItemByBarcode(barcode, branchID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *InventoryHandler) UpdateItem(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	var item domain.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = uint(id)
	item.TenantID = tenantID
	item.BranchID = branchID.(uint)
	if err := h.usecase.UpdateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item updated"})
}

func (h *InventoryHandler) DeleteItem(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	if err := h.usecase.DeleteItem(uint(id), branchID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
}

func (h *InventoryHandler) ToggleActive(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	if err := h.usecase.ToggleActive(uint(id), branchID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item status toggled"})
}

type updateStockRequest struct {
	Quantity float64 `json:"quantity" binding:"required"`
}

func (h *InventoryHandler) UpdateStock(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}
	itemIDStr := c.Param("id")
	itemID, _ := strconv.ParseUint(itemIDStr, 10, 32)

	var req updateStockRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.UpdateStock(uint(itemID), branchID.(uint), req.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "stock updated"})
}
