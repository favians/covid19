package api

import (
	"github.com/favians/golang_starter/api/handlers"

	"github.com/labstack/echo"
)

func JwtGroup(g *echo.Group) {
	g.GET("/main", handlers.MainJwt)
}
