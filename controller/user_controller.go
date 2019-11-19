package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ballweera/refactor-go-shopping/service"
)

// Register adds new user
func Register(w http.ResponseWriter, r *http.Request) {
	userService := service.NewUserService()
	u := userService.Register()
	json.NewEncoder(w).Encode(u)
}
