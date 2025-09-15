package models

import "time"

type AppointmentSchedule struct {
	ID          uint   `json:"id"`
	PatientID   *uint  `json:"patient_id"`
	PatientName string `json:"patient_name"`
	WeekDay     string `json:"week_day"`
	Date        string `json:"date"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Status      string `json:"status"`
	UserId      uint   `json:"user_id"`

	Patient Patient `json:"patient" gorm:"foreignKey:PatientID;constraint:OnDelete:SET NULL;"`
	User    User    `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:SET NULL;"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}

func (AppointmentSchedule) TableName() string {
	return "appointment_schedule"
}
