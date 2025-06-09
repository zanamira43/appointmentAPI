package models

import "time"

type Payment struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	PatientID     uint   `json:"patient_id"`
	ServiceTypeID uint   `json:"service_type_id"`
	OfferID       uint   `json:"offer_id"`
	Amount        int    `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	Status        string `json:"status"`

	Patient     Patient     `json:"patient" gorm:"foreignKey:PatientID;constraint:OnDelete:SET NULL;"`
	ServiceType ServiceType `json:"service_type" gorm:"foreignKey:ServiceTypeID;constraint:OnDelete:SET NULL;"`
	Offer       Offer       `json:"offer" gorm:"foreignKey:OfferID;constraint:OnDelete:SET NULL;"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}
