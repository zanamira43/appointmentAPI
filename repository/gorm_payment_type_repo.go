package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/gorm"
)

type GormPaymentTypeRepository struct {
	DB *gorm.DB
}

func NewGormPaymentTypeRepository(db *gorm.DB) *GormPaymentTypeRepository {
	return &GormPaymentTypeRepository{DB: db}
}

// insert new payment type data into sql database
func (r *GormPaymentTypeRepository) CreatePaymentType(paymentType *dto.PaymentType) error {
	return r.DB.Create(&paymentType).Error
}

// get all payment types data from sql database
func (r *GormPaymentTypeRepository) GetPaymentTypes(page, limit int, search string) ([]models.PaymentType, int64, error) {
	var paymentTypes []models.PaymentType
	var total int64

	query := r.DB.Model(&models.PaymentType{})

	if search != "" {
		searchPattern := "%" + search + "%"
		query.Where("name LIKE ? ", searchPattern)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	err = query.Find(&paymentTypes).Error
	if err != nil {
		return nil, 0, err
	}
	return paymentTypes, total, nil
}

// get payment type data by id from sql database
func (r *GormPaymentTypeRepository) GetPaymentType(id uint) (*models.PaymentType, error) {
	var paymentType models.PaymentType
	err := r.DB.Where("id = ?", id).First(&paymentType).Error
	if err != nil {
		return nil, err
	}
	return &paymentType, nil
}

// update payment type data by id from sql database
func (r *GormPaymentTypeRepository) UpdatePaymentType(id uint, dtoPaymentType *dto.PaymentType) (*models.PaymentType, error) {
	var paymentType models.PaymentType
	err := r.DB.Where("id = ?", id).First(&paymentType).Error
	if err != nil {
		return nil, err
	}

	if dtoPaymentType.Name != "" {
		paymentType.Name = dtoPaymentType.Name
	}

	err = r.DB.Save(&paymentType).Error
	if err != nil {
		return nil, err
	}
	return &paymentType, nil
}

// delete payment type data by id from sql database
func (r *GormPaymentTypeRepository) DeletePaymentType(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&models.PaymentType{}).Error
}
