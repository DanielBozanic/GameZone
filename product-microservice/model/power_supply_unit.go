package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PowerSupplyUnit struct {
	Id	uuid.UUID	`gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);unique;not null"`
	PowerRating string `gorm:"type:varchar(40);not null"`
	Type string `gorm:"type:varchar(40);not null"`
	FormFactor string `gorm:"type:varchar(40);not null"`
	Manufacturer string `gorm:"type:varchar(40);not null"`
	Price decimal.Decimal `gorm:"type:numeric;not null"`
	Amount uint `gorm:"not null"`
}