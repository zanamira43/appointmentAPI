package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/helpers"
	"github.com/zanamira43/appointment-api/repository"
)

type UserHandler struct {
	UserRepository *repository.GormUserRepository
}

func NewUserHandler(userRepo *repository.GormUserRepository) *UserHandler {
	return &UserHandler{UserRepository: userRepo}
}

// create new user
func (h *UserHandler) CreateUser(c echo.Context) error {
	newUser := new(dto.UserRequest)
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := helpers.ValidateUser(newUser); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user := &dto.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Phone:     newUser.Phone,
		Role:      newUser.Role,
		Active:    newUser.Active,
	}

	user.SetPassword(newUser.Password)

	err := h.UserRepository.CreateUser(user)
	if err != nil {
		log.Error("Failed to register new user", err.Error())
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to register new user",
		})
	}
	return c.JSON(http.StatusOK, user)
}

// get all users
func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.UserRepository.GetAllUsers()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get users")
	}
	return c.JSON(http.StatusOK, users)
}

// get user by id
func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid User Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid User Id")
	}
	user, err := h.UserRepository.GetUserByID(id)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "User Not Found")
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid User Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid User Id")
	}

	var data dto.User
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := h.UserRepository.UpdateUser(id, &data)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "User Not Found")
	}

	return c.JSON(http.StatusOK, user)
}

// delete user
func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid User Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid User Id")
	}
	err = h.UserRepository.DeleteUser(id)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "User Not Found")
	}
	return c.NoContent(http.StatusNoContent)
}
