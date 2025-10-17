package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/gorm"
)

type GormPaymentRepository struct {
	DB *gorm.DB
}

func NewGormPaymentRepository(db *gorm.DB) *GormPaymentRepository {
	return &GormPaymentRepository{DB: db}
}

// insert new  payment type data into sql database
func (r *GormPaymentRepository) CreatepPayments(payment *dto.Payment) error {
	return r.DB.Create(&payment).Error
}

// get all payment data from sql database
func (r *GormPaymentRepository) GetPayments(page, limit int, search string) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64

	query := r.DB.Model(&models.Payment{})

	if search != "" {
		searchPattern := "%" + search + "%"
		query.Where("payment_date LIKE ? ", searchPattern)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Preload("PaymentType").Offset(offset).Limit(limit)

	}
	err = query.Find(&payments).Error
	if err != nil {
		return nil, 0, err
	}

	return payments, total, nil
}

// get all payments data from sql database by patient id
func (r *GormPaymentRepository) GetPaymentsByPatientID(page, limit int, search string, patientId uint) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64

	// create query with patient_id filter
	query := r.DB.Where("patient_id = ?", patientId).Model(&models.Payment{})

	// search by payment date
	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("payment_date LIKE ? ", searchPattern)
	}

	// get total number of payments for this patient
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	err = query.Preload("PaymentType").Find(&payments).Error
	if err != nil {
		return nil, 0, err
	}
	return payments, total, nil
}

// get payment data by id from sql database
func (r *GormPaymentRepository) GetPayment(id uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.DB.Where("id = ?", id).First(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

// update payment data by id from sql database
func (r *GormPaymentRepository) UpdatePayment(id uint, dtopayment *dto.Payment) (*models.Payment, error) {
	var payment models.Payment
	err := r.DB.Where("id = ?", id).First(&payment).Error
	if err != nil {
		return nil, err
	}

	if dtopayment.PatientID != 0 {
		payment.PatientID = dtopayment.PatientID
	}

	if dtopayment.PaymentTypeID != 0 {
		payment.PaymentTypeID = dtopayment.PaymentTypeID
	}

	if dtopayment.Amount != 0 {
		payment.Amount = dtopayment.Amount
	}
	if dtopayment.PaymentDate != "" {
		payment.PaymentDate = dtopayment.PaymentDate
	}

	err = r.DB.Save(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

// delete payment data by id from sql database
func (r *GormPaymentRepository) DeletePayment(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&dto.Payment{}).Error
}
