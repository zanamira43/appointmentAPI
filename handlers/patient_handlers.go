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

type PatientHandler struct {
	PatientRepository *repository.GormPatientRepository
}

func NewPatientHandler(patientRepo *repository.GormPatientRepository) *PatientHandler {
	return &PatientHandler{PatientRepository: patientRepo}
}

// Create New Patient
func (h *PatientHandler) CreatePatient(c echo.Context) error {
	patient := new(dto.Patient)

	// Bind the request body to the patient struct

	if err := c.Bind(&patient); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate the request body
	if err := helpers.ValidatePatient(patient); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	auth := NewAuth((*repository.GormUserRepository)(h.PatientRepository))
	user, err := auth.GetUserByCookie(c)
	if err != nil {
		log.Error("User Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "User not Found")
	}

	err = h.PatientRepository.CreatePatient(patient, user.ID)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to create patient")
	}
	return c.JSON(http.StatusOK, patient)
}

// Get all patients
func (h *PatientHandler) GetAllPatients(c echo.Context) error {
	// parse Search parameters
	search := c.QueryParam("search")
	searchByCode := c.QueryParam("searchByCode")

	if searchByCode != "" {
		search = searchByCode
	}

	// Parse pagination parameters
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	// Optional: Set maximum limit to prevent abuse
	if limit > 100 {
		limit = 100
	}

	auth := NewAuth((*repository.GormUserRepository)(h.PatientRepository))
	user, err := auth.GetUserByCookie(c)
	if err != nil {
		log.Error("User Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "User not Found")
	}

	patients, total, err := h.PatientRepository.GetAllPatients(page, limit, search, user.ID, user.Role)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get patients")
	}

	// Calculate pagination metadata (only if paginated)
	var totalPages int
	var hasNext, hasPrev bool

	if page > 0 && limit > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(limit)))
		hasNext = page < totalPages
		hasPrev = page > 1
	}
	// Return the response
	responses := response.PaginatedResponse{
		Data:       patients,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: int(total / int64(limit)),
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}
	return c.JSON(http.StatusOK, responses)
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

	return c.JSON(http.StatusOK, patient)
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
	return c.JSON(http.StatusOK, patient)
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

	// Validate the request body
	if err := helpers.ValidatePatient(&dtoPatient); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	auth := NewAuth((*repository.GormUserRepository)(h.PatientRepository))
	user, err := auth.GetUserByCookie(c)
	if err != nil {
		log.Error("User Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "User not Found")
	}

	patient := new(models.Patient)
	patient, err = h.PatientRepository.UpdatePatient(id, &dtoPatient, user.ID)
	if err != nil {
		log.Error("Failed to update patient", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update patient")
	}
	return c.JSON(http.StatusOK, patient)
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

// patient outcome
func (h *PatientHandler) Outcome(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Patient Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Patient Id")
	}

	patientOutcomeResponse, err := h.PatientRepository.PatinetOutcome(id)

	if err != nil {
		log.Error("Patient Not Found", err.Error())
		return c.JSON(http.StatusBadRequest, "Patient Not Found")
	}

	return c.JSON(http.StatusOK, patientOutcomeResponse)
}
