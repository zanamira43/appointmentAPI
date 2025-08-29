package dto

type Session struct {
	PatientID   uint   `json:"patient_id"`
	Duration    int    `json:"duration"` // in minutes
	Status      string `json:"status"`
	Notes       string `json:"result"`
	SessionDate string `json:"session_date"`
}
