package service

import (
	"github.com/formulationapp/formulationapp/internal/model"
	"github.com/formulationapp/formulationapp/internal/repository"
)

type WorkspaceService interface {
	GetUserWorkspaces(userID uint) ([]model.Workspace, error)
	GetWorkspace(userID uint, workspaceID uint) (*model.Workspace, error)
}

type workspaceService struct {
	workspaceRepository  repository.WorkspaceRepository
	membershipRepository repository.MembershipRepository
}

func newWorkspaceService(workspaceRepository repository.WorkspaceRepository, membershipRepository repository.MembershipRepository) WorkspaceService {
	return &workspaceService{
		workspaceRepository:  workspaceRepository,
		membershipRepository: membershipRepository,
	}
}

func (w workspaceService) GetUserWorkspaces(userID uint) ([]model.Workspace, error) {
	memberships, err := w.membershipRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	var workspaces []model.Workspace
	for _, membership := range memberships {
		workspace, err := w.workspaceRepository.GetByID(membership.WorkspaceID)
		if err != nil {
			return nil, err
		}
		workspaces = append(workspaces, workspace)
	}
	return workspaces, nil
}

func (w workspaceService) GetWorkspace(userID uint, workspaceID uint) (*model.Workspace, error) {
	//TODO implement me
	panic("implement me")
}
