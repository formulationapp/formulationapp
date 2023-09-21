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
	formController := newFormController(services.User(), services.Form())
	answerController := newAnswerController(services.User(), services.Answer())
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
	e.POST("/api/auth/login", c.userController.Login)
	e.POST("/api/auth/register", c.userController.Register)

	e.GET("/api/workspaces", c.workspaceController.ListWorkspaces)

	e.GET("/api/workspaces/:workspaceID/forms", c.formController.ListForms)
	e.GET("/api/workspaces/:workspaceID/forms/:formID", c.formController.GetForm)
	e.POST("/api/workspaces/:workspaceID/forms", c.formController.CreateForm)
	e.PUT("/api/workspaces/:workspaceID/forms/:formID", c.formController.UpdateForm)
	e.DELETE("/api/workspaces/:workspaceID/forms/:formID", c.formController.DeleteForm)

	e.GET("/api/workspaces/:workspaceID/forms/:formID/answers", c.answerController.ListAnswers)

	e.GET("/api/forms/:secret", c.formController.ViewForm)
	e.PUT("/api/forms/:secret/answers/:submission", c.answerController.PutAnswer)
}
