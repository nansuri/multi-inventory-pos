package repository

import (
	"github.com/username/multi-inventory-manager/internal/domain"
	"gorm.io/gorm"
)

type branchRepository struct {
	db *gorm.DB
}

func NewBranchRepository(db *gorm.DB) domain.BranchRepository {
	return &branchRepository{db: db}
}

func (r *branchRepository) Create(branch *domain.Branch) error {
	return r.db.Create(branch).Error
}

func (r *branchRepository) GetByID(id uint, tenantID uint) (*domain.Branch, error) {
	var branch domain.Branch
	err := r.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&branch).Error
	return &branch, err
}

func (r *branchRepository) List(tenantID uint) ([]domain.Branch, error) {
	var branches []domain.Branch
	err := r.db.Where("tenant_id = ?", tenantID).Find(&branches).Error
	return branches, err
}

func (r *branchRepository) Update(branch *domain.Branch) error {
	return r.db.Save(branch).Error
}

func (r *branchRepository) Delete(id uint, tenantID uint) error {
	return r.db.Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&domain.Branch{}).Error
}
