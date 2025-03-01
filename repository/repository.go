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
	CreatePatient(patient *dto.Patient) error
	GetAllPatients() ([]models.Patient, error)
	GetPatientByID(id uint) (*models.Patient, error)
	GetPatientBySlug(slug string) (*models.Patient, error)
	UpdatePatient(id uint, patient *dto.Patient) (*models.Patient, error)
	DeletePatient(id uint) error
}
