package dto

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Phone     string `json:"phone" gorm:"unique"`
	Password  []byte `json:"-"`
}

// hashing user password before store it in database
func (user *User) SetPassword(password string) {
	HashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = HashedPassword
}
