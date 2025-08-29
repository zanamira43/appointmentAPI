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

	if dto.Phone == "" {
		return errors.New("phone is required")
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
	if dto.PhoneNumber == "" {
		return errors.New("phone number is required")
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

	return nil
}

// validate offer
func ValidateSession(dto *dto.SessionDto) error {
	if dto.PatientID == 0 {
		return errors.New("patient is required")
	}

	if dto.SessionDate == "" {
		return errors.New("session date is required")
	}

	if dto.Status == "" {
		return errors.New("status is required")
	}

	return nil
}

// validate service type
func ValidatePayment(dto *dto.PaymentDto) error {
	if dto.PatientID == 0 {
		return errors.New("patient is required")
	}
	if dto.SessionID == 0 {
		return errors.New("session is required")
	}
	if dto.Amount == 0 {
		return errors.New("amount is required")
	}
	if dto.PaymentDate == "" {
		return errors.New("payment date is required")
	}
	if dto.Status == "" {
		return errors.New("status is required")
	}

	return nil
}
