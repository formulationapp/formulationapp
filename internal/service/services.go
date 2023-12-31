package service

import (
	"github.com/formulationapp/formulationapp/internal/client"
	"github.com/formulationapp/formulationapp/internal/dto"
	"github.com/formulationapp/formulationapp/internal/repository"
)

type Services interface {
	User() UserService
	Membership() MembershipService
	Workspace() WorkspaceService
	Form() FormService
	Answer() AnswerService
	AI() AIService
}

type services struct {
	userService       UserService
	membershipService MembershipService
	workspaceService  WorkspaceService
	formService       FormService
	answerService     AnswerService
	aiService         AIService
}

func NewServices(repositories repository.Repositories, clients client.Clients, config dto.Config) Services {
	userService := newUserService(repositories.User(), repositories.Membership(), repositories.Workspace(), config)
	membershipService := newMembershipService(repositories.Membership())
	workspaceService := newWorkspaceService(repositories.Workspace(), repositories.Membership())
	formService := newFormService(repositories.Form(), workspaceService)
	answerService := newAnswerService(workspaceService, repositories.Form(), repositories.Answer())
	aiService := newAIService(clients.AI())

	return &services{
		userService:       userService,
		membershipService: membershipService,
		workspaceService:  workspaceService,
		formService:       formService,
		answerService:     answerService,
		aiService:         aiService,
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

func (s services) AI() AIService {
	return s.aiService
}
