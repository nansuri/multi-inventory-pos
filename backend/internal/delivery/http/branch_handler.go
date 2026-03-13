package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/username/multi-inventory-manager/internal/domain"
)

type BranchHandler struct {
	usecase domain.BranchUsecase
}

func NewBranchHandler(u domain.BranchUsecase) *BranchHandler {
	return &BranchHandler{usecase: u}
}

func (h *BranchHandler) Create(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	var branch domain.Branch
	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branch.TenantID = tenantID
	if err := h.usecase.CreateBranch(&branch); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, branch)
}

func (h *BranchHandler) List(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	branches, err := h.usecase.ListBranches(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, branches)
}

func (h *BranchHandler) Update(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)
	id, _ := strconv.Atoi(c.Param("id"))

	var branch domain.Branch
	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branch.ID = uint(id)
	branch.TenantID = tenantID
	if err := h.usecase.UpdateBranch(&branch); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, branch)
}

func (h *BranchHandler) Delete(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.usecase.DeleteBranch(uint(id), tenantID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Branch deleted"})
}
