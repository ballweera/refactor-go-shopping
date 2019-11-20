package user

// Service is service of user information
type Service struct {
}

// NewService instantiates UserService
func NewService() *Service {
	return &Service{}
}

// Register creates new user
func (s *Service) Register() User {
	userDAO := NewDAO()
	return userDAO.Register()
}
