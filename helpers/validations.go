package helpers

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zanamira43/appointment-api/dto"
)

// get id parameter form echo context
func GetParam(c echo.Context) (uint, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	return uint(id), nil
}

// validate user

func ValidateRegisterUser(dto *dto.Register) error {
	if dto.FirstName == "" {
		return errors.New("first name is required")
	}
	if dto.LastName == "" {
		return errors.New("last name  is required")
	}

	if dto.Email == "" {
		return errors.New("email is required")
	}

	if dto.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

// validate patient
func ValidatePatient(dto *dto.Patient) error {
	if dto.Name == "" {
		return errors.New("name is required")
	}

	if dto.Gender == "" {
		return errors.New("gender is required")
	}

	if dto.Age == 0 {
		return errors.New("age is required")
	}

	if dto.Profession == "" {
		return errors.New("profession is required")
	}

	if dto.Address == "" {
		return errors.New("address is required")
	}

	if dto.PhoneNumber == "" {
		return errors.New("address is required")
	}

	return nil
}

// validate offer
func ValidateOffer(dto *dto.Offer) error {
	if dto.Title == "" {
		return errors.New("title is required")
	}

	if dto.ServiceTypeID == 0 {
		return errors.New("service type is required")
	}

	if dto.Price == 0 {
		return errors.New("price is required")
	}

	return nil
}
