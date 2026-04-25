package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zanamira43/appointment-api/repository"
)

type DashboardHnalder struct {
	Repo repository.GormDashboardRepository
}

func NewDashboardHandler(repo *repository.GormDashboardRepository) *DashboardHnalder {
	return &DashboardHnalder{Repo: *repo}
}

func (h DashboardHnalder) Outcome(c echo.Context) error {

	outcome, err := h.Repo.DashboardOutCome()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, outcome)
}
