package model

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	City              string `gorm:"size:100;not null"`
	Locality          string `gorm:"size:100"`
	Region            string `gorm:"size:100;not null"`
	Street            string `gorm:"size:100;not null"`
	HouseNumber       uint   `gorm:"not null"`
	Letter            string `gorm:"size:10;"`
	Building          *uint
	Appartment_number *uint
}
