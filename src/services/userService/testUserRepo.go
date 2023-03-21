package userService

// TODO: Write mock user provider meant for testing
type TestUserProvider struct{}

func NewTestUserProvider() UserRepository {
	return &TestUserProvider{}
}

// UpdateToken() updates the token of the user
func (u *TestUserProvider) UpdateToken(userId string, newToken string) error {
	return nil
}
