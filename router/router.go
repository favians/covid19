package router

import (
	"net/http"

	"github.com/favians/golang_starter/api"
	"github.com/favians/golang_starter/api/middlewares"

	"github.com/labstack/echo"
	eMiddleware "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// router groups
	jwtGroup := e.Group("/jwt")

	// set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetCompleteLogMiddlware(e)
	middlewares.SetJwtAdminMiddlewares(jwtGroup)

	e.Use(eMiddleware.CORSWithConfig(eMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// set main routes
	api.MainGroup(e)

	// set token and auth routes
	api.TokenGroup(e)
	api.AuthGroup(e)

	// set group routes
	api.JwtGroup(jwtGroup)

	return e
}
