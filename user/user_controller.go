package user

import (
	"encoding/json"
	"net/http"
)

// Handler is controller of user
type Handler struct {
	service *Service
}

// NewUserController instantiate UserController
func NewUserController() *Handler {
	return &Handler{
		service: NewUserService(),
	}
}

// Register adds new user
func (ctrl *Handler) Register(w http.ResponseWriter, r *http.Request) {
	u := ctrl.service.Register()
	json.NewEncoder(w).Encode(u)
}
