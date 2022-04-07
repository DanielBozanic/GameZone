package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"product/model/enums"
)

type ConsoleDTO struct {
	Id uuid.UUID
	Name string 
	Price decimal.Decimal 
	Platform enums.Platform
	Manufacturer string
	Amount uint
}