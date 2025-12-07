package models

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	UserID        uint   `json:"user_id"`
	Slug          string `json:"slug" gorm:"unique;index"`
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	Age           int    `json:"age"`
	MarriedStatus string `json:"married_status"`
	Profession    string `json:"profession"`
	Address       string `json:"address"`
	PhoneNumber   string `json:"phone_number"`

	Problem Problem `json:"problem"`

	Payments []Payment `json:"payments"`
	Sessions []Session `json:"sessions"`

	User User `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL;OnUpdate:SET NULL;"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}

func (p *Patient) BeforeCreate(tx *gorm.DB) error {
	// Generate a random number and convert it to a string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := r.Intn(89999999) + 10000000 // Adjust the range as needed

	p.Slug = strconv.Itoa(int(randomNum))

	if p.Slug == "" {
		return errors.New("slug is required")
	}

	return nil
}
