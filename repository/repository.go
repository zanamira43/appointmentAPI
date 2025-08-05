package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
)

// user repository interface
type UserRepository interface {
	CreateUser(user *dto.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(id string, user *dto.User) (*models.User, error)
	DeleteUser(id string) error
}

// project repository interface
type PatientRepository interface {
	CreatePatient(patient *dto.PatientDto) error
	GetAllPatients() ([]models.Patient, error)
	GetPatientByID(id uint) (*models.Patient, error)
	GetPatientBySlug(slug string) (*models.Patient, error)
	UpdatePatient(id uint, patient *dto.PatientDto) (*models.Patient, error)
	DeletePatient(id uint) error
}

type SessionRepository interface {
	CreateSession(session *dto.SessionDto) error
	GetAllSessions() ([]models.Session, error)
	GetSessionByID(id uint) (*models.Session, error)
	UpdateSession(id uint, session *dto.SessionDto) (*models.Session, error)
	DeleteSession(id uint) error
}

// payment repository interface
type PyamentRepository interface {
	CreatePayments(payment *dto.PaymentDto) error
	GetAllPayments() ([]models.Payment, error)
	GetPaymentsByID(id uint) (*models.Payment, error)
	UpdatePaymentsByID(id uint, payment *dto.PaymentDto) (*models.Payment, error)
	DeletePaymentsByID(id uint) error
}
