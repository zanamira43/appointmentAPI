package models

import "time"

// person info model
// it's 2 month course plan
type PersonInfo struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number" gorm:"unique"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}
