package router

import (
	"github.com/favians/golang_starter/api"
	"github.com/favians/golang_starter/api/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// router groups
	jwtGroup := e.Group("/jwt")

	// set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetCompleteLogMiddlware(e)

	middlewares.SetJwtAdminMiddlewares(jwtGroup)

	// set main routes
	api.MainGroup(e)

	// set token and auth routes
	api.TokenGroup(e)
	api.AuthGroup(e)

	// set group routes
	api.JwtGroup(jwtGroup)

	return e
}
