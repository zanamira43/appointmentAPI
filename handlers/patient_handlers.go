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

type PatientHandler struct {
	PatientRepository *repository.GormPatientRepository
}

func NewPatientHandler(patientRepo *repository.GormPatientRepository) *PatientHandler {
	return &PatientHandler{PatientRepository: patientRepo}
}

// Create New Patient
func (h *PatientHandler) CreatePatient(c echo.Context) error {
	patient := new(dto.Patient)

	if err := c.Bind(&patient); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate the request body
	if err := helpers.ValidatePatient(patient); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.PatientRepository.CreatePatient(patient)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to create patient")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": patient,
	})
}

// Get all patients
func (h *PatientHandler) GetAllPatients(c echo.Context) error {
	patients, err := h.PatientRepository.GetAllPatients()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get patients")
	}
	return c.JSON(http.StatusOK, patients)
}

// Get single patient by id
func (h *PatientHandler) GetPatient(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Patient Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Patient Id")
	}

	patient, err := h.PatientRepository.GetPatientByID(id)
	if err != nil {
		log.Error("Patient Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "Patient Not Found")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": patient,
	})
}

// Get single patient by slug
func (h *PatientHandler) GetPatientbySlug(c echo.Context) error {
	// slug := c.Param("slug")
	slug := c.QueryParam("slug")
	if slug == "" {
		log.Error("Invalid Patient Slug")
		return c.JSON(http.StatusBadRequest, "Invalid Patient Slug")
	}
	patient, err := h.PatientRepository.GetPatientBySlug(slug)
	if err != nil {
		log.Error("Patient Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "Patient Not Found")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": patient,
	})
}

// update single patient by id
func (h *PatientHandler) UpdatePatient(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Patient Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Patient Id")
	}

	var dtoPatient dto.Patient
	if err := c.Bind(&dtoPatient); err != nil {
		log.Error("Invalid Request data", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Request data")
	}

	patient := new(models.Patient)
	patient, err = h.PatientRepository.UpdatePatient(id, &dtoPatient)
	if err != nil {
		log.Error("Failed to update patient", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update patient")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": patient,
	})
}

// delete patient by id
func (h *PatientHandler) DeletePatient(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Patient Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Patient Id")
	}
	err = h.PatientRepository.DeletePatient(id)
	if err != nil {
		log.Error("patient not found", err.Error())
		return c.JSON(http.StatusInternalServerError, "patient not found")
	}
	return c.NoContent(http.StatusNoContent)
}
