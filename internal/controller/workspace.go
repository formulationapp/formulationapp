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
	userService      service.UserService
}

func newWorkspaceController(workspaceService service.WorkspaceService, userService service.UserService) WorkspaceController {
	return &workspaceController{
		workspaceService: workspaceService,
		userService:      userService,
	}
}

func (w workspaceController) ListWorkspaces(c echo.Context) error {
	user, err := w.userService.Authenticate(extractToken(c))
	if err != nil {
		return err
	}
	workspaces, err := w.workspaceService.GetUserWorkspaces(user.ID)
	if err != nil {
		return err
	}
	return c.JSON(200, workspaces)
}
