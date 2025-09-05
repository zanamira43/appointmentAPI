package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/zanamira43/appointment-api/database"
	"github.com/zanamira43/appointment-api/handlers"
	"github.com/zanamira43/appointment-api/repository"
)

type Middleware struct{}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

// authenticated middleware
func (m *Middleware) IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userRepo := repository.NewGormUserRepository(database.DB)
		authHandler := handlers.NewAuth(userRepo)
		_, err := authHandler.GetUserByCookie(c)
		if err != nil {
			log.Error("Sorry! User UnAuthenticated", err.Error())
			return c.JSON(http.StatusBadRequest, "Sorry! User UnAuthenticated")
		}

		return next(c)
	}
}

func (m *Middleware) IsUserAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userRepo := repository.NewGormUserRepository(database.DB)
		authHandler := handlers.NewAuth(userRepo)
		user, _ := authHandler.GetUserByCookie(c)
		if user.Role != "admin" {
			log.Error("Sorry! User Is Not Admin")
			return c.JSON(http.StatusForbidden, "Sorry! User Is Not Admin")
		}

		return next(c)
	}
}
