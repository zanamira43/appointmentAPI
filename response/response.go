package response

import "time"

// Response wrapper for paginated results
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Total      int64       `json:"total"`
	TotalPages int         `json:"total_pages"`
	HasNext    bool        `json:"has_next"`
	HasPrev    bool        `json:"has_prev"`
}

type PatientResponse struct {
	ID          uint   `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Profession  string `json:"profession"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`

	Payments  []PaymentsResposes `json:"payments"`
	Sessions  []SessionResposes  `json:"sessions"`
	CreatedAt time.Time          `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time          `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}

type SessionResposes struct {
	ID          uint   `json:"Id" gorm:"primaryKey"`
	PatientID   uint   `json:"patient_id"`
	Duration    int    `json:"duration"` // in minutes
	Status      string `json:"status"`
	Notes       string `json:"result"`
	SessionDate string `json:"session_date"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}

type PaymentsResposes struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	PatientID     uint   `json:"patient_id"`
	SessionID     uint   `json:"session_id"`
	Amount        int    `json:"amount"`
	PaymentDate   string `json:"payment_date"`
	PaymentMethod string `json:"payment_method"`
	Status        string `json:"status"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
}

// patient outcome Response struct
type PatientOutcomeResponse struct {
	ID                      uint    `json:"id"`
	Name                    string  `json:"name"`
	NeedSessionsCount       int     `json:"need_sessions_count"`
	IsDollarPaymnet         bool    `json:"is_dollar_payment"`
	SessionPrice            float32 `json:"session_price_one_month"`
	SumReceivedSessionCount int64   `json:"sum_received_session_count"`
	TotalReceivedPayments   float32 `json:"total_received_payments"`
}
