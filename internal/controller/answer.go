package controller

import "github.com/formulationapp/formulationapp/internal/service"

type AnswerController interface {
}

type answerController struct {
	answerService service.AnswerService
}

func newAnswerController(answerService service.AnswerService) AnswerController {
	return &answerController{
		answerService: answerService,
	}
}
