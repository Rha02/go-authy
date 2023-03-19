package userService

// UserRepository is the interface for the user repository
type UserRepository interface {
	UpdateToken(userId string, newToken string) error
}
