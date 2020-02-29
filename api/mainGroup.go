package api

import (
	"github.com/favians/golang_starter/api/handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/yallo", handlers.Yallo)

	//cats handler
	e.GET("/cat/:data", handlers.GetCats)
	e.POST("/cat", handlers.AddCat)

	//dog handler
	e.POST("/dogs", handlers.AddDog)

	//hamster handler
	e.POST("/hamsters", handlers.AddHamster)
}
