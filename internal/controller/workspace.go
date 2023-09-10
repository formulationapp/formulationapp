package controller

import (
	"github.com/formulationapp/formulationapp/internal/service"
	"github.com/labstack/echo/v4"
)

type WorkspaceController interface {
	ListWorkspaces(c echo.Context) error
}

type workspaceController struct {
	workspaceService service.WorkspaceService
}

func newWorkspaceController(workspaceService service.WorkspaceService) WorkspaceController {
	return &workspaceController{
		workspaceService: workspaceService,
	}
}

func (w workspaceController) ListWorkspaces(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
