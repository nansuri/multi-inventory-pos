package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/username/multi-inventory-manager/internal/domain"
)

type ProductHandler struct {
	usecase domain.ProductUsecase
}

func NewProductHandler(u domain.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: u}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}

	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.TenantID = tenantID
	product.BranchID = branchID.(uint)
	if err := h.usecase.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}

	products, err := h.usecase.ListProducts(branchID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = uint(id)
	product.TenantID = tenantID
	product.BranchID = branchID.(uint)
	if _, err := h.usecase.GetProductByID(product.ID, branchID.(uint)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	if err := h.usecase.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated"})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	if err := h.usecase.DeleteProduct(uint(id), branchID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}

type setRecipeRequest struct {
	Recipes []domain.Recipe `json:"recipes" binding:"required"`
}

func (h *ProductHandler) SetRecipe(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}
	productIDStr := c.Param("id")
	productID, _ := strconv.ParseUint(productIDStr, 10, 32)

	var req setRecipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.SetRecipe(uint(productID), branchID.(uint), req.Recipes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "recipe updated"})
}

type prepareRequest struct {
	Pax int `json:"pax" binding:"required,min=1"`
}

func (h *ProductHandler) Prepare(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}
	productIDStr := c.Param("id")
	productID, _ := strconv.ParseUint(productIDStr, 10, 32)

	var req prepareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.PrepareProduct(uint(productID), branchID.(uint), req.Pax); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "preparation recorded, inventory updated"})
}

func (h *ProductHandler) ListPreparations(c *gin.Context) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch context required"})
		return
	}

	logs, err := h.usecase.ListProductionLogs(branchID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}
