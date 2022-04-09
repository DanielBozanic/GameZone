package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type SolidStateDriveDTO struct {
	Id	uuid.UUID
	Name string
	Capacity string
	Interface string
	MaxSequentialRead string
	MaxSequentialWrite string
	Form string
	Dimensions string
	Price decimal.Decimal
	Manufacturer string
	Amount uint
}