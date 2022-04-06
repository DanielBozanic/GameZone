package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"product/model/enums"
)

type VideoGameDTO struct {
	Id uuid.UUID
	Name string 
	Price decimal.Decimal 
	Digital bool
	Platform enums.Platform
	Publisher string
	Rating uint 
	Genre string
	ReleaseDate string
	Amount uint
}