package controller

import "github.com/formulationapp/formulationapp/internal/service"

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

func NewControllers(services service.Services) Controllers {
	userController := newUserController(services.User())
	workspaceController := newWorkspaceController(services.Workspace())
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
