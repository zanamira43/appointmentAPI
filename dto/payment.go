package dto

type PaymentType struct {
	Name string `json:"name"`
}

type Payment struct {
	PatientID       uint   `json:"patient_id"`
	PaymentTypeID   uint   `json:"payment_type_id"`
	IsDollarPayment bool   `json:"is_dollar_payment"`
	Amount          int    `json:"amount"`
	PaymentDate     string `json:"payment_date"`
}
