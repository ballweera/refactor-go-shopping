package product

// Service is service wrapper of product information
type Service struct {
	productDAO *DAO
}

// NewService instantiates service of product module
func NewService() *Service {
	return &Service{
		productDAO: NewDAO(),
	}
}

// GetAllProducts return list of products
func (s *Service) GetAllProducts() []Product {
	return s.productDAO.GetAllProducts()
}
