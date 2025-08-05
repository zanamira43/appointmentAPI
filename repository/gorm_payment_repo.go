package repository

import (
	"time"

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
func (r *GormPaymentRepository) CreatepPayments(payment *dto.PaymentDto) error {
	return r.DB.Create(&payment).Error
}

// get all payment data from sql database
func (r *GormPaymentRepository) GetPayments() ([]models.Payment, error) {
	var payments []models.Payment
	err := r.DB.Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
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
func (r *GormPaymentRepository) UpdatePayment(id uint, dtopayment *dto.PaymentDto) (*models.Payment, error) {
	var payment models.Payment
	err := r.DB.Where("id = ?", id).First(&payment).Error
	if err != nil {
		return nil, err
	}

	if dtopayment.PatientID != 0 {
		payment.PatientID = dtopayment.PatientID
	}
	if dtopayment.SessionID != 0 {
		payment.SessionID = dtopayment.SessionID
	}

	if dtopayment.Amount != 0 {
		payment.Amount = dtopayment.Amount
	}
	if dtopayment.PaymentMethod != "" {
		payment.PaymentMethod = dtopayment.PaymentMethod
	}
	if dtopayment.PaymentDate != (time.Time{}) {
		payment.PaymentDate = dtopayment.PaymentDate
	}
	if dtopayment.Status != "" {
		payment.Status = dtopayment.Status
	}

	err = r.DB.Save(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

// delete payment data by id from sql database
func (r *GormPaymentRepository) DeletePayment(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&dto.PaymentDto{}).Error
}
