package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type VideoGameDTO struct {
	Id uuid.UUID
	Name string 
	Price decimal.Decimal 
	Digital bool
	Platform string
	Publisher string
	Rating uint 
	Genre string
	ReleaseDate string
	Amount uint
}