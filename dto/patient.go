package dto

type Patient struct {
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	Age           int    `json:"age"`
	MarriedStatus string `json:"married_status"`
	Profession    string `json:"profession"`
	Address       string `json:"address"`
	PhoneNumber   string `json:"phone_number"`
}
