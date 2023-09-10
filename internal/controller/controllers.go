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
}

type controllers struct {
	userController       UserController
	workspaceController  WorkspaceController
	membershipController MembershipController
	formController       FormController
	answerController     AnswerController
}

func NewControllers(services service.Services, e *echo.Echo) Controllers {
	userController := newUserController(services.User())
	workspaceController := newWorkspaceController(services.Workspace())
	membershipController := newMembershipController(services.Membership())
	formController := newFormController(services.Form())
	answerController := newAnswerController(services.Answer())

	e.POST("/auth/login", userController.Login)
	e.POST("/auth/register", userController.Register)

	e.GET("/workspaces", workspaceController.ListWorkspaces)

	e.GET("/workspaces/:workspaceID/forms", formController.ListForms)
	e.GET("/workspaces/:workspaceID/forms/:formID", formController.GetForm)
	e.POST("/workspaces/:workspaceID/forms", formController.CreateForm)
	e.PUT("/workspaces/:workspaceID/forms", formController.UpdateForm)
	e.DELETE("/workspaces/:workspaceID/forms", formController.DeleteForm)

	e.GET("/workspaces/:workspaceID/forms/:formID/answers", answerController.ListAnswers)

	e.PUT("/forms/:secret/answers", answerController.PutAnswer)

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
