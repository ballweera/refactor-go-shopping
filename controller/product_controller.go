package controller

import "net/http"

import "encoding/json"

// GetAllProducts returns all products information
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := map[string]string{
		"status": "ok",
	}

	json.NewEncoder(w).Encode(products)
}
