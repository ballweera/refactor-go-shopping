package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ballweera/refactor-go-shopping/internal/user"
)

// Handler is controller of user
type Handler struct {
	service *user.Service
}

// NewUserHandler instantiate handler of user module
func NewUserHandler() *Handler {
	return &Handler{
		service: user.NewUserService(),
	}
}

// Register adds new user
func (ctrl *Handler) Register(w http.ResponseWriter, r *http.Request) {
	u := ctrl.service.Register()
	json.NewEncoder(w).Encode(u)
}
