package api

import (
	"github.com/favians/golang_starter/api/handlers"

	"github.com/labstack/echo"
)

func AuthGroup(e *echo.Echo) {
	e.GET("/login", handlers.LoginUser)
	e.GET("/login/admin", handlers.LoginSuperAdmin)
}
