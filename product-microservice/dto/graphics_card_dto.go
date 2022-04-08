package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type GraphicsCardDTO struct {
	Id	uuid.UUID	
	Name string
	ModelName string
	BusWidth string
	MemorySize string
	MemoryType string
	PCIInterface string
	GPUSpeed string
	Manufacturer string
	CUDAStreamProcessors uint
	Cooling string
	HDMI uint
	DisplayPort uint
	TDP string
	PowerConnector string
	Dimensions string
	Price decimal.Decimal
	Amount uint
}