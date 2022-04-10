package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type MonitorDTO struct {
	Id	uuid.UUID
	Name string
	Size string
	AspectRatio string
	Resolution string
	ContrastRatio string
	ResponseTime string
	PanelType string
	ViewingAngle string
	Brightness string
	RefreshRate string
	Manufacturer string
	Price decimal.Decimal
	Amount uint
}