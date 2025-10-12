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

type PaymentTypeHandler struct {
	PaymentTypeRepository *repository.GormPaymentTypeRepository
}

func NewPaymentTypeHandler(Repo *repository.GormPaymentTypeRepository) *PaymentTypeHandler {
	return &PaymentTypeHandler{PaymentTypeRepository: Repo}
}

// Create New Payment Type
func (h *PaymentTypeHandler) CreatePaymentType(c echo.Context) error {
	paymentType := new(dto.PaymentType)

	if err := c.Bind(&paymentType); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate the request body
	if err := helpers.ValidatePaymentType(paymentType); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.PaymentTypeRepository.CreatePaymentType(paymentType)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to create payment type")
	}
	return c.JSON(http.StatusOK, paymentType)
}

// Get all payment types
func (h *PaymentTypeHandler) GetPaymentTypes(c echo.Context) error {

	search := c.QueryParam("search")

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

	paymentTypes, total, err := h.PaymentTypeRepository.GetPaymentTypes(page, limit, search)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get payment types")
	}

	// Calculate pagination metadata (only if paginated)
	var totalPages int
	var hasNext, hasPrev bool

	if page > 0 && limit > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(limit)))
		hasNext = page < totalPages
		hasPrev = page > 1
	}

	response := response.PaginatedResponse{
		Data:       paymentTypes,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: int(total / int64(limit)),
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}
	return c.JSON(http.StatusOK, response)
}

// Get single payment type
func (h *PaymentTypeHandler) GetPaymentType(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid payment type Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid payment type Id")
	}

	data, err := h.PaymentTypeRepository.GetPaymentType(id)
	if err != nil {
		log.Error("data Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "data Not Found")
	}
	return c.JSON(http.StatusOK, data)
}

// update single payment type by id
func (h *PaymentTypeHandler) UpdatePaymentType(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid payment type Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid payment type Id")
	}

	var dto dto.PaymentType
	if err := c.Bind(&dto); err != nil {
		log.Error("Invalid Request data", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Request data")
	}

	// Validate the request body
	if err := helpers.ValidatePaymentType(&dto); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	data := new(models.PaymentType)
	data, err = h.PaymentTypeRepository.UpdatePaymentType(id, &dto)
	if err != nil {
		log.Error("Failed to update payment type", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update payment type")
	}
	return c.JSON(http.StatusOK, data)
}

// delete payment type by id
func (h *PaymentTypeHandler) DeletePaymentType(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid payment type Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid payment type Id")
	}
	err = h.PaymentTypeRepository.DeletePaymentType(id)
	if err != nil {
		log.Error("data not found", err.Error())
		return c.JSON(http.StatusNotFound, "data not found")
	}
	return c.NoContent(http.StatusNoContent)
}
