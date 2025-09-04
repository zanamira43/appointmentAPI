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
func (r *GormPatientRepository) CreatePatient(patient *dto.Patient) error {
	return r.DB.Create(&patient).Error
}

// get all patient data from sql database
func (r *GormPatientRepository) GetAllPatients(page, limit int, search string) ([]models.Patient, int64, error) {
	var patients []models.Patient

	var total int64

	offset := (page - 1) * limit

	// get total number of patients
	err := r.DB.Model(&patients).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.DB.Offset(offset).Limit(limit).Where("name LIKE ?", "%"+search+"%").Or("slug LIKE ?", "%"+search+"%").Or("phone_number LIKE ?", "%"+search+"%").Find(&patients).Error
	if err != nil {
		return nil, 0, err
	}

	return patients, total, nil

}

// get patient data by id from sql database
func (r *GormPatientRepository) GetPatientByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	err := r.DB.Preload("Sessions").Preload("Payments").Where("id = ?", id).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

// get patient data by slug from sql database
func (r *GormPatientRepository) GetPatientBySlug(slug string) (*models.Patient, error) {
	var patient models.Patient
	err := r.DB.Preload("Sessions").Preload("Payments").Where("slug = ?", slug).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

// update patient data by id from sql database
func (r *GormPatientRepository) UpdatePatient(id uint, dtoPatient *dto.Patient) (*models.Patient, error) {
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
