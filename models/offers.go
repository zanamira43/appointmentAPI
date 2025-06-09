package models

import "time"

type Offer struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Title         string `json:"title"`
	ServiceTypeID uint   `json:"service_type_id"`
	Price         int    `json:"price"`

	ServiceType ServiceType `json:"service_type" gorm:"foreignKey:ServiceTypeID;constraint:OnDelete:SET NULL;"`
	Payments    []Payment   `json:"payments"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}
