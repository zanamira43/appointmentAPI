package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/zanamira43/appointment-api/database"
	"github.com/zanamira43/appointment-api/handlers"
	"github.com/zanamira43/appointment-api/helpers"
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
	api.GET("/patients/:id/outcome", patientHandler.Outcome)

	// user routes end points
	userRepo = repository.NewGormUserRepository(database.DB)
	userHandler := handlers.NewUserHandler(userRepo)
	api.POST("/users", userHandler.CreateUser, middleware.IsUserAdmin)
	api.GET("/users", userHandler.GetAllUsers, middleware.IsUserAdmin)
	api.GET("/users/:id", userHandler.GetUser, middleware.IsUserAdmin)
	api.PUT("/users/:id", userHandler.UpdateUser, middleware.IsUserAdmin)
	api.PUT("/users/:id/password", userHandler.UpdateUserPassword, middleware.IsUserAdmin)
	api.DELETE("/users/:id", userHandler.DeleteUser, middleware.IsUserAdmin)

	// appointment schedule routes end points
	timeTableRepo := repository.NewGormTimeTableRepository(database.DB)
	timeTableHandler := handlers.NewTimeTableHandler(timeTableRepo)
	api.GET("/timetables", timeTableHandler.GetTimeTables)
	api.POST("/timetables", timeTableHandler.CreateTimeTables)
	api.GET("/timetables/:id", timeTableHandler.GetTimeTable)
	api.PUT("/timetables/:id", timeTableHandler.UpdateTimeTable)
	api.DELETE("/timetables/:id", timeTableHandler.DeleteTimeTable)

	// problem  routes end points
	problemRepo := repository.NewGormProblemRepository(database.DB)
	problemHandler := handlers.NewProblemHandler(problemRepo)
	api.POST("/problems", problemHandler.CreateProblem, middleware.IsUserAdmin)
	api.GET("/problems", problemHandler.GetProblems, middleware.IsUserAdmin)
	api.GET("/problems/:id", problemHandler.GetProblem, middleware.IsUserAdmin)
	api.GET("/problems/patient/:id", problemHandler.GetProblemByPatientId, middleware.IsUserAdmin)
	api.PUT("/problems/:id", problemHandler.UpdateProblem, middleware.IsUserAdmin)
	api.DELETE("/problems/:id", problemHandler.DeleteProblem, middleware.IsUserAdmin)

	// patient images end points
	// get image
	api.Static("/image", "./uploads")
	// upload image route endpoints for categories
	api.POST("/image/upload", helpers.UploadImage, middleware.IsUserAdmin)
	// remove patient image from storage
	api.POST("/image/delete", problemHandler.DeletePatientImageUrl, middleware.IsUserAdmin)

	// session routes end points
	sessionRepo := repository.NewGormSessionRepository(database.DB)
	sessionHandler := handlers.NewSessionHandler(sessionRepo)
	api.POST("/sessions", sessionHandler.CreateSessions)
	api.GET("/sessions", sessionHandler.GetSessions)
	api.GET("/sessions/patient/:id", sessionHandler.GetSessionsByPatientId)
	api.GET("/sessions/:id", sessionHandler.GetSession)
	api.PUT("/sessions/:id", sessionHandler.UpdateSession)
	api.DELETE("/sessions/:id", sessionHandler.DeleteSession)

	// payment routes end points
	PaymentRepo := repository.NewGormPaymentRepository(database.DB)
	PaymentHandler := handlers.NewPaymentHandler(PaymentRepo)
	api.POST("/payments", PaymentHandler.CreatePayments)
	api.GET("/payments", PaymentHandler.GetPayments)
	api.GET("/payments/patient/:id", PaymentHandler.GetPaymentsByPatientId)
	api.GET("/payments/:id", PaymentHandler.GetPayment)
	api.PUT("/payments/:id", PaymentHandler.UpdatePayment)
	api.DELETE("/payments/:id", PaymentHandler.DeletePayment)

	// payment type routes end points
	PaymentTypeRepo := repository.NewGormPaymentTypeRepository(database.DB)
	PaymentTypeHandler := handlers.NewPaymentTypeHandler(PaymentTypeRepo)
	api.POST("/payment-types", PaymentTypeHandler.CreatePaymentType)
	api.GET("/payment-types", PaymentTypeHandler.GetPaymentTypes)
	api.GET("/payment-types/:id", PaymentTypeHandler.GetPaymentType)
	api.PUT("/payment-types/:id", PaymentTypeHandler.UpdatePaymentType)
	api.DELETE("/payment-types/:id", PaymentTypeHandler.DeletePaymentType)

	// setting routes end points
	settingRepo := repository.NewGormSettingsRepository(database.DB)
	settingHandler := handlers.NewSettingHandler(&settingRepo)
	api.GET("/system/setting", settingHandler.GetSetting)
	api.PUT("/system/setting", settingHandler.UpdateSetting)

	// notebook routes endPoint
	notebookRepo := repository.NewGormNoteBookRepostiry(database.DB)
	notebookHandler := handlers.NewNoteBookHandler(notebookRepo)
	api.POST("/notebooks", notebookHandler.CreateNotebook, middleware.IsUserAdmin)
	api.GET("/notebooks", notebookHandler.GetAllNotebooks, middleware.IsUserAdmin)
	api.GET("/notebooks/:id", notebookHandler.GetNotebook, middleware.IsUserAdmin)
	api.PUT("/notebooks/:id", notebookHandler.UpdateNotebook, middleware.IsUserAdmin)
	api.DELETE("/notebooks/:id", notebookHandler.DeleteNotebook, middleware.IsUserAdmin)

	// person info for 2 month plan routes endPoint
	personInfoRepo := repository.NewGormPersonInfoRepository(database.DB)
	personInfoHandler := handlers.NewPersonInfoHandler(personInfoRepo)
	api.POST("/persons", personInfoHandler.CreatePersonInfo)
	api.GET("/persons", personInfoHandler.GetAllPersonInfo)
	api.GET("/persons/:id", personInfoHandler.GetPersonInfo)
	api.PUT("/persons/:id", personInfoHandler.UpdatePersonInfo)
	api.DELETE("/persons/:id", personInfoHandler.DeletePersonInfo)

}
