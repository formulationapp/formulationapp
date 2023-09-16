package util

import (
	"github.com/labstack/echo/v4"
	"math/rand"
	"strconv"
	"strings"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetTokenFromContext(c echo.Context) string {
	return strings.ReplaceAll(c.Request().Header.Get("Authorization"), "Bearer ", "")
}

func ParseParamID(param string) uint {
	id, err := strconv.Atoi(param)
	if err != nil {
		return 0
	}
	return uint(id)
}
