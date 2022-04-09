package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Motherboard struct {
	Id	uuid.UUID	`gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);unique;not null"`
	ProcessorType string `gorm:"type:varchar(30);not null"`
	Socket  string  `gorm:"type:varchar(30);not null"`
	SupportedProcessors string `gorm:"type:text;not null"`
	Chipset string  `gorm:"type:varchar(30);not null"`
	Memory  string `gorm:"type:text;not null"`
	MultiGraphicsTechnology string `gorm:"type:text;not null"`
	ExpansionSlots string  `gorm:"type:text;not null"`
	StorageInterface string `gorm:"type:text;not null"`
	WirelessCommunicationModule string `gorm:"type:text;not null"`
	Audio string `gorm:"type:text;not null"`
	USB string `gorm:"type:text;not null"`
	BackPanelConnectors string `gorm:"type:text;not null"`
	InternalConnectors string `gorm:"type:text;not null"`
	BIOS string `gorm:"type:text;not null"`
	Manufacturer string `gorm:"type:varchar(40);not null"`
	Price decimal.Decimal `gorm:"type:numeric;not null"`
	Amount uint `gorm:"not null"`
}