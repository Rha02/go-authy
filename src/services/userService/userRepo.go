package userService

// TODO: Write a user provider
type UserProvider struct{}

func NewUserProvider() UserRepository {
	return &UserProvider{}
}

// UpdateToken() updates the token of the user
func (u *UserProvider) UpdateToken(userId string, newToken string) error {
	return nil
}
