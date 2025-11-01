package models

import "gorm.io/gorm"

type Settings struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:false;default:1" json:"id"`
	SystemName  string `json:"system_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	BillPrefix  string `json:"bill_prefix"`
}

// BeforeCreate hook to ensure only one record exists
func (s *Settings) BeforeCreate(tx *gorm.DB) error {
	s.ID = 1
	return nil
}
