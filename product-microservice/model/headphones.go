package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Headphones struct {
	Id	uuid.UUID	`gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);unique;not null"`
	Description string `gorm:"type:text;not null"`
	VirtualSurroundEncoding string `gorm:"type:varchar(30);not null"`
	Sensitivity string `gorm:"type:varchar(30);not null"`
	ConnectionType string `gorm:"type:varchar(40);not null"`
	Wireless bool `gorm:"type:bool;not null"`
	DriverSize string `gorm:"type:varchar(20);not null"`
	Microphone bool `gorm:"type:bool;not null"`
	Color string `gorm:"type:varchar(20);not null"`
	Weight string `gorm:"type:varchar(20);not null"`
	FrequencyResponse string `gorm:"type:varchar(30);not null"`
	Manufacturer string `gorm:"type:varchar(40);not null"`
	Price decimal.Decimal `gorm:"type:numeric;not null"`
	Amount uint `gorm:"not null"`
}