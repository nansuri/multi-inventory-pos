package domain

type AuthRepository interface {
	CreateTenant(tenant *Tenant) error
	CreateUser(user *User) error
	CreateRole(role *Role) error
	GetUserByUsername(username string) (*User, error)
	GetUserByID(id uint, tenantID uint) (*User, error)
}

type AuthUsecase interface {
	Register(username, password, tenantName string) (string, error)
	Login(username, password string) (string, error)
	GetMe(userID uint, tenantID uint) (*User, error)
}
