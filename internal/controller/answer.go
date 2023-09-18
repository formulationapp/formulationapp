package controller

import (
	"github.com/formulationapp/formulationapp/internal/dto"
	"github.com/formulationapp/formulationapp/internal/service"
	"github.com/formulationapp/formulationapp/internal/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AnswerController interface {
	PutAnswer(c echo.Context) error
	ListAnswers(c echo.Context) error
}

type answerController struct {
	userService   service.UserService
	answerService service.AnswerService
}

func (a answerController) PutAnswer(c echo.Context) error {
	var request dto.PutAnswerRequest
	err := c.Bind(request)
	if err != nil {
		return err
	}
	answer, err := a.answerService.PutAnswer(c.Param("submission"), c.Param("secret"), request.Field, request.Value)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, answer)
}

func (a answerController) ListAnswers(c echo.Context) error {
	user, err := a.userService.Authenticate(extractToken(c))
	if err != nil {
		return err
	}
	answers, err := a.answerService.GetAnswers(user.ID, util.ParseParamID(c.Param("workspaceID")), util.ParseParamID(c.Param("formID")))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, answers)
}

func newAnswerController(userService service.UserService, answerService service.AnswerService) AnswerController {
	return &answerController{
		userService:   userService,
		answerService: answerService,
	}
}
