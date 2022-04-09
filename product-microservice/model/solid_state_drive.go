package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type SolidStateDrive struct {
	Id	uuid.UUID	`gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);unique;not null"`
	Capacity string `gorm:"type:varchar(30);not null"`
	Interface string  `gorm:"type:varchar(30);not null"`
	MaxSequentialRead string `gorm:"type:varchar(30);not null"`
	MaxSequentialWrite string `gorm:"type:varchar(30);not null"`
	Form string `gorm:"type:varchar(30);not null"`
	Dimensions string `gorm:"type:varchar(40);not null"`
	Price decimal.Decimal `gorm:"type:numeric;not null"`
	Manufacturer string `gorm:"type:varchar(40);not null"`
	Amount uint `gorm:"not null"`
}