package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"github.com/zanamira43/appointment-api/response"
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
	GetAllPatients(page, limit int, search string) ([]models.Patient, error)
	GetPatientByID(id uint) (*models.Patient, error)
	GetPatientBySlug(slug string) (*models.Patient, error)
	UpdatePatient(id uint, patient *dto.Patient) (*models.Patient, error)
	DeletePatient(id uint) error
	PatinetOutcome(id uint) (*response.PatientOutcomeResponse, error)
}

type TimeTableRepository interface {
	CreateTimeTables(dtt *dto.TimeTable) error
	GetAllTimeTables(page, limit int, search string) ([]models.TimeTable, error)
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
	GetAllSessions(page, limit int, search string) ([]models.Session, error)
	GetSessionsByPatientID(patientId uint, page, limit int, search string) ([]models.Session, error)
	GetSessionByID(id uint) (*models.Session, error)
	UpdateSession(id uint, session *dto.Session) (*models.Session, error)
	DeleteSession(id uint) error
}

// payment repository interface
type PaymentTypeRepository interface {
	CreatePaymentType(paymentType *dto.PaymentType) error
	GetPaymentTypes(page, limit int, search string) ([]models.PaymentType, int64, error)
	GetPaymentType(id uint) (*models.PaymentType, error)
	UpdatePaymentType(id uint, dtoPaymentType *dto.PaymentType) (*models.PaymentType, error)
	DeletePaymentType(id uint) error
}

// payment repository interface
type PyamentRepository interface {
	CreatePayments(payment *dto.Payment) error
	GetPayments(page, limit int, search string) ([]models.Payment, int64, error)
	GetPaymentsByPatientID(page, limit int, search string, patientId uint) ([]models.Payment, int64, error)
	GetPaymentsByID(id uint) (*models.Payment, error)
	UpdatePaymentsByID(id uint, payment *dto.Payment) (*models.Payment, error)
	DeletePaymentsByID(id uint) error
}

// setting repository
type SettingRepository interface {
	GetSetting() (*models.Settings, error)
	UpdateSetting(settings *dto.Settings) error
}

type NoteBookRepository interface {
	CreateNoteBook(noteBook *dto.NoteBook) error
	GetAllNoteBooks(page, limit int, search string) ([]models.NoteBook, int64, error)
	GetNoteBookByID(id uint) (*models.NoteBook, error)
	UpdateNoteBookByID(id uint, noteBook *dto.NoteBook) (*models.NoteBook, error)
	DeleteNoteBookByID(id uint) error
}
