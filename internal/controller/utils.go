package controller

import (
	"github.com/labstack/echo/v4"
	"strings"
)

func extractToken(c echo.Context) string {
	return strings.ReplaceAll(c.Request().Header.Get("Authorization"), "Bearer ", "")
}
