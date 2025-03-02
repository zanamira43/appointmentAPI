package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zanamira43/appointment-api/database"
	"github.com/zanamira43/appointment-api/routes"
)

func main() {
	// Echo instance initialize API
	e := echo.New()

	// loading env
	err := godotenv.Load()
	if err != nil {
		log.Printf("Unable to load .env file %v", err)
	}

	// initialize database
	err = database.Connect()
	if err != nil {
		log.Fatalf("Unable to connect to database %v\n", err)
		panic("failed to connect database\n")
	}

	// deal with the middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("2MB"))

	// deal with cors for frontend
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "http://*", "https://*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	// setup routes
	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT")))
}
