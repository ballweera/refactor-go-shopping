package user

// Service is service of user information
type Service struct {
}

// NewUserService instantiates UserService
func NewUserService() *Service {
	return &Service{}
}

// Register creates new user
func (s *Service) Register() User {
	userDAO := NewUserDAO()
	return userDAO.Register()
}
