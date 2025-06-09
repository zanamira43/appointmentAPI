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
	Password  []byte `json:"-"`
	RoleID    uint   `json:"role_id"`

	Role Role `json:"role" gorm:"foreignKey:RoleID;constraint:OnDelete:SET NULL;"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}

func (user *User) ComparedPassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
