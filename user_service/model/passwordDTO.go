package model

type PasswordDTO struct {
	CurrPass    string `json:"current_password"`
	NewPassword string `json:"new_password"`
}
