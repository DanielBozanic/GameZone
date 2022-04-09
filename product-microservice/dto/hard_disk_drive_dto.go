package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type HardDiskDriveDTO struct {
	Id	uuid.UUID
	Name string
	Capacity string
	DiskSpeed string
	Interface string
	TransferRate string
	Form string
	Price decimal.Decimal
	Manufacturer string
	Amount uint
}