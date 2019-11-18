package product

// ProductService is service wrapper of product information
type ProductService struct {
	productDAO *ProductDAO
}

// NewProductService instantiates ProductService
func NewProductService() *ProductService {
	return &ProductService{
		productDAO: NewProductDAO(),
	}
}

// GetAllProducts return list of products
func (s *ProductService) GetAllProducts() []Product {
	return s.productDAO.GetAllProducts()
}
