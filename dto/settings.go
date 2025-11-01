package dto

type Settings struct {
	ID          uint   `json:"id"`
	SystemName  string `json:"system_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	BillPrefix  string `json:"bill_prefix"`
}
