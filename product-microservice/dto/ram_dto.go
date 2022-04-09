package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type RamDTO struct {
	Id	uuid.UUID
	Name string 
	MemoryType string
	Capacity string
	Speed string
	Voltage string
	Latency string
	Price decimal.Decimal
	Manufacturer string
	Amount uint
}