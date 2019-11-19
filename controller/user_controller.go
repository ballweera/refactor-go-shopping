package controller

import "net/http"

import "encoding/json"

// Register adds new user
func Register(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"name": "ballweera"})
}
