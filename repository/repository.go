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

// offer repository interface
type OfferRepository interface {
	CreateOffers(offer *dto.Offer) error
	GetAllOffers() ([]models.Offer, error)
	GetOfferByID(id uint) (*models.Offer, error)
	UpdateOfferByID(id uint, offer *dto.Offer) (*models.Offer, error)
	DeleteOfferByID(id uint) error
}

// service type repository interface
type ServiceTypeRepository interface {
	CreateServiceTypes(offer *dto.ServiceType) error
	GetAllServiceTypes() ([]models.ServiceType, error)
	GetServiceTypeByID(id uint) (*models.ServiceType, error)
	UpdateServiceTypeByID(id uint, offer *dto.ServiceType) (*models.ServiceType, error)
	DeleteServiceTypeByID(id uint) error
}
