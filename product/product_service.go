package product

// Service is service wrapper of product information
type Service struct {
	productDAO *DAO
}

// NewProductService instantiates service of product module
func NewProductService() *Service {
	return &Service{
		productDAO: NewProductDAO(),
	}
}

// GetAllProducts return list of products
func (s *Service) GetAllProducts() []Product {
	return s.productDAO.GetAllProducts()
}
