package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/helpers"
	"github.com/zanamira43/appointment-api/models"
	"github.com/zanamira43/appointment-api/repository"
	"github.com/zanamira43/appointment-api/response"
)

type TimeTableHandler struct {
	TimeTableRepository *repository.GormTimeTableRepository
}

func NewTimeTableHandler(repo *repository.GormTimeTableRepository) *TimeTableHandler {
	return &TimeTableHandler{
		TimeTableRepository: repo,
	}
}

func (h *TimeTableHandler) GetTimeTables(c echo.Context) error {
	search := c.QueryParam("search")
	searchByWeekDays := c.QueryParam("searchByWeekDays")

	// Parse pagination parameters
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 5 // default limit
	}

	// Optional: Set maximum limit to prevent abuse
	if limit > 100 {
		limit = 100
	}

	if searchByWeekDays != "" {
		search = searchByWeekDays
	}

	datas, total, err := h.TimeTableRepository.GetAllTimeTables(page, limit, search)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Failed to get time tables")
	}

	res, err := response.Pagination(datas, total, page, limit)
	if err != nil {
		log.Error("Failed to Create Paginagtion", err)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *TimeTableHandler) CreateTimeTables(c echo.Context) error {
	var data dto.TimeTable

	// bind request body for appointment
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate the request body
	if err := helpers.ValidateTimeTables(&data); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	authHandler := NewAuth((*repository.GormUserRepository)(h.TimeTableRepository))
	user, err := authHandler.GetUserByCookie(c)
	if err != nil {
		log.Error("User Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "User not Found")
	}

	err = h.TimeTableRepository.CreateTimeTables(&data, user.ID)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, "Failed to create New TimeTable")
	}
	return c.JSON(http.StatusOK, data)
}

func (h *TimeTableHandler) GetTimeTable(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("invaild time table id", err.Error())
		return c.JSON(http.StatusBadRequest, "invaild time table id")
	}

	data, err := h.TimeTableRepository.GetTimeTableByID(id)
	if err != nil {
		log.Error("time table not found", err.Error())
		return c.JSON(http.StatusNotFound, "time table not found")
	}

	return c.JSON(http.StatusOK, data)
}

func (h *TimeTableHandler) UpdateTimeTable(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("invaild time table id", err.Error())
		return c.JSON(http.StatusBadRequest, "invaild time table id")
	}

	var data dto.TimeTable
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	authHandler := NewAuth((*repository.GormUserRepository)(h.TimeTableRepository))
	user, err := authHandler.GetUserByCookie(c)
	if err != nil {
		log.Error("User Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "User not Found")
	}

	update_data := new(models.TimeTable)
	update_data, err = h.TimeTableRepository.UpdateTimeTableByID(id, &data, user.ID)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update time table")
	}
	return c.JSON(http.StatusOK, update_data)
}

func (h *TimeTableHandler) DeleteTimeTable(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("invaild time table id", err.Error())
		return c.JSON(http.StatusBadRequest, "invaild time table id")
	}

	err = h.TimeTableRepository.DeleteTimeTableByID(id)
	if err != nil {
		log.Error("time table not found", err.Error())
		return c.JSON(http.StatusNotFound, "time table not found")
	}
	return c.NoContent(http.StatusNoContent)
}
