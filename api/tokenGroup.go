package api

import (
	"github.com/favians/golang_starter/api/handlers"
	"github.com/favians/golang_starter/api/middlewares"

	"github.com/labstack/echo"
)

func TokenGroup(e *echo.Echo) {

	//API Group For Normal User
	g := e.Group("/users")
	middlewares.SetJwtMiddlewares(g)
	g.GET("", handlers.GetUserById)
	g.GET("/list", handlers.GetUsers)
	g.POST("", handlers.AddUser)
	g.PUT("", handlers.EditUser)
	g.DELETE("", handlers.DeleteUser)

	//API Group For Admin
	g = e.Group("admin/users")
	middlewares.SetJwtAdminMiddlewares(g)
	g.GET("", handlers.GetUserById)
	g.GET("/list", handlers.GetUsers)
	g.POST("", handlers.AddUser)
	g.PUT("", handlers.EditUser)
	g.DELETE("", handlers.DeleteUser)

	//API Group For Either Admin or User
	g = e.Group("general/users")
	middlewares.SetJwtGeneralMiddlewares(g)
	g.GET("", handlers.GetUserById)
	g.GET("/list", handlers.GetUsers)
	g.POST("", handlers.AddUser)
	g.PUT("", handlers.EditUser)
	g.DELETE("", handlers.DeleteUser)
}
