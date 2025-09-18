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
	api.POST("/user/updateinfo", authHandler.UpdateInfo)
	api.POST("/user/updatepassword", authHandler.UpdatePassword)
	api.POST("/user/logout", authHandler.Logout)

	// patient routes
	patientRepo := repository.NewGormPatientRepository(database.DB)
	patientHandler := handlers.NewPatientHandler(patientRepo)
	api.POST("/patients", patientHandler.CreatePatient)
	api.GET("/patients", patientHandler.GetAllPatients)
	api.GET("/patients/:id", patientHandler.GetPatient)
	api.PUT("/patients/:id", patientHandler.UpdatePatient)
	api.DELETE("/patients/:id", patientHandler.DeletePatient)
	api.GET("/patients/search", patientHandler.GetPatientbySlug)

	// session routes end points
	sessionRepo := repository.NewGormSessionRepository(database.DB)
	sessionHandler := handlers.NewSessionHandler(sessionRepo)
	api.POST("/sessions", sessionHandler.CreateSessions)
	api.GET("/sessions", sessionHandler.GetSessions)
	api.GET("/sessions/:id", sessionHandler.GetSession)
	api.PUT("/sessions/:id", sessionHandler.UpdateSession)
	api.DELETE("/sessions/:id", sessionHandler.DeleteSession)

	// payment routes end points
	PaymentRepo := repository.NewGormPaymentRepository(database.DB)
	PaymentHandler := handlers.NewPaymentHandler(PaymentRepo)
	api.POST("/payments", PaymentHandler.CreatePayments)
	api.GET("/payments", PaymentHandler.GetPayments)
	api.GET("/payments/:id", PaymentHandler.GetPayment)
	api.PUT("/payments/:id", PaymentHandler.UpdatePayment)
	api.DELETE("/payments/:id", PaymentHandler.DeletePayment)

	// user routes end points
	userRepo = repository.NewGormUserRepository(database.DB)
	userHandler := handlers.NewUserHandler(userRepo)
	api.POST("/users", userHandler.CreateUser, middleware.IsUserAdmin)
	api.GET("/users", userHandler.GetAllUsers, middleware.IsUserAdmin)
	api.GET("/users/:id", userHandler.GetUser, middleware.IsUserAdmin)
	api.PUT("/users/:id", userHandler.UpdateUser, middleware.IsUserAdmin)
	api.DELETE("/users/:id", userHandler.DeleteUser, middleware.IsUserAdmin)

	// appointment schedule routes end points
	timeTableRepo := repository.NewGormTimeTableRepository(database.DB)
	timeTableHandler := handlers.NewTimeTableHandler(timeTableRepo)
	api.GET("/timetables", timeTableHandler.GetTimeTables)
	api.POST("/timetables", timeTableHandler.CreateTimeTables)
	api.GET("/timetables/:id", timeTableHandler.GetTimeTable)
	api.PUT("/timetables/:id", timeTableHandler.UpdateTimeTable)
	api.DELETE("/timetables/:id", timeTableHandler.DeleteTimeTable)
}
