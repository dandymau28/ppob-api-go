package dto

type RegisterRequest struct {
	NoHandphone string `json:"no_handphone" validate:"required"`
	MacAddress  string `json:"mac_address" validate:"required"`
}
