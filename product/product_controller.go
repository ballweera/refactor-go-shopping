package product

import (
	"encoding/json"
	"net/http"
)

// Handler is the controller of product information
type Handler struct {
	service *Service
}

// NewProductHandler returns handler of product module
func NewProductHandler() *Handler {
	return &Handler{service: NewProductService()}
}

// GetAllProducts returns all products information
func (ctrl *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := ctrl.service.GetAllProducts()
	json.NewEncoder(w).Encode(products)
}
