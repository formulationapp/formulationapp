package controller

import (
	"context"
	"encoding/json"
	"github.com/formulationapp/formulationapp/internal/dto"
	"github.com/formulationapp/formulationapp/internal/model"
	"github.com/formulationapp/formulationapp/internal/service"
	"github.com/formulationapp/formulationapp/internal/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

type FormController interface {
	CreateForm(c echo.Context) error
	ListForms(c echo.Context) error
	GetForm(c echo.Context) error
	UpdateForm(c echo.Context) error
	DeleteForm(c echo.Context) error
	GenerateForm(c echo.Context) error

	ViewForm(c echo.Context) error
}

type formController struct {
	userService service.UserService
	formService service.FormService
	aiService   service.AIService
}

func (f formController) CreateForm(c echo.Context) error {
	userID, err := f.userService.DecodeToken(util.GetTokenFromContext(c))
	if err != nil {
		return err
	}

	var payload dto.CreateFormRequest
	err = c.Bind(&payload)
	if err != nil {
		return err
	}

	form, err := f.formService.CreateForm(userID, util.ParseParamID(c.Param("workspaceID")), payload)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, form)
}

func (f formController) ListForms(c echo.Context) error {
	userID, err := f.userService.DecodeToken(util.GetTokenFromContext(c))
	if err != nil {
		return err
	}

	forms, err := f.formService.GetForms(userID, util.ParseParamID(c.Param("workspaceID")))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, forms)
}

func (f formController) GetForm(c echo.Context) error {
	userID, err := f.userService.DecodeToken(util.GetTokenFromContext(c))
	if err != nil {
		return err
	}

	form, err := f.formService.GetForm(userID, util.ParseParamID(c.Param("workspaceID")), util.ParseParamID(c.Param("formID")))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, form)
}

func (f formController) UpdateForm(c echo.Context) error {
	userID, err := f.userService.DecodeToken(util.GetTokenFromContext(c))
	if err != nil {
		return err
	}

	var payload dto.UpdateFormRequest
	err = c.Bind(&payload)
	if err != nil {
		return err
	}

	form, err := f.formService.UpdateForm(userID, util.ParseParamID(c.Param("workspaceID")), util.ParseParamID(c.Param("formID")), payload)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, form)
}

func (f formController) DeleteForm(c echo.Context) error {
	userID, err := f.userService.DecodeToken(util.GetTokenFromContext(c))
	if err != nil {
		return err
	}

	err = f.formService.DeleteForm(userID, util.ParseParamID(c.Param("workspaceID")), util.ParseParamID(c.Param("formID")))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, model.Form{})
}

func (f formController) ViewForm(c echo.Context) error {
	form, err := f.formService.ViewForm(c.Param("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, form)
}

func (f formController) GenerateForm(c echo.Context) error {
	var payload dto.GenerateFormRequest
	err := c.Bind(&payload)
	if err != nil {
		return err
	}
	form, err := f.aiService.GenerateForm(context.Background(), payload.Prompt)
	if err != nil {
		return err
	}
	var data interface{}
	err = json.Unmarshal([]byte(form), &data)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}

func newFormController(userService service.UserService, formService service.FormService, aiService service.AIService) FormController {
	return &formController{
		userService: userService,
		formService: formService,
		aiService:   aiService,
	}
}
