package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductPurchaseDTO struct {
	ProductId uuid.UUID
	ProductName string
	TotalPrice decimal.Decimal
	Amount uint
}