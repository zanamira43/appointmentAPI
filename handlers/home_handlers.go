package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// home funct
func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, &echo.Map{
		"1) Mesage":   "Welcome to Clinic Appoitment API - Project by Golang & Echo Framework",
		"2) Base URL": c.Request().Host + "/api/",
	})

}
