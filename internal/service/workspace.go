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
	workspaceRepository repository.WorkspaceRepository
}

func newWorkspaceService(workspaceRepository repository.WorkspaceRepository) WorkspaceService {
	return &workspaceService{
		workspaceRepository: workspaceRepository,
	}
}

func (w workspaceService) GetUserWorkspaces(userID uint) ([]model.Workspace, error) {
	//TODO implement me
	panic("implement me")
}

func (w workspaceService) GetWorkspace(userID uint, workspaceID uint) (*model.Workspace, error) {
	//TODO implement me
	panic("implement me")
}
