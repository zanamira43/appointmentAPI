package models

import "time"

type WeekDay struct {
	ID   uint   `json:"Id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`

	Patients  []Patient `json:"patients" gorm:"many2many:patient_week_days;"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}
