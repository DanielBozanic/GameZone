package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type MotherboardDTO struct {
	Id	uuid.UUID
	Name string
	ProcessorType string
	Socket  string
	SupportedProcessors string
	Chipset string
	Memory  string
	MultiGraphicsTechnology string
	ExpansionSlots string
	StorageInterface string
	WirelessCommunicationModule string
	Audio string
	USB string
	BackPanelConnectors string
	InternalConnectors string
	BIOS string
	Manufacturer string
	Price decimal.Decimal
	Amount uint
}