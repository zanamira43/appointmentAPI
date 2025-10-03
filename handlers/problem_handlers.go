package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/helpers"
	"github.com/zanamira43/appointment-api/repository"
)

type ProblemHandler struct {
	ProblemRepository *repository.GormProblemRepository
}

func NewProblemHandler(problemRepository *repository.GormProblemRepository) *ProblemHandler {
	return &ProblemHandler{ProblemRepository: problemRepository}
}

// get all problems
func (ph *ProblemHandler) GetProblems(c echo.Context) error {
	problems, err := ph.ProblemRepository.GetAllProblems()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, problems)
}

// get  single problem by ID
func (ph *ProblemHandler) GetProblem(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Problem Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Problem Id")
	}

	problem, err := ph.ProblemRepository.GetProblemByID(id)
	if err != nil {
		log.Error("Problem Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "Problem Not Found")
	}
	return c.JSON(http.StatusOK, problem)
}

// get single problem by patient Id
func (ph *ProblemHandler) GetProblemByPatientId(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Problem Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Problem Id")
	}

	problem, err := ph.ProblemRepository.GetProblemByPatientId(id)
	if err != nil {
		log.Error("Problem Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "Problem Not Found")
	}
	return c.JSON(http.StatusOK, problem)
}

func (ph *ProblemHandler) CreateProblem(c echo.Context) error {
	var data dto.Problem
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate the request body
	if err := helpers.ValidateProblems(&data); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// create problem from Gorm repository method
	err := ph.ProblemRepository.CreateProblem(&data)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to create problem")
	}
	return c.JSON(http.StatusOK, data)
}

func (ph *ProblemHandler) UpdateProblem(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Problem Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Problem Id")
	}

	var data dto.Problem
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	problem, err := ph.ProblemRepository.UpdateProblemByID(id, &data)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update problem")
	}
	return c.JSON(http.StatusOK, problem)
}

// delete single problem by id handler
func (ph *ProblemHandler) DeleteProblem(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Problem Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Problem Id")
	}
	err = ph.ProblemRepository.DeleteProblemByID(id)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to delete problem")
	}
	return c.NoContent(http.StatusNoContent)
}

// delete patient image
// delete image url & rempove it in ropo
func (ph *ProblemHandler) DeletePatientImageUrl(c echo.Context) error {
	var image dto.Image

	if err := c.Bind(&image); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if image.PatientImageUrl == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "patientImageUrl is required"})
	}

	if err := helpers.DeleteImageFromStorage(image.PatientImageUrl); err != nil {
		// Map specific errors to status codes if needed
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusOK)
}
