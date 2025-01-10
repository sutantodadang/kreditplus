package customers

type CreateCustomerRequest struct {
	IdentificationNumber string          `json:"identification_number"`
	Fullname             string          `json:"full_name"`
	LegalName            string          `json:"legal_name"`
	PlaceOfBirth         string          `json:"place_of_birth"`
	DateOfBirth          string          `json:"date_of_birth"`
	Salary               string          `json:"salary"`
	PhotoKTP             string          `json:"photo_ktp"`
	PhotoSelfie          string          `json:"photo_selfie"`
	UserID               string          `json:"user_id"`
	CustomerLimits       []CustomerLimit `json:"customer_limits"`
}

type CustomerLimit struct {
	Tenor       int     `json:"tenor"`
	LimitAmount float64 `json:"limit_amount"`
}
