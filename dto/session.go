package dto

type Session struct {
	PatientID          uint   `json:"patient_id"`
	Subject            string `json:"subject"`
	CommunicationTypes string `json:"communication_types"`
	SessionDate        string `json:"session_date"`
	Detail             string `json:"detail"`
	Status             string `json:"status"`
}
