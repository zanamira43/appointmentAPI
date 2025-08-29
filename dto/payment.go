package dto

type Payment struct {
	PatientID     uint   `json:"patient_id"`
	SessionID     uint   `json:"session_id"`
	Amount        int    `json:"amount"`
	PaymentDate   string `json:"payment_date"`
	PaymentMethod string `json:"payment_method"`
	Status        string `json:"status"`
}
