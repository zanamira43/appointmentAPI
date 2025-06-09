package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/helpers"
	"github.com/zanamira43/appointment-api/models"
	"github.com/zanamira43/appointment-api/repository"
)

type ServiceTypeHandler struct {
	ServiceTypeRepository *repository.GormServiceTypeRepository
}

func NewServiceTypeHandler(ServiceTypeRepo *repository.GormServiceTypeRepository) *ServiceTypeHandler {
	return &ServiceTypeHandler{ServiceTypeRepository: ServiceTypeRepo}
}

// Create New Service Type
func (h *ServiceTypeHandler) CreateServiceTypes(c echo.Context) error {
	serviceType := new(dto.ServiceType)

	if err := c.Bind(&serviceType); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate the request body
	if err := helpers.ValidateServiceType(serviceType); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.ServiceTypeRepository.CreateServiceTypes(serviceType)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to create service type")
	}
	return c.JSON(http.StatusOK, serviceType)
}

// Get all service types
func (h *ServiceTypeHandler) GetAllServiceTypes(c echo.Context) error {
	serviceTypes, err := h.ServiceTypeRepository.GetAllServiceTypes()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get service types")
	}
	return c.JSON(http.StatusOK, serviceTypes)
}

// Get single service type
func (h *ServiceTypeHandler) GetServiceType(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid service type Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid service type Id")
	}

	data, err := h.ServiceTypeRepository.GetServiceTypeByID(id)
	if err != nil {
		log.Error("service type Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "service Not Found")
	}
	return c.JSON(http.StatusOK, data)
}

// update single service by id
func (h *ServiceTypeHandler) UpdateServiceType(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Service Type Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Service Type Id")
	}

	var dto dto.ServiceType
	if err := c.Bind(&dto); err != nil {
		log.Error("Invalid Request data", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Request data")
	}

	// Validate the request body
	if err := helpers.ValidateServiceType(&dto); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	data := new(models.ServiceType)
	data, err = h.ServiceTypeRepository.UpdateServiceTypeByID(id, &dto)
	if err != nil {
		log.Error("Failed to update service type", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update service type")
	}
	return c.JSON(http.StatusOK, data)
}

// delete service type by id
func (h *ServiceTypeHandler) DeleteServiceType(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid service type Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalidservice typer Id")
	}
	err = h.ServiceTypeRepository.DeleteServiceTypeByID(id)
	if err != nil {
		log.Error("service type not found", err.Error())
		return c.JSON(http.StatusInternalServerError, "service type not found")
	}
	return c.NoContent(http.StatusNoContent)
}
