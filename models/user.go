package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `json:"Id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Phone     string `json:"phone" gorm:"unique"`
	Password  []byte `json:"-"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}

func (user *User) ComparedPassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
