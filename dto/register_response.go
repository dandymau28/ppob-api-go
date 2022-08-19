package dto

import "time"

type RegisterResponse struct {
	ID          string    `json:"id"`
	NoHandphone string    `json:"noHandphone"`
	MacAddress  string    `json:"macAddress"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updateAt"`
}
