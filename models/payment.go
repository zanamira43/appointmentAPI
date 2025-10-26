package models

import "time"

type PaymentType struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Payment struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	PatientID       uint   `json:"patient_id"`
	PaymentTypeID   uint   `json:"payment_type_id"`
	IsDollarPayment bool   `json:"is_dollar_payment"`
	Amount          int    `json:"amount"`
	PaymentDate     string `json:"payment_date"`

	Patient     Patient     `json:"patient" gorm:"foreignKey:PatientID;constraint:OnDelete:SET NULL;"`
	PaymentType PaymentType `json:"payment_type" gorm:"foreignKey:PaymentTypeID;constraint:OnDelete:SET NULL;"`
	CreatedAt   time.Time   `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}
