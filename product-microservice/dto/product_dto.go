package dto

import (
	"product/model"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)


type ProductDTO struct {
	Id	uuid.UUID
	Name string
	Manufacturer string
	Price decimal.Decimal
	Amount uint
	Type model.Type
	Image string
}