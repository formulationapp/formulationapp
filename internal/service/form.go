package service

import (
	"github.com/formulationapp/formulationapp/internal/dto"
	"github.com/formulationapp/formulationapp/internal/model"
	"github.com/formulationapp/formulationapp/internal/repository"
)

type FormService interface {
	CreateForm(userID, workspaceID uint, form *dto.CreateFormRequest) (model.Form, error)
	GetForms(userID, workspaceID uint) ([]model.Form, error)
	GetForm(userID, formID uint) (model.Form, error)
	UpdateForm(userID, formID uint, form *dto.UpdateFormRequest) (model.Form, error)
	DeleteForm(userID, formID uint) error
}

type formService struct {
	formRepository repository.FormRepository
}

func newFormService(formRepository repository.FormRepository) FormService {
	return &formService{
		formRepository: formRepository,
	}
}

func (f formService) CreateForm(userID, workspaceID uint, form *dto.CreateFormRequest) (model.Form, error) {
	//TODO implement me
	panic("implement me")
}

func (f formService) GetForms(userID, workspaceID uint) ([]model.Form, error) {
	//TODO implement me
	panic("implement me")
}

func (f formService) GetForm(userID, formID uint) (model.Form, error) {
	//TODO implement me
	panic("implement me")
}

func (f formService) UpdateForm(userID, formID uint, form *dto.UpdateFormRequest) (model.Form, error) {
	//TODO implement me
	panic("implement me")
}

func (f formService) DeleteForm(userID, formID uint) error {
	//TODO implement me
	panic("implement me")
}
