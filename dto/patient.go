package dto

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type PatientDto struct {
	Slug        string `json:"slug" gorm:"unique;index"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Profession  string `json:"profession"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

func (p *PatientDto) BeforeCreate(tx *gorm.DB) error {
	// Generate a random number and convert it to a string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := r.Intn(89999999) + 10000000 // Adjust the range as needed

	p.Slug = strconv.Itoa(int(randomNum))

	if p.Slug == "" {
		return errors.New("slug is required")
	}

	return nil
}
