package api

import (
	"github.com/favians/golang_starter/api/handlers"
	"github.com/favians/golang_starter/api/middlewares"

	"github.com/labstack/echo"
)

func TokenGroup(e *echo.Echo) {

	//API Group For Admin
	g := e.Group("/users")
	middlewares.SetJwtMiddlewares(g)
	g.GET("", handlers.GetAdminById)
	g.GET("/list", handlers.GetAdmins)
	g.POST("", handlers.AddAdmin)
	g.PUT("", handlers.EditAdmin)
	g.DELETE("", handlers.DeleteAdmin)

	//API Group For SuperAdmin
	g = e.Group("admin/users")
	middlewares.SetJwtAdminMiddlewares(g)
	g.GET("", handlers.GetAdminById)
	g.GET("/list", handlers.GetAdmins)
	g.POST("", handlers.AddAdmin)
	g.PUT("", handlers.EditAdmin)
	g.DELETE("", handlers.DeleteAdmin)

	//API Group For Either Admin or SuperAdmin
	g = e.Group("/pasien")
	middlewares.SetJwtGeneralMiddlewares(g)
	g.GET("", handlers.GetPasienById)
	g.GET("/list", handlers.GetPasien)
	g.POST("", handlers.AddPasien)
	g.PUT("", handlers.EditPasien)
	g.DELETE("", handlers.DeletePasien)

	//API Group For Either Admin or SuperAdmin
	g = e.Group("/rumahsakit")
	middlewares.SetJwtGeneralMiddlewares(g)
	g.GET("", handlers.GetRumahSakitById)
	g.GET("/list", handlers.GetRumahSakit)
	g.POST("", handlers.AddRumahSakit)
	g.PUT("", handlers.EditRumahSakit)
	g.DELETE("", handlers.DeleteRumahSakit)

	//API Group For Either Admin or SuperAdmin
	g = e.Group("/report")
	middlewares.SetJwtGeneralMiddlewares(g)
	g.GET("", handlers.GetReportById)
	g.GET("/list", handlers.GetReport)
	g.POST("", handlers.AddReport)
	g.PUT("", handlers.EditReport)
	g.DELETE("", handlers.DeleteReport)

}
