package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ballweera/refactor-go-shopping/model"
)

// GetAllProducts returns all products information
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := []model.Product{
		model.Product{SKU: "12345", Name: "iPhone", Price: 1000.00, Unit: 100},
	}
	json.NewEncoder(w).Encode(products)
}
