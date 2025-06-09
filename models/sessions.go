package models

import "time"

type Session struct {
	ID        uint   `json:"Id" gorm:"primaryKey"`
	PatientID uint   `json:"patient_id"`
	ProblemID uint   `json:"problem_id"`
	Result    string `json:"result"`
	Status    string `json:"status"`

	Patient Patient `json:"patient" gorm:"foreignKey:PatientID;constraint:OnDelete:SET NULL;"`
	Problem Problem `json:"Problem" gorm:"foreignKey:ProblemID;constraint:OnDelete:SET NULL;"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}
