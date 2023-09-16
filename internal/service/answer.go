package service

import (
	"errors"
	"github.com/formulationapp/formulationapp/internal/dto"
	"github.com/formulationapp/formulationapp/internal/model"
	"github.com/formulationapp/formulationapp/internal/repository"
)

type AnswerService interface {
	PutAnswer(submission, secret, field, value string) (model.Answer, error)
	GetAnswers(userID, workspaceID, formID uint) ([]map[string]string, error)
}

type answerService struct {
	workspaceService WorkspaceService
	formRepository   repository.FormRepository
	answerRepository repository.AnswerRepository
}

func newAnswerService(workspaceService WorkspaceService, formRepository repository.FormRepository, answerRepository repository.AnswerRepository) AnswerService {
	return &answerService{
		workspaceService: workspaceService,
		formRepository:   formRepository,
		answerRepository: answerRepository,
	}
}

func (a answerService) PutAnswer(submission, secret, field, value string) (model.Answer, error) {
	form, err := a.formRepository.GetBySecret(secret)
	if err != nil {
		return model.Answer{}, err
	}

	return a.answerRepository.Create(model.Answer{
		FormID:     form.ID,
		Submission: submission,
		Field:      field,
		Value:      value,
		Submitted:  false,
	})
}

func (a answerService) GetAnswers(userID, workspaceID, formID uint) ([]map[string]string, error) {
	var form model.Form
	access, err := a.workspaceService.UserHasAccessToWorkspace(userID, workspaceID)
	if err != nil {
		return nil, err
	}
	if !access {
		return nil, dto.AppError(errors.New("access denied"))
	}
	form, err = a.formRepository.GetByIDAndWorkspaceID(formID, workspaceID)
	if err != nil {
		return nil, err
	}

	answers, err := a.answerRepository.GetByFormID(form.ID)
	if err != nil {
		return nil, err
	}

	results := make([]map[string]string, 0)
	submissions := make(map[string][]model.Answer)
	for _, answer := range answers {
		if _, ok := submissions[answer.Submission]; !ok {
			submissions[answer.Submission] = make([]model.Answer, 0)
		}
		submissions[answer.Submission] = append(submissions[answer.Submission], answer)
	}

	for _, l := range submissions {
		m := make(map[string]string)
		for _, answer := range l {
			m[answer.Field] = answer.Value
		}
		results = append(results, m)
	}
	return results, nil
}
