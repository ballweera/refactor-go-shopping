package service

import "github.com/ballweera/refactor-go-shopping/model"

// UserService is service of user information
type UserService struct {
}

// NewUserService instantiates UserService
func NewUserService() *UserService {
	return &UserService{}
}

// Register creates new user
func (s *UserService) Register() model.User {
	return model.User{
		ID:    1,
		Name:  "Brown",
		Email: "brown@cafe.com",
	}
}
