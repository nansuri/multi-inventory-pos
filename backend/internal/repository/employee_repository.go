package repository

import (
	"github.com/username/multi-inventory-manager/internal/domain"
	"gorm.io/gorm"
)

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) domain.EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) CreateRole(role *domain.Role) error {
	return r.db.Create(role).Error
}

func (r *employeeRepository) GetRoleByID(id uint, tenantID uint) (*domain.Role, error) {
	var role domain.Role
	err := r.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&role).Error
	return &role, err
}

func (r *employeeRepository) ListRoles(tenantID uint) ([]domain.Role, error) {
	var roles []domain.Role
	err := r.db.Where("tenant_id = ?", tenantID).Find(&roles).Error
	return roles, err
}

func (r *employeeRepository) UpdateRole(role *domain.Role) error {
	return r.db.Save(role).Error
}

func (r *employeeRepository) DeleteRole(id uint, tenantID uint) error {
	return r.db.Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&domain.Role{}).Error
}

func (r *employeeRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *employeeRepository) ListUsers(tenantID uint) ([]domain.User, error) {
	var users []domain.User
	err := r.db.Preload("Role").Where("tenant_id = ? AND is_owner = false", tenantID).Find(&users).Error
	return users, err
}

func (r *employeeRepository) GetUserByID(id uint, tenantID uint) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("Role").Where("id = ? AND tenant_id = ?", id, tenantID).First(&user).Error
	return &user, err
}

func (r *employeeRepository) UpdateUser(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *employeeRepository) DeleteUser(id uint, tenantID uint) error {
	return r.db.Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&domain.User{}).Error
}
