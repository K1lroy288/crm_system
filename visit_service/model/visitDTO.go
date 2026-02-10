package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type VisitDTO struct {
	FirstName            string          `json:"firstname"`
	LastName             string          `json:"lastname"`
	Surname              string          `json:"surname"`
	Phone                string          `json:"phone"`
	MasterID             *uint           `json:"masterid"`
	City                 string          `json:"city"`
	Locality             string          `json:"locality"`
	Region               string          `json:"region"`
	Street               string          `json:"street"`
	HouseNumber          uint            `json:"housenumber"`
	Letter               string          `json:"letter"`
	Building             *uint           `json:"building"`
	Appartment_number    *uint           `json:"appartmentnumber"`
	ContractNumber       string          `json:"contractnumber"`
	ContractDate         time.Time       `json:"contractdate"`
	ScheduledDate        *time.Time      `json:"scheduleddate"`
	ScheduledTime        string          `json:"scheduledtime"`
	EquipmentDescription string          `json:"equipmentdescription"`
	AssignedMonth        string          `json:"assignedmonth"`
	Amount               decimal.Decimal `json:"amount"`
}
