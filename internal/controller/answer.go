package controller

import (
	"github.com/formulationapp/formulationapp/internal/service"
	"github.com/labstack/echo/v4"
)

type AnswerController interface {
	PutAnswer(c echo.Context) error
	ListAnswers(c echo.Context) error
}

type answerController struct {
	answerService service.AnswerService
}

func (a answerController) PutAnswer(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a answerController) ListAnswers(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func newAnswerController(answerService service.AnswerService) AnswerController {
	return &answerController{
		answerService: answerService,
	}
}
