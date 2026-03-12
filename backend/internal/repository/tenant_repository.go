package repository

import (
	"github.com/username/multi-inventory-manager/internal/domain"
	"gorm.io/gorm"
)

type tenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) domain.TenantRepository {
	return &tenantRepository{db: db}
}

func (r *tenantRepository) GetTenantByID(id uint) (*domain.Tenant, error) {
	var tenant domain.Tenant
	err := r.db.First(&tenant, id).Error
	return &tenant, err
}

func (r *tenantRepository) UpdateTenant(tenant *domain.Tenant) error {
	return r.db.Save(tenant).Error
}
