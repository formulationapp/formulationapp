package service

import "github.com/formulationapp/formulationapp/internal/repository"

type FormService interface {
}

type formService struct {
	formRepository repository.FormRepository
}

func newFormService(formRepository repository.FormRepository) FormService {
	return &formService{
		formRepository: formRepository,
	}
}
