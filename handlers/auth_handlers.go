package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/helpers"
	"github.com/zanamira43/appointment-api/models"
	"github.com/zanamira43/appointment-api/repository"
	"github.com/zanamira43/appointment-api/utils"
)

type Auth struct {
	Repo *repository.GormUserRepository
}

func NewAuth(repo *repository.GormUserRepository) *Auth {
	return &Auth{Repo: repo}
}

// register new user
func (h *Auth) SignUP(c echo.Context) error {

	// var data map[string]string
	// create user register instance
	var data dto.Register

	// parse data from fornt end
	if err := c.Bind(&data); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := helpers.ValidateRegisterUser(&data); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}
	// matching password
	if data.Password != data.PasswordConfirm {
		log.Error("password does not match")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "password does not match",
		})
	}

	// sending user info into database
	user := &dto.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Phone:     data.Phone,
	}

	// hashing password
	user.SetPassword(data.Password)
	err := h.Repo.CreateUser(user)
	if err != nil {
		log.Error("Failed to register new user", err.Error())
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to register new user",
		})
	}

	// return user after registered in database
	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

// login exists user
func (h *Auth) Login(c echo.Context) error {
	// var data map[string]string
	data := new(dto.Login)

	if err := c.Bind(&data); err != nil {
		return err
	}

	// checking user email
	user := new(models.User)
	user, err := h.Repo.GetUserByPhone(data.Phone)
	if err != nil {
		log.Error("Incorrect phone number")
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Incorrect phone number",
		})
	}

	// checking hashing password
	err = user.ComparedPassword(data.Password)
	if err != nil {
		log.Error("incorrect password", err.Error())
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "incorrect password",
		})
	}

	// create token by user id
	token, err := utils.GenerateJwt(user.ID)
	if err != nil {
		log.Error("there is a problem in token creating process", err.Error())
		return c.JSON(http.StatusInternalServerError, "there is a problem in token creating process")
	}

	// set token to cookies
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/api",
		HttpOnly: true,
		Secure:   true, // Set to true if using HTTPS
		Expires:  time.Now().Add(time.Hour * 24),
		SameSite: http.SameSiteNoneMode,
		// SameSite: http.SameSiteLaxMode,
	}
	c.SetCookie(&cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"user":  user,
		"token": token,
	})
}

// / get id by cookie
func (h *Auth) GetUserByCookie(c echo.Context) (*models.User, error) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return nil, err
	}

	userId, err := utils.ParseJwt(cookie.Value)
	if err != nil {
		return nil, err
	}

	id, _ := strconv.Atoi(userId)
	user, err := h.Repo.GetUserByID(uint(id))
	if err != nil {

		return nil, err
	}

	return user, nil
}

// authenticated user
func (h *Auth) User(c echo.Context) error {
	user, err := h.GetUserByCookie(c)
	if err != nil {
		log.Error("Sorry! User UnAuthenticated", err.Error())
		return c.JSON(http.StatusNotFound, "Sorry! User UnAuthenticated")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

// update user info
func (h *Auth) UpdateInfo(c echo.Context) error {
	var data dto.User

	if err := c.Bind(&data); err != nil {
		return err
	}

	// get user by cookie
	user, err := h.GetUserByCookie(c)
	if err != nil {
		log.Error("Sorry! User UnAuthenticated", err.Error())
		return c.JSON(http.StatusNotFound, "Sorry! User UnAuthenticated")
	}

	// update user info
	updatedUser, err := h.Repo.UpdateUser(user.ID, &data)
	if err != nil {
		log.Error("Failed to update user", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update user")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": updatedUser,
	})
}

// update user password
func (h *Auth) UpdatePassword(c echo.Context) error {
	var data dto.Register

	if err := c.Bind(&data); err != nil {
		return err
	}

	// comparing the password
	if data.Password != data.PasswordConfirm {
		log.Error("password does not match")
		return c.JSON(http.StatusBadRequest, "password does not match")
	}

	// hashing new password
	updatepassword := new(dto.User)
	updatepassword.SetPassword(data.Password)

	// get user by cookie
	user, err := h.GetUserByCookie(c)
	if err != nil {
		log.Error("Sorry! User UnAuthenticated", err.Error())
		return c.JSON(http.StatusNotFound, "Sorry! User UnAuthenticated")
	}

	// update user password
	user, err = h.Repo.UpdateUser(user.ID, updatepassword)
	if err != nil {
		log.Error("Failed to update password", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update password")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "password updated successfully",
		"user":    user,
	})
}

// logout method
func (h *Auth) Logout(c echo.Context) error {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/api",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true, // Set to true if using HTTPS
		// SameSite: http.SameSiteLaxMode, // Use Lax mode
		SameSite: http.SameSiteNoneMode, // Use None mode for cross-site
		MaxAge:   -1,
	}

	c.SetCookie(&cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "user logged out successfully",
	})
}
