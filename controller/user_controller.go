package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ballweera/refactor-go-shopping/service"
)

// UserController is controller of user
type UserController struct {
	service *service.UserService
}

// NewUserController instantiate UserController
func NewUserController() *UserController {
	return &UserController{
		service: service.NewUserService(),
	}
}

// Register adds new user
func (ctrl *UserController) Register(w http.ResponseWriter, r *http.Request) {
	u := ctrl.service.Register()
	json.NewEncoder(w).Encode(u)
}
