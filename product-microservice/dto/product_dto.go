package dto

import (
	"product/model"

	"github.com/shopspring/decimal"
)


type ProductDTO struct {
	Id	int
	Name string
	Description string
	Manufacturer string
	Price decimal.Decimal
	Amount uint
	Type model.Type
	Image FileDTO
	MainPage *bool
	Archived *bool
}