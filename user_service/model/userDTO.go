package model

type UserDTO struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Password  string    `json:"password"`
	Roles     []RoleDTO `json:"roles"`
}
