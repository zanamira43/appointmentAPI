package dto

type AppointmentSchedule struct {
	PatientID   uint   `json:"patient_id"`
	PatientName string `json:"patient_name"`
	WeekDay     string `json:"week_day"`
	Date        string `json:"date"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Status      string `json:"status"`
}
