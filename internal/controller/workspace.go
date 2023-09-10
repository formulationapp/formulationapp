package controller

import "github.com/formulationapp/formulationapp/internal/service"

type WorkspaceController interface {
}

type workspaceController struct {
	workspaceService service.WorkspaceService
}

func newWorkspaceController(workspaceService service.WorkspaceService) WorkspaceController {
	return &workspaceController{
		workspaceService: workspaceService,
	}
}
