package models

import (
	"time"

	"gorm.io/datatypes"
)

type Problem struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	PatientID         uint           `json:"patient_id"`
	MianpProblems     datatypes.JSON `json:"main_problems"`
	SecondaryProblems datatypes.JSON `json:"secondary_problems"`
	NeedSessionsCount int            `json:"need_sessions_count"`
	IsDollarPayment   bool           `json:"is_dollar_payment"`
	SessionPrice      float32        `json:"session_price"`
	SessionTotalPrice float32        `json:"session_total_price"`
	PatientImage      string         `json:"patient_image"`
	Details           string         `json:"details" gorm:"type:longtext"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}
