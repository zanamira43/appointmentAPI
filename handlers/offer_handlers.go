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

type SessionHandler struct {
	SessionRepository *repository.GormSessionRepository
}

func NewSessionHandler(Repo *repository.GormSessionRepository) *SessionHandler {
	return &SessionHandler{SessionRepository: Repo}
}

// Create New session
func (h *SessionHandler) CreateSessions(c echo.Context) error {
	sessiondto := new(dto.SessionDto)

	if err := c.Bind(&sessiondto); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate the request body
	if err := helpers.ValidateSession(sessiondto); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.SessionRepository.CreateSession(sessiondto)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to create session")
	}
	return c.JSON(http.StatusOK, sessiondto)
}

// Get all sessions
func (h *SessionHandler) GetSessions(c echo.Context) error {
	sessions, err := h.SessionRepository.GetAllSessions()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get sessions")
	}
	return c.JSON(http.StatusOK, sessions)
}

// Get single session
func (h *SessionHandler) GetSession(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Offer Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Offer Id")
	}

	session, err := h.SessionRepository.GetSessionByID(id)
	if err != nil {
		log.Error("Offer Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "Offer Not Found")
	}
	return c.JSON(http.StatusOK, session)
}

// update single session by id
func (h *SessionHandler) UpdateSession(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Offer Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Offer Id")
	}

	var sessionDto dto.SessionDto
	if err := c.Bind(&sessionDto); err != nil {
		log.Error("Invalid Request data", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Request data")
	}

	// Validate the request body
	if err := helpers.ValidateSession(&sessionDto); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	session := new(models.Session)
	session, err = h.SessionRepository.UpdateSession(id, &sessionDto)
	if err != nil {
		log.Error("Failed to update session", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update session")
	}
	return c.JSON(http.StatusOK, session)
}

// delete offer by id
func (h *SessionHandler) DeleteSession(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Session Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Session Id")
	}

	err = h.SessionRepository.DeleteSession(id)
	if err != nil {
		log.Error("Session not found", err.Error())
		return c.JSON(http.StatusInternalServerError, "Session not found")
	}
	return c.NoContent(http.StatusNoContent)
}
