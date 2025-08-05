package dto

import "time"

type PaymentDto struct {
	PatientID     uint      `json:"patient_id"`
	SessionID     uint      `json:"session_id"`
	Amount        int       `json:"amount"`
	PaymentDate   time.Time `json:"payment_date"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status"`
}
