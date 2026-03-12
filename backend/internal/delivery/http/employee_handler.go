package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/username/multi-inventory-manager/internal/domain"
)

type EmployeeHandler struct {
	usecase domain.EmployeeUsecase
}

func NewEmployeeHandler(u domain.EmployeeUsecase) *EmployeeHandler {
	return &EmployeeHandler{usecase: u}
}

func (h *EmployeeHandler) CreateRole(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	var role domain.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role.TenantID = tenantID
	if err := h.usecase.CreateRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

func (h *EmployeeHandler) ListRoles(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	roles, err := h.usecase.ListRoles(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func (h *EmployeeHandler) UpdateRole(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	var role domain.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role.ID = uint(id)
	role.TenantID = tenantID
	if err := h.usecase.UpdateRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "role updated"})
}

func (h *EmployeeHandler) DeleteRole(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	if err := h.usecase.DeleteRole(uint(id), tenantID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "role deleted"})
}

type createEmployeeRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   uint   `json:"role_id" binding:"required"`
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	var req createEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.CreateEmployee(req.Username, req.Password, tenantID, req.RoleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "employee created"})
}

func (h *EmployeeHandler) ListEmployees(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	employees, err := h.usecase.ListEmployees(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	if err := h.usecase.DeleteEmployee(uint(id), 0); err != nil { // Note: 0 tenantID because repository handles it via tenant_id check
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "employee deleted"})
}
