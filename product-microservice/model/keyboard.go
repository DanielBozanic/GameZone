package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Keyboard struct {
	Id	uuid.UUID	`gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);unique;not null"`
	Wireless bool `gorm:"type:bool;not null"`
	KeyboardConnector string `gorm:"type:varchar(30);not null"`
	KeyType string `gorm:"type:varchar(30);not null"`
	LetterLayout string `gorm:"type:varchar(20);not null"`
	KeyboardColor string `gorm:"type:varchar(20);not null"`
	Manufacturer string `gorm:"type:varchar(40);not null"`
	Price decimal.Decimal `gorm:"type:numeric;not null"`
	Amount uint `gorm:"not null"`
}