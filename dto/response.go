package dto

type LoginResponse struct {
	Token string `json:"token"`
	ID    uint   `json:"id"`
}
