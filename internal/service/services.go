package service

import "github.com/formulationapp/formulationapp/internal/repository"

type Services interface {
	User() UserService
	Membership() MembershipService
	Workspace() WorkspaceService
	Form() FormService
	Answer() AnswerService
}

type services struct {
	userService       UserService
	membershipService MembershipService
	workspaceService  WorkspaceService
	formService       FormService
	answerService     AnswerService
}

func NewServices(repositories repository.Repositories) Services {
	userService := newUserService(repositories.User())
	membershipService := newMembershipService(repositories.Membership())
	workspaceService := newWorkspaceService(repositories.Workspace())
	formService := newFormService(repositories.Form())
	answerService := newAnswerService(repositories.Answer())

	return &services{
		userService:       userService,
		membershipService: membershipService,
		workspaceService:  workspaceService,
		formService:       formService,
		answerService:     answerService,
	}
}

func (s services) User() UserService {
	return s.userService
}

func (s services) Membership() MembershipService {
	return s.membershipService
}

func (s services) Workspace() WorkspaceService {
	return s.workspaceService
}

func (s services) Form() FormService {
	return s.formService
}

func (s services) Answer() AnswerService {
	return s.answerService
}