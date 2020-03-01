package middlewares

import (
	"encoding/json"

	"github.com/favians/golang_starter/bootstrap"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// SetCompleteLogMiddlware Middleware for logging request and response
func SetCompleteLogMiddlware(e *echo.Echo) {

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {

		var bodyJSON interface{}
		var bodyRESP interface{}

		json.Unmarshal(reqBody, &bodyJSON)
		json.Unmarshal(resBody, &bodyRESP)

		bootstrap.App.Log.LogRequest(c, bodyJSON, bodyRESP)
	}))
}
