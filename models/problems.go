package models

import "time"

type Problem struct {
	ID              uint     `json:"Id" gorm:"primaryKey"`
	PatientID       uint     `json:"patient_id"`
	Titles          []string `json:"titles" gorm:"type:json"`
	Effects         []string `json:"effects" gorm:"type:json"`
	Durations       string   `json:"durations" gorm:"type:varchar(255)"`
	IsHeartBeatFast bool     `json:"is_heart_beat_fast" gorm:"type:boolean;default:false"`
	IsHateOrAngry   bool     `json:"is_heat_or_angry" gorm:"type:boolean;default:false"`
	HaveDrugs       string   `json:"have_drugs" gorm:"type:varchar(255)"`

	Patient  Patient   `json:"patient" gorm:"foreignKey:PatientID;constraint:onDelete:SET NULL;"`
	Sessions []Session `json:"sessions"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}
