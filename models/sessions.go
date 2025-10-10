package models

import "time"

type Session struct {
	ID                 uint   `json:"id" gorm:"primaryKey"`
	PatientID          uint   `json:"patient_id"`
	Subject            string `json:"subject"`
	CommunicationTypes string `json:"communication_types"`
	SessionDate        string `json:"session_date"`
	Duration           int    `json:"duration"` // in minutes
	Status             string `json:"status"`

	Patient   Patient   `json:"-" gorm:"foreignKey:PatientID;constraint:OnDelete:SET NULL;"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}
