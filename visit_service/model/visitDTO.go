package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type VisitDTO struct {
	Client struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Surname   string `json:"surname"`
		Phone     string `json:"phone"`
	} `json:"client"`

	Address struct {
		City             string `json:"city"`
		Locality         string `json:"locality"`
		Region           string `json:"region"`
		Street           string `json:"street"`
		HouseNumber      uint   `json:"house_number"`
		Letter           string `json:"letter"`
		Building         *uint  `json:"building"`
		AppartmentNumber *uint  `json:"appartment_number"`
	} `json:"address"`

	ID                   uint            `json:"id"`
	ContractNumber       string          `json:"contract_number"`
	ContractDate         time.Time       `json:"contract_date"`
	ScheduledDate        *time.Time      `json:"scheduled_date"`
	ScheduledTime        string          `json:"scheduled_time"`
	EquipmentDescription string          `json:"equipment_description"`
	AssignedMonth        string          `json:"assigned_month"`
	Amount               decimal.Decimal `json:"amount"`
	MasterID             *uint           `json:"master_id"`
}
