package dto

import (
	"gorm.io/datatypes"
)

type Problem struct {
	PatientID         uint           `json:"patient_id"`
	MianpProblems     datatypes.JSON `json:"main_problems"`
	SecondaryProblems datatypes.JSON `json:"secondary_problems"`
	NeedSessionsCount int            `json:"need_sessions_count"`
	IsDollarPayment   bool           `json:"is_dollar_payment"`
	SessionPrice      float32        `json:"session_price_one_month"`
	PatientImage      string         `json:"patient_image"`
	Details           string         `json:"details" gorm:"type:longtext"`
}
