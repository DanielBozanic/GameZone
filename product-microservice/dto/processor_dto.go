package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProcessorDTO struct {
	Id	uuid.UUID
	Name string
	Type string
	Manufacturer string
	Socket  string
	NumberOfCores uint
	Threads uint
	TDP string
	IntegratedGraphics string
	BaseClockRate string
	TurboClockRate string
	Price decimal.Decimal
	Amount uint
}