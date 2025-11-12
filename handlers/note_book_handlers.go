package handlers

import (
	"math"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/helpers"
	"github.com/zanamira43/appointment-api/models"
	"github.com/zanamira43/appointment-api/repository"
	"github.com/zanamira43/appointment-api/response"
)

type NoteBookHandler struct {
	Repo *repository.GormNoteBookRepostiry
}

func NewNoteBookHandler(repo *repository.GormNoteBookRepostiry) *NoteBookHandler {
	return &NoteBookHandler{Repo: repo}
}

// create new notebook
func (h *NoteBookHandler) CreateNotebook(c echo.Context) error {
	notebook := new(dto.NoteBook)
	if err := c.Bind(&notebook); err != nil {
		return c.JSON(400, err.Error())
	}

	err := helpers.ValidateNotebook(notebook)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.Repo.CreateNoteBook(notebook)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to create notebook")
	}
	return c.JSON(http.StatusOK, notebook)
}

// get all notebooks
func (h *NoteBookHandler) GetAllNotebooks(c echo.Context) error {

	search := c.QueryParam("search")

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

	notebooks, total, err := h.Repo.GetAllNoteBooks(page, limit, search)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get notebooks")
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
		Data:       notebooks,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: int(total / int64(limit)),
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}
	return c.JSON(http.StatusOK, responses)
}

// get single notebook
func (h *NoteBookHandler) GetNotebook(c echo.Context) error {
	id, err := helpers.GetParam(c)

	if err != nil {
		log.Error("Invalid Notebook Id")
		return c.JSON(http.StatusBadRequest, "Invalid Notebook Id")
	}

	notebook, err := h.Repo.GetNoteBookByID(id)
	if err != nil {
		log.Error("Notebook Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "Notebook Not Found")
	}

	return c.JSON(http.StatusOK, notebook)
}

// update notebook
func (h *NoteBookHandler) UpdateNotebook(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid NoteBook Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Notebook Id")
	}

	var notebookDto dto.NoteBook
	if err := c.Bind(&notebookDto); err != nil {
		log.Error("Invalid Request data", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Request data")
	}

	err = helpers.ValidateNotebook(&notebookDto)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	notebook := new(models.NoteBook)
	notebook, err = h.Repo.UpdateNoteBookByID(id, &notebookDto)
	if err != nil {
		log.Error("Failed to update notebook", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update notebook")
	}

	return c.JSON(http.StatusOK, notebook)
}

// delete notebook
func (h *NoteBookHandler) DeleteNotebook(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid NoteBook Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Notebook Id")
	}

	err = h.Repo.DeleteNoteBookByID(id)
	if err != nil {
		log.Error("Notebook not found", err.Error())
		return c.JSON(http.StatusInternalServerError, "Notebook not found")
	}

	return c.NoContent(http.StatusNoContent)
}
