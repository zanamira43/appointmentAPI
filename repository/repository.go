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
	GetUserByPhone(phone_number string) (*models.User, error)
	UpdateUser(id string, user *dto.User) (*models.User, error)
	UpdateUserPassword(id string, user *dto.User) (*models.User, error)
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

type TimeTableRepository interface {
	CreateTimeTables(dtt *dto.TimeTable) error
	GetAllTimeTables() ([]models.TimeTable, error)
	GetTimeTableByID(id uint) (*models.TimeTable, error)
	UpdateTimeTableByID(id uint, dtt *dto.TimeTable) (*models.TimeTable, error)
	DeleteTimeTableByID(id uint) error
}
type ProblemRepository interface {
	CreateProblems(problem *dto.Problem) error
	GetAllProblems() ([]models.Problem, error)
	GetProblemByID(id uint) (*models.Problem, error)
	GetProblemByPatientId(patientId uint) (*models.Problem, error)
	UpdateProblemByID(id uint, problem *dto.Problem) (*models.Problem, error)
	DeleteProblemByID(id uint) error
}

type SessionRepository interface {
	CreateSession(session *dto.Session) error
	GetAllSessions() ([]models.Session, error)
	GetSessionsByPatientID(patientId uint) ([]models.Session, error)
	GetSessionByID(id uint) (*models.Session, error)
	UpdateSession(id uint, session *dto.Session) (*models.Session, error)
	DeleteSession(id uint) error
}

// payment repository interface
type PyamentRepository interface {
	CreatePayments(payment *dto.Payment) error
	GetAllPayments() ([]models.Payment, error)
	GetPaymentsByID(id uint) (*models.Payment, error)
	UpdatePaymentsByID(id uint, payment *dto.Payment) (*models.Payment, error)
	DeletePaymentsByID(id uint) error
}
