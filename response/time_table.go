package response

// response for time table for day
type TimeTableForDay struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	PatientName string `json:"patient_name"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}
