package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/zanamira43/appointment-api/database"
	"github.com/zanamira43/appointment-api/handlers"
	"github.com/zanamira43/appointment-api/middleware"
	"github.com/zanamira43/appointment-api/repository"
)

func SetupRoutes(app *echo.Echo) {
	// home route end point
	app.GET("/", handlers.Home)

	api := app.Group("/api")

	// patient routes
	patientRepo := repository.NewGormPatientRepository(database.DB)
	patientHandler := handlers.NewPatientHandler(patientRepo)
	api.POST("/patients", patientHandler.CreatePatient)
	api.GET("/patients", patientHandler.GetAllPatients)
	api.GET("/patients/:id", patientHandler.GetPatient)
	api.PUT("/patients/:id", patientHandler.UpdatePatient)
	api.DELETE("/patients/:id", patientHandler.DeletePatient)
	api.GET("/patients/search", patientHandler.GetPatientbySlug)

	// appointment routes
	// user endpoints
	userRepo := repository.NewGormUserRepository(database.DB)
	authHandler := handlers.NewAuth(userRepo)
	api.POST("/signup", authHandler.SignUP)
	api.POST("/login", authHandler.Login)

	// middleware instance
	middleware := middleware.NewMiddleware()
	api.Use(middleware.IsAuthenticated)
	api.GET("/user/info", authHandler.User)
	api.GET("/user/updateinfo", authHandler.UpdateInfo)
	api.GET("/user/updatepassword", authHandler.UpdatePassword)
	api.POST("/user/logout", authHandler.Logout)
}
