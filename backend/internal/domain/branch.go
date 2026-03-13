package domain

import (
	"time"

	"gorm.io/gorm"
)

type Branch struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	TenantID  uint           `gorm:"not null;index" json:"tenant_id"`
	Tenant    Tenant         `gorm:"foreignKey:TenantID" json:"tenant,omitempty"`
	Name      string         `gorm:"not null" json:"name"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type BranchRepository interface {
	Create(branch *Branch) error
	GetByID(id uint, tenantID uint) (*Branch, error)
	List(tenantID uint) ([]Branch, error)
	Update(branch *Branch) error
	Delete(id uint, tenantID uint) error
}

type BranchUsecase interface {
	CreateBranch(branch *Branch) error
	GetBranch(id uint, tenantID uint) (*Branch, error)
	ListBranches(tenantID uint) ([]Branch, error)
	UpdateBranch(branch *Branch) error
	DeleteBranch(id uint, tenantID uint) error
}
