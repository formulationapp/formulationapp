package service

import (
	"errors"
	"github.com/formulationapp/formulationapp/internal/dto"
	"github.com/formulationapp/formulationapp/internal/model"
	"github.com/formulationapp/formulationapp/internal/repository"
	"github.com/formulationapp/formulationapp/internal/util"
)

type FormService interface {
	CreateForm(userID, workspaceID uint, form dto.CreateFormRequest) (model.Form, error)
	GetForms(userID, workspaceID uint) ([]model.Form, error)
	GetForm(userID, workspaceID, formID uint) (model.Form, error)
	UpdateForm(userID, workspaceID, formID uint, form dto.UpdateFormRequest) (model.Form, error)
	DeleteForm(userID, workspaceID, formID uint) error
}

type formService struct {
	formRepository   repository.FormRepository
	workspaceService WorkspaceService
}

func newFormService(formRepository repository.FormRepository, workspaceService WorkspaceService) FormService {
	return &formService{
		formRepository:   formRepository,
		workspaceService: workspaceService,
	}
}

func (f formService) CreateForm(userID, workspaceID uint, form dto.CreateFormRequest) (model.Form, error) {
	access, err := f.workspaceService.UserHasAccessToWorkspace(userID, workspaceID)
	if err != nil {
		return model.Form{}, err
	}
	if !access {
		return model.Form{}, dto.AppError(errors.New("access denied"))
	}

	return f.formRepository.Create(model.Form{
		Name:        form.Name,
		Definition:  form.Definition,
		WorkspaceID: workspaceID,
		Secret:      util.RandStringRunes(15),
	})
}

func (f formService) GetForms(userID, workspaceID uint) ([]model.Form, error) {
	access, err := f.workspaceService.UserHasAccessToWorkspace(userID, workspaceID)
	if err != nil {
		return nil, err
	}
	if !access {
		return nil, dto.AppError(errors.New("access denied"))
	}

	return f.formRepository.GetByWorkspaceID(workspaceID)
}

func (f formService) GetForm(userID, workspaceID, formID uint) (model.Form, error) {
	var form model.Form
	access, err := f.workspaceService.UserHasAccessToWorkspace(userID, workspaceID)
	if err != nil {
		return form, err
	}
	if !access {
		return form, dto.AppError(errors.New("access denied"))
	}

	return f.formRepository.GetByIDAndWorkspaceID(formID, workspaceID)
}

func (f formService) UpdateForm(userID, workspaceID, formID uint, payload dto.UpdateFormRequest) (model.Form, error) {
	var form model.Form
	access, err := f.workspaceService.UserHasAccessToWorkspace(userID, workspaceID)
	if err != nil {
		return form, err
	}
	if !access {
		return form, dto.AppError(errors.New("access denied"))
	}

	form, err = f.formRepository.GetByIDAndWorkspaceID(formID, workspaceID)
	if err != nil {
		return form, nil
	}

	form.Name = payload.Name
	form.Definition = payload.Definition

	return f.formRepository.Update(form)
}

func (f formService) DeleteForm(userID, workspaceID, formID uint) error {
	var form model.Form
	access, err := f.workspaceService.UserHasAccessToWorkspace(userID, workspaceID)
	if err != nil {
		return err
	}
	if !access {
		return dto.AppError(errors.New("access denied"))
	}

	form, err = f.formRepository.GetByIDAndWorkspaceID(formID, workspaceID)
	if err != nil {
		return nil
	}

	return f.formRepository.Delete(form)
}
