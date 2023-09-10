package controller

import (
	"github.com/formulationapp/formulationapp/internal/service"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

type userController struct {
	userService service.UserService
}

func (u userController) Login(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (u userController) Register(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func newUserController(userService service.UserService) UserController {
	return &userController{userService}
}
