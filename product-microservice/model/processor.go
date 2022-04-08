package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Processor struct {
	Id	uuid.UUID	`gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);unique;not null"`
	Type string `gorm:"type:varchar(40);not null"`
	Manufacturer string `gorm:"type:varchar(30);not null"`
	Socket  string  `gorm:"type:varchar(30);not null"`
	NumberOfCores uint `gorm:"not null"`
	Threads uint `gorm:"type:varchar(20);not null"`
	TDP string `gorm:"type:varchar(30);not null"`
	IntegratedGraphics string `gorm:"type:varchar(30);default=None"`
	BaseClockRate string `gorm:"type:varchar(40);not null"`
	TurboClockRate string `gorm:"type:varchar(40)"`
	Price decimal.Decimal `gorm:"type:numeric;not null"`
	Amount uint `gorm:"not null"`
}