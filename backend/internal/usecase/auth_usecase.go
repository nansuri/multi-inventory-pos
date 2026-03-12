package usecase

import (
	"errors"
	"log"

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
		IsOwner:      true,
	}

	if err := u.repo.CreateUser(user); err != nil {
		return "", err
	}

	// 4. Seed Default Roles
	defaultRoles := []domain.Role{
		{
			TenantID:    tenant.ID,
			Name:        "Manager",
			Permissions: []string{"dashboard", "inventory", "recipes", "production", "production_log", "pos_order", "pos_payment", "order_history", "employees", "settings"},
		},
		{
			TenantID:    tenant.ID,
			Name:        "Chef",
			Permissions: []string{"dashboard", "inventory", "recipes", "production", "production_log"},
		},
		{
			TenantID:    tenant.ID,
			Name:        "Waiter",
			Permissions: []string{"dashboard", "pos_order", "pos_payment"},
		},
	}

	for _, role := range defaultRoles {
		err := u.repo.CreateRole(&role)
		if err != nil {
			log.Printf("Error seeding role %s for tenant %d: %v", role.Name, tenant.ID, err)
		} else {
			log.Printf("Successfully seeded role %s for tenant %d", role.Name, tenant.ID)
		}
	}

	// 5. Generate Token
	return auth.GenerateToken(user.ID, user.TenantID, "owner", u.jwtSecret)
}

func (u *authUsecase) Login(username, password string) (string, error) {
	user, err := u.repo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !auth.CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New("invalid credentials")
	}

	roleName := "employee"
	if user.IsOwner {
		roleName = "owner"
	} else if user.Role != nil {
		roleName = user.Role.Name
	}

	return auth.GenerateToken(user.ID, user.TenantID, roleName, u.jwtSecret)
}

func (u *authUsecase) GetMe(userID uint, tenantID uint) (*domain.User, error) {
	return u.repo.GetUserByID(userID, tenantID)
}
