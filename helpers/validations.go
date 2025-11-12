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

func ValidateUser(dto *dto.UserRequest) error {
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

	if dto.Role == "" {
		return errors.New("role	 is required")
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

	if dto.MarriedStatus == " " {
		return errors.New("married status is required")
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
func ValidateSession(dto *dto.Session) error {
	if dto.PatientID == 0 {
		return errors.New("patient is required")
	}

	if dto.Subject == "" {
		return errors.New("subject is required")
	}

	if dto.CommunicationTypes == "" {
		return errors.New("communication types is required")
	}

	if dto.SessionDate == "" {
		return errors.New("session date is required")
	}

	if dto.Status == "" {
		return errors.New("status is required")
	}

	return nil
}

func ValidateTimeTables(dto *dto.TimeTable) error {

	if dto.PatientName == "" {
		return errors.New("patient name is required")
	}
	if len(dto.WeekDay) == 0 {
		return errors.New("week day is required")
	}

	if dto.StartTime == "" {
		return errors.New("start time is required")
	}
	if dto.EndTime == "" {
		return errors.New("end time is required")
	}

	return nil
}

func ValidateProblems(dto *dto.Problem) error {
	if dto.PatientID == 0 {
		return errors.New("patient is required")
	}

	if len(dto.MianpProblems) == 0 {
		return errors.New("main problems is required")
	}

	if len(dto.SecondaryProblems) == 0 {
		return errors.New("secondary problems is required")
	}

	if dto.NeedSessionsCount == 0 {
		return errors.New("need sessions count is required")
	}

	if dto.SessionPrice == 0 {
		return errors.New("session price is required")
	}
	return nil
}

// validate payment type
func ValidatePaymentType(dto *dto.PaymentType) error {
	if dto.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

// validate service type
func ValidatePayment(dto *dto.Payment) error {

	if dto.PatientID == 0 {
		return errors.New("patient is required")
	}
	if dto.PaymentTypeID == 0 {
		return errors.New("payment type is required")
	}
	if dto.Amount == 0 {
		return errors.New("amount is required")
	}
	if dto.PaymentDate == "" {
		return errors.New("payment date is required")
	}

	return nil
}

func ValidateNotebook(dto *dto.NoteBook) error {
	if dto.Content == "" {
		return errors.New("content is required")
	}

	return nil
}
