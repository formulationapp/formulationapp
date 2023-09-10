package service

import (
	"github.com/formulationapp/formulationapp/internal/repository"
)

type AnswerService interface {
	PutAnswer(secret, field, value string) error
	GetAnswers(userID, formID string) ([]map[string]string, error)
}

type answerService struct {
	answerRepository repository.AnswerRepository
}

func newAnswerService(answerRepository repository.AnswerRepository) AnswerService {
	return &answerService{
		answerRepository: answerRepository,
	}
}

func (a answerService) PutAnswer(secret, field, value string) error {
	//TODO implement me
	panic("implement me")
}

func (a answerService) GetAnswers(userID, formID string) ([]map[string]string, error) {
	//TODO implement me
	panic("implement me")
}
