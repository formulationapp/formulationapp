package repository

import "gorm.io/gorm"

type Repositories interface {
	User() UserRepository
	Membership() MembershipRepository
	Workspace() WorkspaceRepository
	Form() FormRepository
	Answer() AnswerRepository
}

type repositories struct {
	userRepository       UserRepository
	membershipRepository MembershipRepository
	workspaceRepository  WorkspaceRepository
	formRepository       FormRepository
	answerRepository     AnswerRepository
}

func NewRepositories(db *gorm.DB) Repositories {
	userRepository := newUserRepository(db)
	membershipRepository := newMembershipRepository(db)
	workspaceRepository := newWorkspaceRepository(db)
	formRepository := newFormRepository(db)
	answerRepository := newAnswerRepository(db)
	return &repositories{
		userRepository:       userRepository,
		membershipRepository: membershipRepository,
		workspaceRepository:  workspaceRepository,
		formRepository:       formRepository,
		answerRepository:     answerRepository,
	}
}

func (r repositories) User() UserRepository {
	return r.userRepository
}

func (r repositories) Membership() MembershipRepository {
	return r.membershipRepository
}

func (r repositories) Workspace() WorkspaceRepository {
	return r.workspaceRepository
}

func (r repositories) Form() FormRepository {
	return r.formRepository
}

func (r repositories) Answer() AnswerRepository {
	return r.answerRepository
}
