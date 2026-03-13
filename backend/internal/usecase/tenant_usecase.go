package usecase

import (
	"github.com/username/multi-inventory-manager/internal/domain"
)

type tenantUsecase struct {
	repo domain.TenantRepository
}

func NewTenantUsecase(repo domain.TenantRepository) domain.TenantUsecase {
	return &tenantUsecase{repo: repo}
}

func (u *tenantUsecase) GetTenant(id uint) (*domain.Tenant, error) {
	return u.repo.GetTenantByID(id)
}

func (u *tenantUsecase) UpdateTenant(id uint, name string, currency string, language string) error {
	tenant, err := u.repo.GetTenantByID(id)
	if err != nil {
		return err
	}

	tenant.Name = name
	tenant.Currency = currency
	tenant.Language = language

	return u.repo.UpdateTenant(tenant)
}
