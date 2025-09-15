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

type AppointmentHandler struct {
	AppointmentRepository *repository.GormAppointmentRepository
}

func NewAppointmentHandler(repo *repository.GormAppointmentRepository) *AppointmentHandler {
	return &AppointmentHandler{
		AppointmentRepository: repo,
	}
}

func (h *AppointmentHandler) GetAppointments(c echo.Context) error {
	search := c.QueryParam("search")

	// Parse pagination parameters
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10 // default limit
	}

	// Optional: Set maximum limit to prevent abuse
	if limit > 100 {
		limit = 100
	}

	appointments, total, err := h.AppointmentRepository.GetAllAppointments(page, limit, search)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Failed to get appointments")
	}

	res, err := response.Pagination(appointments, total, page, limit)
	if err != nil {
		log.Error("Failed to Create Paginagtion", err)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *AppointmentHandler) CreateAppointments(c echo.Context) error {
	var data dto.AppointmentSchedule

	// bind request body for appointment
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate the request body
	if err := helpers.ValidateAppointments(&data); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	authHandler := NewAuth((*repository.GormUserRepository)(h.AppointmentRepository))
	user, err := authHandler.GetUserByCookie(c)
	if err != nil {
		log.Error("User Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "User not Found")
	}

	err = h.AppointmentRepository.CreateAppointments(&data, user.ID)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, "Failed to create New Appointment")
	}
	return c.JSON(http.StatusOK, data)
}

func (h *AppointmentHandler) GetAppointment(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("invaild appointment id", err.Error())
		return c.JSON(http.StatusBadRequest, "invaild appointment id")
	}

	appointment, err := h.AppointmentRepository.GetAppointmentByID(id)
	if err != nil {
		log.Error("appointment not found", err.Error())
		return c.JSON(http.StatusNotFound, "appointment not found")
	}

	return c.JSON(http.StatusOK, appointment)
}

func (h *AppointmentHandler) UpdateAppointment(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("invaild appointment id", err.Error())
		return c.JSON(http.StatusBadRequest, "invaild appointment id")
	}

	var data dto.AppointmentSchedule
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	authHandler := NewAuth((*repository.GormUserRepository)(h.AppointmentRepository))
	user, err := authHandler.GetUserByCookie(c)
	if err != nil {
		log.Error("User Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "User not Found")
	}

	appointment := new(models.AppointmentSchedule)
	appointment, err = h.AppointmentRepository.UpdateAppointmentByID(id, &data, user.ID)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update appointment")
	}
	return c.JSON(http.StatusOK, appointment)
}

func (h *AppointmentHandler) DeleteAppointment(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("invaild appointment id", err.Error())
		return c.JSON(http.StatusBadRequest, "invaild appointment id")
	}

	err = h.AppointmentRepository.DeleteAppointmentByID(id)
	if err != nil {
		log.Error("appointment not found", err.Error())
		return c.JSON(http.StatusNotFound, "appointment not found")
	}
	return c.NoContent(http.StatusNoContent)
}
