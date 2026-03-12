package domain

import (
	"time"

	"gorm.io/gorm"
)

type Tenant struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Currency  string         `gorm:"default:USD" json:"currency"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type TenantRepository interface {
	GetTenantByID(id uint) (*Tenant, error)
	UpdateTenant(tenant *Tenant) error
}

type TenantUsecase interface {
	GetTenant(id uint) (*Tenant, error)
	UpdateTenant(id uint, name string, currency string) error
}
