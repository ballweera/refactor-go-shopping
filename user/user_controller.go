package user

import (
	"encoding/json"
	"net/http"
)

// Controller is controller of user
type Controller struct {
	service *Service
}

// NewUserController instantiate UserController
func NewUserController() *Controller {
	return &Controller{
		service: NewUserService(),
	}
}

// Register adds new user
func (ctrl *Controller) Register(w http.ResponseWriter, r *http.Request) {
	u := ctrl.service.Register()
	json.NewEncoder(w).Encode(u)
}
