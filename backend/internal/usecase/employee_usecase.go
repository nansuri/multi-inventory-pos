package usecase

import (
	"github.com/username/multi-inventory-manager/internal/domain"
	"github.com/username/multi-inventory-manager/pkg/auth"
)

type employeeUsecase struct {
	repo domain.EmployeeRepository
}

func NewEmployeeUsecase(repo domain.EmployeeRepository) domain.EmployeeUsecase {
	return &employeeUsecase{repo: repo}
}

func (u *employeeUsecase) CreateRole(role *domain.Role) error {
	return u.repo.CreateRole(role)
}

func (u *employeeUsecase) ListRoles(tenantID uint) ([]domain.Role, error) {
	return u.repo.ListRoles(tenantID)
}

func (u *employeeUsecase) UpdateRole(role *domain.Role) error {
	return u.repo.UpdateRole(role)
}

func (u *employeeUsecase) DeleteRole(id uint, tenantID uint) error {
	return u.repo.DeleteRole(id, tenantID)
}

func (u *employeeUsecase) CreateEmployee(username, password string, tenantID uint, branchID *uint, roleID uint) error {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Username:     username,
		PasswordHash: hashedPassword,
		TenantID:     tenantID,
		BranchID:     branchID,
		RoleID:       &roleID,
		IsOwner:      false,
	}

	return u.repo.CreateUser(user)
}

func (u *employeeUsecase) ListEmployees(tenantID uint) ([]domain.User, error) {
	return u.repo.ListUsers(tenantID)
}

func (u *employeeUsecase) GetEmployeeByID(id uint, tenantID uint) (*domain.User, error) {
	return u.repo.GetUserByID(id, tenantID)
}

func (u *employeeUsecase) UpdateEmployee(user *domain.User) error {
	return u.repo.UpdateUser(user)
}

func (u *employeeUsecase) ToggleActive(id uint, tenantID uint) error {
	user, err := u.repo.GetUserByID(id, tenantID)
	if err != nil {
		return err
	}

	user.IsActive = !user.IsActive
	return u.repo.UpdateUser(user)
}

func (u *employeeUsecase) DeleteEmployee(id uint, tenantID uint) error {
	return u.repo.DeleteUser(id, tenantID)
}
