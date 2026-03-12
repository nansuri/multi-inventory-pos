package repository

import (
	"github.com/username/multi-inventory-manager/internal/domain"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateTenant(tenant *domain.Tenant) error {
	return r.db.Create(tenant).Error
}

func (r *authRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) CreateRole(role *domain.Role) error {
	return r.db.Create(role).Error
}

func (r *authRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("Role").Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *authRepository) GetUserByID(id uint, tenantID uint) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("Role").Where("id = ? AND tenant_id = ?", id, tenantID).First(&user).Error
	return &user, err
}
