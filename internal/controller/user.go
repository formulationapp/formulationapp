package controller

import "github.com/formulationapp/formulationapp/internal/service"

type UserController interface {
}

type userController struct {
	userService service.UserService
}

func newUserController(userService service.UserService) UserController {
	return &userController{userService}
}
