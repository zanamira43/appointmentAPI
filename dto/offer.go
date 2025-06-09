package dto

type Offer struct {
	Title         string `json:"title"`
	ServiceTypeID uint   `json:"service_type_id"`
	Price         int    `json:"price"`
}

type ServiceType struct {
	Name string `json:"name"`
}
