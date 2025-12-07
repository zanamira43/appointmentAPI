package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"github.com/zanamira43/appointment-api/response"
	"gorm.io/gorm"
)

type GormPatientRepository struct {
	DB *gorm.DB
}

func NewGormPatientRepository(db *gorm.DB) *GormPatientRepository {
	return &GormPatientRepository{DB: db}
}

// insert new patient data into sql database
func (r *GormPatientRepository) CreatePatient(patientDto *dto.Patient, UserID uint) error {

	request := models.Patient{
		Name:          patientDto.Name,
		Gender:        patientDto.Gender,
		Age:           patientDto.Age,
		MarriedStatus: patientDto.MarriedStatus,
		Profession:    patientDto.Profession,
		Address:       patientDto.Address,
		PhoneNumber:   patientDto.PhoneNumber,
		UserID:        UserID,
	}

	return r.DB.Create(&request).Error
}

// get all patient data from sql database
func (r *GormPatientRepository) GetAllPatients(page, limit int, search string, userID uint, userRole string) ([]models.Patient, int64, error) {
	var patients []models.Patient
	var total int64

	// create blank query to build upon
	query := r.DB.Model(&models.Patient{})

	// If the user is not an admin, they should see all patients created by non-admins.
	// Admins can see all patients.
	if userRole != "admin" {
		query = query.Joins("JOIN users ON users.id = patients.user_id AND users.role != ?", "admin")
	}

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("patients.name LIKE ? OR patients.slug LIKE ? OR patients.phone_number LIKE ?", searchPattern, searchPattern, searchPattern)
	}

	// get total number of patients
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	err = query.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, first_name")
	}).Order("id desc").Find(&patients).Error
	if err != nil {
		return nil, 0, err
	}

	return patients, total, nil

}

// get patient data by id from sql database
func (r *GormPatientRepository) GetPatientByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	err := r.DB.Preload("Problem").Preload("Sessions").Preload("Payments").Where("id = ?", id).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

// get patient data by slug from sql database
func (r *GormPatientRepository) GetPatientBySlug(slug string) (*models.Patient, error) {
	var patient models.Patient
	err := r.DB.Preload("Problem").Preload("Sessions").Preload("Payments").Where("slug = ?", slug).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

// update patient data by id from sql database
func (r *GormPatientRepository) UpdatePatient(id uint, dtoPatient *dto.Patient, UserID uint) (*models.Patient, error) {
	var patient models.Patient
	err := r.DB.Where("id = ?", id).First(&patient).Error
	if err != nil {
		return nil, err
	}
	if dtoPatient.Name != "" {
		patient.Name = dtoPatient.Name
	}
	if dtoPatient.Gender != "" {
		patient.Gender = dtoPatient.Gender
	}
	if dtoPatient.Age != 0 {
		patient.Age = dtoPatient.Age
	}

	if dtoPatient.MarriedStatus != "" {
		patient.MarriedStatus = dtoPatient.MarriedStatus
	}

	if dtoPatient.Profession != "" {
		patient.Profession = dtoPatient.Profession
	}
	if dtoPatient.Address != "" {
		patient.Address = dtoPatient.Address
	}
	if dtoPatient.PhoneNumber != "" {
		patient.PhoneNumber = dtoPatient.PhoneNumber
	}

	if UserID != 0 {
		patient.UserID = UserID
	}

	err = r.DB.Save(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

// delete patient data by id from sql database
func (r *GormPatientRepository) DeletePatient(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&dto.Patient{}).Error
}

// patinet outcome calculation
func (r *GormPatientRepository) PatinetOutcome(id uint) (*response.PatientOutcomeResponse, error) {
	var patient models.Patient

	var SumReceivedSessionCount int64

	err := r.DB.
		Select("id", "name").
		Preload("Problem", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "patient_id", "need_sessions_count", "is_dollar_payment", "session_price", "session_total_price")
		}).
		Preload("Sessions", func(db *gorm.DB) *gorm.DB {

			return db.Select("id", "patient_id", "status").Where("status LIKE ?", "completed").
				Count(&SumReceivedSessionCount)
		}).
		Preload("Payments", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "patient_id", "amount")
		}).
		First(&patient, id).Error

	if err != nil {
		return nil, err
	}

	// Calculate totals
	var totalReceivedPayments float32
	for _, payment := range patient.Payments {
		totalReceivedPayments += float32(payment.Amount)
	}

	response := &response.PatientOutcomeResponse{
		ID:   patient.ID,
		Name: patient.Name,

		NeedSessionsCount: patient.Problem.NeedSessionsCount,
		IsDollarPaymnet:   patient.Problem.IsDollarPayment,
		SessionPrice:      patient.Problem.SessionPrice,
		SessionTotalPrice: patient.Problem.SessionTotalPrice,

		SumReceivedSessionCount:   SumReceivedSessionCount,
		SessionReceivedTotalPrice: float32(SumReceivedSessionCount) * patient.Problem.SessionPrice,
		TotalReceivedPayments:     totalReceivedPayments,
		RemainingBalance:          totalReceivedPayments - (float32(SumReceivedSessionCount) * patient.Problem.SessionPrice),
	}

	return response, nil
}
