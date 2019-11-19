package dao

import "github.com/ballweera/refactor-go-shopping/model"

// UserDAO is DAO of User information
type UserDAO struct{}

// NewUserDAO instantiates UserDAO
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// Register creates new user
func (dao UserDAO) Register() model.User {
	return model.User{
		ID:    1,
		Name:  "Brown",
		Email: "brown@cafe.com",
	}
}
