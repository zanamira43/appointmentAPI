package models

import "time"

type Payment struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	PatientID     uint      `json:"patient_id"`
	SessionID     uint      `json:"session_id"`
	Amount        int       `json:"amount"`
	PaymentDate   time.Time `json:"payment_date"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status"`

	Patient   Patient   `json:"patient" gorm:"foreignKey:PatientID;constraint:OnDelete:SET NULL;"`
	Session   Session   `json:"session" gorm:"foreignKey:SessionID;constraint:OnDelete:SET NULL;"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}
