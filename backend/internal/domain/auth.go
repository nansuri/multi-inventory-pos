package domain

type AuthRepository interface {
	CreateTenant(tenant *Tenant) error
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
}

type AuthUsecase interface {
	Register(username, password, tenantName string) (string, error)
	Login(username, password string) (string, error)
}
