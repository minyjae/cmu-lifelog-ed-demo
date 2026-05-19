package services

type RequireRoleService interface {
	// @Read
	GetRoleByEmail(email string) (string, error)
}
