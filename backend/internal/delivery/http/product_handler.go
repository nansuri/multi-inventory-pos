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

	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.TenantID = tenantID
	if err := h.usecase.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	products, err := h.usecase.ListProducts(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

type setRecipeRequest struct {
	Recipes []domain.Recipe `json:"recipes" binding:"required"`
}

func (h *ProductHandler) SetRecipe(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)
	productIDStr := c.Param("id")
	productID, _ := strconv.ParseUint(productIDStr, 10, 32)

	var req setRecipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.SetRecipe(uint(productID), tenantID, req.Recipes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "recipe updated"})
}

type prepareRequest struct {
	Pax int `json:"pax" binding:"required,min=1"`
}

func (h *ProductHandler) Prepare(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)
	productIDStr := c.Param("id")
	productID, _ := strconv.ParseUint(productIDStr, 10, 32)

	var req prepareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.PrepareProduct(uint(productID), tenantID, req.Pax); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "preparation recorded, inventory updated"})
}

func (h *ProductHandler) ListPreparations(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	logs, err := h.usecase.ListPreparationLogs(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}
