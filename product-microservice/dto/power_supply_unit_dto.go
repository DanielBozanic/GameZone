package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PowerSupplyUnitDTO struct {
	Id	uuid.UUID
	Name string
	PowerRating string
	Type string
	FormFactor string
	Manufacturer string
	Price decimal.Decimal
	Amount uint
}