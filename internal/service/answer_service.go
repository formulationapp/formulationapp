package service

import "github.com/formulationapp/formulationapp/internal/repository"

type AnswerService interface {
}

type answerService struct {
	answerRepository repository.AnswerRepository
}

func newAnswerService(answerRepository repository.AnswerRepository) AnswerService {
	return &answerService{
		answerRepository: answerRepository,
	}
}
