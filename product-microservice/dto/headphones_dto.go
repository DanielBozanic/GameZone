package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type HeadphonesDTO struct {
	Id	uuid.UUID
	Name string
	Description string
	VirtualSurroundEncoding string
	Sensitivity string
	ConnectionType string
	Wireless bool
	DriverSize string
	Microphone bool
	Color string
	Weight string
	FrequencyResponse string
	Manufacturer string
	Price decimal.Decimal
	Amount uint
}