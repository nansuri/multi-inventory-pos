package usecase

import (
	"github.com/username/multi-inventory-manager/internal/domain"
)

type branchUsecase struct {
	repo domain.BranchRepository
}

func NewBranchUsecase(repo domain.BranchRepository) domain.BranchUsecase {
	return &branchUsecase{repo: repo}
}

func (u *branchUsecase) CreateBranch(branch *domain.Branch) error {
	return u.repo.Create(branch)
}

func (u *branchUsecase) GetBranch(id uint, tenantID uint) (*domain.Branch, error) {
	return u.repo.GetByID(id, tenantID)
}

func (u *branchUsecase) ListBranches(tenantID uint) ([]domain.Branch, error) {
	return u.repo.List(tenantID)
}

func (u *branchUsecase) UpdateBranch(branch *domain.Branch) error {
	return u.repo.Update(branch)
}

func (u *branchUsecase) DeleteBranch(id uint, tenantID uint) error {
	return u.repo.Delete(id, tenantID)
}
