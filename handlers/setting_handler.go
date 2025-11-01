package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/repository"
)

type SettingHandler struct {
	Repo *repository.GormSettingsRepository
}

func NewSettingHandler(repo *repository.GormSettingsRepository) SettingHandler {
	return SettingHandler{Repo: repo}
}

// get setting
func (h SettingHandler) GetSetting(c echo.Context) error {

	setting, err := h.Repo.GetSetting()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to fetch setting")
	}
	return c.JSON(http.StatusOK, setting)
}

// update setting
// get setting
func (h SettingHandler) UpdateSetting(c echo.Context) error {
	var settingDto dto.Settings

	if err := c.Bind(&settingDto); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	setting, err := h.Repo.UpdateSetting(&settingDto)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusNotFound, "Failed to fetch setting")
	}
	return c.JSON(http.StatusOK, setting)
}
