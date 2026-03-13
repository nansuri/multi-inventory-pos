package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/username/multi-inventory-manager/internal/domain"
)

type TenantHandler struct {
	usecase domain.TenantUsecase
}

func NewTenantHandler(u domain.TenantUsecase) *TenantHandler {
	return &TenantHandler{usecase: u}
}

func (h *TenantHandler) GetTenant(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	tenant, err := h.usecase.GetTenant(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tenant)
}

type updateTenantRequest struct {
	TenantName string `json:"tenant_name" binding:"required"`
	Currency   string `json:"currency" binding:"required"`
	Language   string `json:"language" binding:"required"`
}

func (h *TenantHandler) UpdateTenant(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	var req updateTenantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.UpdateTenant(tenantID, req.TenantName, req.Currency, req.Language); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tenant settings updated"})
}
