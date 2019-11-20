package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ballweera/refactor-go-shopping/internal/user"
)

// UserHandler is controller of user
type UserHandler struct {
	service *user.Service
}

// NewUserHandler instantiate handler of user module
func NewUserHandler() *UserHandler {
	return &UserHandler{
		service: user.NewUserService(),
	}
}

// Register adds new user
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	u := h.service.Register()
	json.NewEncoder(w).Encode(u)
}
