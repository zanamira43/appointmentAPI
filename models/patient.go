package models

import (
	"time"
)

type Patient struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Slug          string `json:"slug" gorm:"unique;index"`
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	Age           int    `json:"age"`
	MarriedStatus string `json:"married_status"`
	Profession    string `json:"profession"`
	Address       string `json:"address"`
	PhoneNumber   string `json:"phone_number"`

	Payments []Payment `json:"payments"`
	Sessions []Session `json:"sessions"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}
