package controller

import (
	"github.com/formulationapp/formulationapp/internal/service"
	"github.com/labstack/echo/v4"
)

type FormController interface {
	CreateForm(c echo.Context) error
	ListForms(c echo.Context) error
	GetForm(c echo.Context) error
	UpdateForm(c echo.Context) error
	DeleteForm(c echo.Context) error
}

type formController struct {
	formService service.FormService
}

func (f formController) CreateForm(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f formController) ListForms(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f formController) GetForm(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f formController) UpdateForm(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f formController) DeleteForm(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func newFormController(formService service.FormService) FormController {
	return &formController{
		formService: formService,
	}
}
