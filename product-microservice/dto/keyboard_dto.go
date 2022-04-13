package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type KeyboardDTO struct {
	Id	uuid.UUID
	Name string
	Wireless bool
	KeyboardConnector string
	KeyType string
	KeyboardColor string
	LetterLayout string
	Manufacturer string
	Price decimal.Decimal
	Amount uint
}