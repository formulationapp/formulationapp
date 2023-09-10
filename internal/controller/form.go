package controller

import "github.com/formulationapp/formulationapp/internal/service"

type FormController interface {
}

type formController struct {
	formService service.FormService
}

func newFormController(formService service.FormService) FormController {
	return &formController{
		formService: formService,
	}
}
