package models

import (
	"time"
)

type Patient struct {
	ID               uint   `json:"id" gorm:"primaryKey"`
	Slug             string `json:"slug" gorm:"unique;index"`
	Name             string `json:"name"`
	Gender           string `json:"gender"`
	Age              int    `json:"age"`
	Profession       string `json:"profession"`
	Address          string `json:"address"`
	PhoneNumber      string `json:"phone_number"`
	ServiceType      string `json:"service_type"`
	HowHeardAboutUse string `json:"how_heard_about_us"`

	Payments []Payment `json:"payments"`
	Problems []Problem `json:"problems"`
	Sessions []Session `json:"sessions"`
	WeekDays []WeekDay `json:"week_days" gorm:"many2many:patient_week_days;"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}
