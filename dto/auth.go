package dto

// register dto just for validation
// and checking password confirmation
type Register struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email" gorm:"unique"`
	Phone           string `json:"phone_number" gorm:"unique"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirmation"`
}

// /////////////////
type Login struct {
	Phone    string `json:"phone_number" gorm:"unique"`
	Password string `json:"password"`
}
