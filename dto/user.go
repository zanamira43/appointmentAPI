package dto

import (
	"golang.org/x/crypto/bcrypt"
)

type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Phone     string `json:"phone_number" gorm:"unique"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Active    bool   `json:"active"`
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Phone     string `json:"phone_number" gorm:"unique"`
	Password  []byte `json:"-"`
	Role      string `json:"role"`
	Active    bool   `json:"active"`
}

type UserPassword struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirmation"`
}

// hashing user password before store it in database
func (user *User) SetPassword(password string) {
	HashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = HashedPassword
}
