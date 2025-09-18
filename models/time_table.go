package models

import (
	"time"

	"gorm.io/datatypes"
)

type TimeTable struct {
	ID          uint           `json:"id"`
	PatientID   *int           `json:"patient_id"`
	PatientName string         `json:"patient_name"`
	WeekDay     datatypes.JSON `json:"week_day"`
	StartTime   string         `json:"start_time"`
	EndTime     string         `json:"end_time"`
	UserId      uint           `json:"user_id"`

	Patient *Patient `json:"patient" gorm:"foreignKey:PatientID;constraint:OnDelete:SET NULL;"`
	User    User     `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:SET NULL;OnUpdate:SET NULL;"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}

func (TimeTable) TableName() string {
	return "time_tables"
}
