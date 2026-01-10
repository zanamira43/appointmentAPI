package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/helpers"
	"github.com/zanamira43/appointment-api/repository"
	"github.com/zanamira43/appointment-api/response"
)

type PersonInfoHandler struct {
	PersonInfoRepository *repository.GormPersonInfoRepository
}

func NewPersonInfoHandler(personRepo *repository.GormPersonInfoRepository) *PersonInfoHandler {
	return &PersonInfoHandler{PersonInfoRepository: personRepo}
}

// create new person info
func (h *PersonInfoHandler) CreatePersonInfo(c echo.Context) error {
	newPerson := new(dto.PersonInfo)
	if err := c.Bind(&newPerson); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := helpers.ValidatePersonInfo(newPerson); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	person := &dto.PersonInfo{
		FullName:    newPerson.FullName,
		PhoneNumber: newPerson.PhoneNumber,
	}

	err := h.PersonInfoRepository.CreatePersonInfo(person)
	if err != nil {
		log.Error("Failed to register new person info", err.Error())
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to register new person info",
		})
	}
	return c.JSON(http.StatusOK, person)
}

// get all person info
func (h *PersonInfoHandler) GetAllPersonInfo(c echo.Context) error {
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

	datas, total, err := h.PersonInfoRepository.GetAllPersonInfo(page, limit, search)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get person info")
	}

	res, err := response.Pagination(datas, total, page, limit)
	if err != nil {
		log.Error("Failed to Create Paginagtion", err)
	}
	return c.JSON(http.StatusOK, res)
}

// get person info by id
func (h *PersonInfoHandler) GetPersonInfo(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Person Info Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Person Info Id")
	}
	person, err := h.PersonInfoRepository.GetPersonInfoByID(id)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Person Info Not Found")
	}
	return c.JSON(http.StatusOK, person)
}

// update
func (h *PersonInfoHandler) UpdatePersonInfo(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Person Info Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Person Info Id")
	}

	var data dto.PersonInfo
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	person, err := h.PersonInfoRepository.UpdatePersonInfo(id, &data)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Person Info Not Found")
	}

	return c.JSON(http.StatusOK, person)
}

// delete person info
func (h *PersonInfoHandler) DeletePersonInfo(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Person Info Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Person Info Id")
	}
	err = h.PersonInfoRepository.DeletePersonInfo(id)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Person Info Not Found")
	}
	return c.NoContent(http.StatusNoContent)
}
