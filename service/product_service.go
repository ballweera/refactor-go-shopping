package service

import (
	"github.com/ballweera/refactor-go-shopping/dao"
	"github.com/ballweera/refactor-go-shopping/model"
)

// ProductService is service wrapper of product information
type ProductService struct {
	productDAO *dao.ProductDAO
}

// NewProductService instantiates ProductService
func NewProductService() *ProductService {
	return &ProductService{
		productDAO: dao.NewProductDAO(),
	}
}

// GetAllProducts return list of products
func (s *ProductService) GetAllProducts() []model.Product {
	return s.productDAO.GetAllProducts()
}
