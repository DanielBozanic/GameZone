package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductPurchase struct {
	Id	uuid.UUID	`gorm:"primaryKey"`
	UserId  int `gorm:"type:int;not null"`
	ProductId uuid.UUID `gorm:"not null"`
	ProductName string `gorm:"not null"`
	Amount uint `gorm:"not null"`
	TotalPrice decimal.Decimal `gorm:"type:numeric;not null"`
	PurchaseDate time.Time
}