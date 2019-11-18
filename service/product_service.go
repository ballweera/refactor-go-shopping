package service

import "github.com/ballweera/refactor-go-shopping/model"

// ProductService is service wrapper of product information
type ProductService struct{}

// GetAllProducts return list of products
func (s *ProductService) GetAllProducts() []model.Product {
	products := []model.Product{
		{SKU: "12345", Name: "iPhone", Price: 100.00, Unit: 100},
		{SKU: "22345", Name: "Pixel 4", Price: 70.00, Unit: 200},
	}
	return products
}
