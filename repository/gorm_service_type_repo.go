package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/gorm"
)

type GormServiceTypeRepository struct {
	DB *gorm.DB
}

func NewGormServiceTypeRepository(db *gorm.DB) *GormServiceTypeRepository {
	return &GormServiceTypeRepository{DB: db}
}

// insert new service type data into sql database
func (r *GormServiceTypeRepository) CreateServiceTypes(serviceType *dto.ServiceType) error {
	return r.DB.Create(&serviceType).Error
}

// get all service data from sql database
func (r *GormServiceTypeRepository) GetAllServiceTypes() ([]models.ServiceType, error) {
	var serviceTypes []models.ServiceType
	err := r.DB.Find(&serviceTypes).Error
	if err != nil {
		return nil, err
	}
	return serviceTypes, nil
}

// get service type data by id from sql database
func (r *GormServiceTypeRepository) GetServiceTypeByID(id uint) (*models.ServiceType, error) {
	var serviceType models.ServiceType
	err := r.DB.Where("id = ?", id).First(&serviceType).Error
	if err != nil {
		return nil, err
	}
	return &serviceType, nil
}

// update service type data by id from sql database
func (r *GormServiceTypeRepository) UpdateServiceTypeByID(id uint, dtoServiceType *dto.ServiceType) (*models.ServiceType, error) {
	var serviceType models.ServiceType
	err := r.DB.Where("id = ?", id).First(&serviceType).Error
	if err != nil {
		return nil, err
	}
	if dtoServiceType.Name != "" {
		serviceType.Name = dtoServiceType.Name
	}

	err = r.DB.Save(&serviceType).Error
	if err != nil {
		return nil, err
	}
	return &serviceType, nil
}

// delete service type data by id from sql database
func (r *GormServiceTypeRepository) DeleteServiceTypeByID(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&dto.ServiceType{}).Error
}
