package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Visit struct {
	gorm.Model
	ClientID             uint `gorm:"not null"`
	MasterID             *uint
	AddressID            uint       `gorm:"not null"`
	ContractNumber       string     `gorm:"size:100;not null"`
	ContractDate         time.Time  `gorm:"type:DATE;not null"`
	ScheduledDate        *time.Time `gorm:"type:DATE;"`
	ScheduledTime        string     `gorm:"size:50"`
	EquipmentDescription string
	AssignedMonth        string          `gorm:"size:20"`
	Amount               decimal.Decimal `gorm:"type:numeric(10,2)"`

	Client  Client  `gorm:"foreignKey:ClientID;references:ID"`
	Address Address `gorm:"foreignKey:AddressID;references:ID"`
}
