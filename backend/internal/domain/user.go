package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	TenantID     uint           `gorm:"not null" json:"tenant_id"`
	Tenant       Tenant         `gorm:"foreignKey:TenantID" json:"tenant,omitempty"`
	BranchID     *uint          `json:"branch_id"`
	Branch       *Branch        `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	Username     string         `gorm:"unique;not null" json:"username"`
	PasswordHash string         `gorm:"not null" json:"-"`
	RoleID       *uint          `json:"role_id"`
	Role         *Role          `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	IsOwner      bool           `gorm:"default:false" json:"is_owner"`
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type Role struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	TenantID    uint           `gorm:"not null;index" json:"tenant_id"`
	Name        string         `gorm:"not null" json:"name"`
	Permissions Permissions    `gorm:"type:json" json:"permissions"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// Permissions is a slice of strings that supports JSON serialization for GORM
type Permissions []string

func (p Permissions) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Permissions) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, p)
}

type EmployeeRepository interface {
	// Role CRUD
	CreateRole(role *Role) error
	GetRoleByID(id uint, tenantID uint) (*Role, error)
	ListRoles(tenantID uint) ([]Role, error)
	UpdateRole(role *Role) error
	DeleteRole(id uint, tenantID uint) error

	// User/Employee Management
	CreateUser(user *User) error
	ListUsers(tenantID uint) ([]User, error)
	GetUserByID(id uint, tenantID uint) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id uint, tenantID uint) error
}

type EmployeeUsecase interface {
	CreateRole(role *Role) error
	ListRoles(tenantID uint) ([]Role, error)
	UpdateRole(role *Role) error
	DeleteRole(id uint, tenantID uint) error

	CreateEmployee(username, password string, tenantID uint, branchID *uint, roleID uint) error
	ListEmployees(tenantID uint) ([]User, error)
	GetEmployeeByID(id uint, tenantID uint) (*User, error)
	UpdateEmployee(user *User) error
	ToggleActive(id uint, tenantID uint) error
	DeleteEmployee(id uint, tenantID uint) error
}
