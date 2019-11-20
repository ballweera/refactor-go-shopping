package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ballweera/refactor-go-shopping/internal/product"
)

// ProductHandler is the handler of product information
type ProductHandler struct {
	service *product.Service
}

// NewProductHandler returns handler of product module
func NewProductHandler() *ProductHandler {
	return &ProductHandler{service: product.NewProductService()}
}

// GetAllProducts returns all products information
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := h.service.GetAllProducts()
	json.NewEncoder(w).Encode(products)
}
