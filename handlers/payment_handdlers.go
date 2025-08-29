package handlers

import (
	"math"
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

type PaymentHandler struct {
	PaymentRepository *repository.GormPaymentRepository
}

func NewPaymentHandler(Repo *repository.GormPaymentRepository) *PaymentHandler {
	return &PaymentHandler{PaymentRepository: Repo}
}

// Create New Payment
func (h *PaymentHandler) CreatePayments(c echo.Context) error {
	payment := new(dto.Payment)

	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate the request body
	if err := helpers.ValidatePayment(payment); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.PaymentRepository.CreatepPayments(payment)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to create payment")
	}
	return c.JSON(http.StatusOK, payment)
}

// Get all service types
func (h *PaymentHandler) GetPayments(c echo.Context) error {
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

	payments, total, err := h.PaymentRepository.GetPayments(page, limit)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get payments")
	}
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	hasNext := page < totalPages
	hasPrev := page > 1

	response := response.PaginatedResponse{
		Data:       payments,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}
	return c.JSON(http.StatusOK, response)
}

// Get single service type
func (h *PaymentHandler) GetPayment(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid payment Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid payment Id")
	}

	data, err := h.PaymentRepository.GetPayment(id)
	if err != nil {
		log.Error("data Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "data Not Found")
	}
	return c.JSON(http.StatusOK, data)
}

// update single service by id
func (h *PaymentHandler) UpdatePayment(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid payment Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid payment Id")
	}

	var dto dto.Payment
	if err := c.Bind(&dto); err != nil {
		log.Error("Invalid Request data", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Request data")
	}

	// Validate the request body
	if err := helpers.ValidatePayment(&dto); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	data := new(models.Payment)
	data, err = h.PaymentRepository.UpdatePayment(id, &dto)
	if err != nil {
		log.Error("Failed to update payment", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update payment ")
	}
	return c.JSON(http.StatusOK, data)
}

// delete service type by id
func (h *PaymentHandler) DeletePayment(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid payment Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid payment Id")
	}
	err = h.PaymentRepository.DeletePayment(id)
	if err != nil {
		log.Error("data not found", err.Error())
		return c.JSON(http.StatusNotFound, "data not found")
	}
	return c.NoContent(http.StatusNoContent)
}
