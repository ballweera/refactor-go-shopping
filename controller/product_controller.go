package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ballweera/refactor-go-shopping/service"
)

// ProductController is the controller of product information
type ProductController struct {
	service *service.ProductService
}

// NewProductController returns instance of ProductController
func NewProductController() *ProductController {
	return &ProductController{service: service.NewProductService()}
}

// GetAllProducts returns all products information
func (ctrl *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := ctrl.service.GetAllProducts()
	json.NewEncoder(w).Encode(products)
}
