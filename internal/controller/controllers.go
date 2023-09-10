package controller

import (
	"github.com/formulationapp/formulationapp/internal/service"
	"github.com/labstack/echo/v4"
)

type Controllers interface {
	User() UserController
	Workspace() WorkspaceController
	Membership() MembershipController
	Form() FormController
	Answer() AnswerController

	Route(e *echo.Echo)
}

type controllers struct {
	userController       UserController
	workspaceController  WorkspaceController
	membershipController MembershipController
	formController       FormController
	answerController     AnswerController
}

func NewControllers(services service.Services) Controllers {
	userController := newUserController(services.User())
	workspaceController := newWorkspaceController(services.Workspace(), services.User())
	membershipController := newMembershipController(services.Membership())
	formController := newFormController(services.Form())
	answerController := newAnswerController(services.Answer())
	return &controllers{
		userController:       userController,
		workspaceController:  workspaceController,
		membershipController: membershipController,
		formController:       formController,
		answerController:     answerController,
	}
}

func (c controllers) User() UserController {
	return c.userController
}

func (c controllers) Workspace() WorkspaceController {
	return c.workspaceController
}

func (c controllers) Membership() MembershipController {
	return c.membershipController
}

func (c controllers) Form() FormController {
	return c.formController
}

func (c controllers) Answer() AnswerController {
	return c.answerController
}

func (c controllers) Route(e *echo.Echo) {
	e.POST("/auth/login", c.userController.Login)
	e.POST("/auth/register", c.userController.Register)

	e.GET("/workspaces", c.workspaceController.ListWorkspaces)

	e.GET("/workspaces/:workspaceID/forms", c.formController.ListForms)
	e.GET("/workspaces/:workspaceID/forms/:formID", c.formController.GetForm)
	e.POST("/workspaces/:workspaceID/forms", c.formController.CreateForm)
	e.PUT("/workspaces/:workspaceID/forms", c.formController.UpdateForm)
	e.DELETE("/workspaces/:workspaceID/forms", c.formController.DeleteForm)

	e.GET("/workspaces/:workspaceID/forms/:formID/answers", c.answerController.ListAnswers)

	e.PUT("/forms/:secret/answers", c.answerController.PutAnswer)
}
