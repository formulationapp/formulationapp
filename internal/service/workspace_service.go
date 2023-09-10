package service

import "github.com/formulationapp/formulationapp/internal/repository"

type WorkspaceService interface {
}

type workspaceService struct {
	workspaceRepository repository.WorkspaceRepository
}

func newWorkspaceService(workspaceRepository repository.WorkspaceRepository) WorkspaceService {
	return &workspaceService{
		workspaceRepository: workspaceRepository,
	}
}
