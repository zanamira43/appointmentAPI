package models

import "time"

type ServiceType struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`

	Offers   []Offer   `json:"offers"`
	Payments []Payment `json:"payments"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}
