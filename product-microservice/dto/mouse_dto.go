package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type MouseDTO struct {
	Id	uuid.UUID
	Name string
	Wireless bool
	Sensor string
	DPI string
	PollingRate string
	Connection string
	Color string
	TrackingSpeed string
	Acceleration string
	Buttons uint
	Weight string
	Lifespan string
	Manufacturer string
	Price decimal.Decimal
	Amount uint
} 