package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Ram struct {
	Id	uuid.UUID	`gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);unique;not null"`
	MemoryType string `gorm:"type:varchar(30);not null"`
	Capacity string `gorm:"type:varchar(30);not null"`
	Speed string `gorm:"type:varchar(30);not null"`
	Voltage string `gorm:"type:varchar(30);not null"`
	Latency string `gorm:"type:varchar(30);not null"`
	Price decimal.Decimal `gorm:"type:numeric;not null"`
	Manufacturer string `gorm:"type:varchar(40);not null"`
	Amount uint `gorm:"not null"`
}