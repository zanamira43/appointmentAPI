package dto

type Image struct {
	PatientImageUrl string `json:"patient_image_url"`
}

type SignatureFileImage struct {
	SignatureFileUrl string `json:"signature_file_url"`
}
