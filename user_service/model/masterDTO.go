package model

type MasterDTO struct {
	ID        int    `json:"master_id"`
	Username  string `json:"master_username"`
	FirstName string `json:"master_firstname"`
	LastName  string `json:"master_lastname"`
}
