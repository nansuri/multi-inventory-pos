package usecase

import (
	"errors"

	"github.com/username/multi-inventory-manager/internal/domain"
	"github.com/username/multi-inventory-manager/pkg/auth"
)

type authUsecase struct {
	repo      domain.AuthRepository
	jwtSecret string
}

func NewAuthUsecase(repo domain.AuthRepository, jwtSecret string) domain.AuthUsecase {
	return &authUsecase{repo: repo, jwtSecret: jwtSecret}
}

func (u *authUsecase) Register(username, password, tenantName string) (string, error) {
	// 1. Create Tenant
	tenant := &domain.Tenant{Name: tenantName}
	if err := u.repo.CreateTenant(tenant); err != nil {
		return "", err
	}

	// 2. Hash Password
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return "", err
	}

	// 3. Create User (Owner)
	user := &domain.User{
		Username:     username,
		PasswordHash: hashedPassword,
		TenantID:     tenant.ID,
		Role:         domain.RoleOwner,
	}

	if err := u.repo.CreateUser(user); err != nil {
		return "", err
	}

	// 4. Generate Token
	return auth.GenerateToken(user.ID, user.TenantID, string(user.Role), u.jwtSecret)
}

func (u *authUsecase) Login(username, password string) (string, error) {
	user, err := u.repo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !auth.CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New("invalid credentials")
	}

	return auth.GenerateToken(user.ID, user.TenantID, string(user.Role), u.jwtSecret)
}
