package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	Id	uuid.UUID	`gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);unique;not null"`
	Manufacturer string `gorm:"type:varchar(40);not null"`
	Price decimal.Decimal `gorm:"type:numeric;not null"`
	Amount uint `gorm:"not null"`
	Type Type `gorm:"not null"`
}