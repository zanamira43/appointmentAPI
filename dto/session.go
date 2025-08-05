package dto

import (
	"time"
)

type SessionDto struct {
	PatientID   uint      `json:"patient_id"`
	Duration    int       `json:"duration"` // in minutes
	Status      string    `json:"status"`
	Notes       string    `json:"result"`
	SessionDate time.Time `json:"session_date"`
}
