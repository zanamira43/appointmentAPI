package models

import "time"

type Session struct {
	ID          uint      `json:"Id" gorm:"primaryKey"`
	PatientID   uint      `json:"patient_id"`
	Duration    int       `json:"duration"` // in minutes
	Status      string    `json:"status"`
	Notes       string    `json:"result"`
	SessionDate time.Time `json:"session_date"`
	Patient     Patient   `json:"patient" gorm:"foreignKey:PatientID;constraint:OnDelete:SET NULL;"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}
