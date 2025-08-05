package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/gorm"
)

type GormPatientRepository struct {
	DB *gorm.DB
}

func NewGormPatientRepository(db *gorm.DB) *GormPatientRepository {
	return &GormPatientRepository{DB: db}
}

// insert new patient data into sql database
func (r *GormPatientRepository) CreatePatient(patient *dto.PatientDto) error {
	return r.DB.Create(&patient).Error
}

// get all patient data from sql database
func (r *GormPatientRepository) GetAllPatients() ([]models.Patient, error) {
	var patients []models.Patient
	err := r.DB.Find(&patients).Error
	if err != nil {
		return nil, err
	}
	return patients, nil
}

// get patient data by id from sql database
func (r *GormPatientRepository) GetPatientByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	err := r.DB.Where("id = ?", id).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

// get patient data by slug from sql database
func (r *GormPatientRepository) GetPatientBySlug(slug string) (*models.Patient, error) {
	var patient models.Patient
	err := r.DB.Where("slug = ?", slug).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

// update patient data by id from sql database
func (r *GormPatientRepository) UpdatePatient(id uint, dtoPatient *dto.PatientDto) (*models.Patient, error) {
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
	if dtoPatient.Profession != "" {
		patient.Profession = dtoPatient.Profession
	}
	if dtoPatient.Address != "" {
		patient.Address = dtoPatient.Address
	}
	if dtoPatient.PhoneNumber != "" {
		patient.PhoneNumber = dtoPatient.PhoneNumber
	}

	err = r.DB.Save(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

// delete patient data by id from sql database
func (r *GormPatientRepository) DeletePatient(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&dto.PatientDto{}).Error
}
