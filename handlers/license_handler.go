package handlers

import (
	"math"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type License struct {
	Status     string    `json:"status"`
	StartDate  time.Time `json:"start_date"`
	ExpiryDate time.Time `json:"expiry_date"`
	Usage      int       `json:"usage"`
	DaysLeft   int       `json:"days_left"`
}

// status fucntion
// if active, have not started yet or expired
func calculateStatus(now, start, expiry time.Time) string {
	if now.Before(start) {
		return "Not Started"
	}

	if now.After(expiry) {
		return "Expired"
	}

	return "Active"
}

// how many  % numbers used
func calculateUsage(now, start, expiry time.Time) int {
	total := expiry.Sub(start).Seconds()
	elapsed := now.Sub(start).Seconds()

	if elapsed < 0 {
		return 0
	}

	if elapsed > total {
		return 100
	}

	percent := (elapsed / total) * 100
	return int(math.Round(percent))
}

// how many days left to become expire
func calculateDaysLeft(now, expiry time.Time) int {
	if now.After(expiry) {
		return 0
	}

	return int(expiry.Sub(now).Hours() / 24)
}
func LicenseHandler(c echo.Context) error {

	start := time.Date(2025, 12, 1, 0, 0, 0, 0, time.UTC)
	expiry := time.Date(2026, 12, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	licens := License{
		Status:     calculateStatus(now, start, expiry),
		StartDate:  start,
		ExpiryDate: expiry,
		Usage:      calculateUsage(now, start, expiry),
		DaysLeft:   calculateDaysLeft(now, expiry),
	}

	return c.JSON(http.StatusOK, licens)
}
