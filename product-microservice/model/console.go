package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Console struct {
	Id	uuid.UUID	`gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);unique;not null"`
	Price decimal.Decimal `gorm:"type:numeric;not null"`
	Platform string `gorm:"type:varchar(40);not null"`
	Manufacturer string `gorm:"type:varchar(40);not null"`
	Amount uint `gorm:"not null"`
}