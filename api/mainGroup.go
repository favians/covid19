package api

import (
	"rest_echo/api/handlers"
	"rest_echo/api/middlewares"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/login", handlers.LoginUser)
	e.GET("/login/admin", handlers.LoginAdmin)

	e.GET("/yallo", handlers.Yallo)
	e.GET("/cats/:data", handlers.GetCats)

	e.POST("/cats", handlers.AddCat)
	e.POST("/dogs", handlers.AddDog)
	e.POST("/hamsters", handlers.AddHamster)

	g := e.Group("/users")
	middlewares.SetJwtMiddlewares(g)
	g.GET("", handlers.GetUserById)
	g.GET("/list", handlers.GetUsers)
	g.POST("", handlers.AddUser)
	g.PUT("", handlers.EditUser)
	g.DELETE("", handlers.DeleteUser)

	g = e.Group("admin/users")
	middlewares.SetJwtAdminMiddlewares(g)
	g.GET("", handlers.GetUserById)
	g.GET("/list", handlers.GetUsers)
	g.POST("", handlers.AddUser)
	g.PUT("", handlers.EditUser)
	g.DELETE("", handlers.DeleteUser)

	g = e.Group("general/users")
	middlewares.SetJwtGeneralMiddlewares(g)
	g.GET("", handlers.GetUserById)
	g.GET("/list", handlers.GetUsers)
	g.POST("", handlers.AddUser)
	g.PUT("", handlers.EditUser)
	g.DELETE("", handlers.DeleteUser)
}
