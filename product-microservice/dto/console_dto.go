package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ConsoleDTO struct {
	Id uuid.UUID
	Name string 
	Price decimal.Decimal 
	Platform string
	Manufacturer string
	Amount uint
}