package model

import (
	"github.com/google/uuid"
)

type Motherboard struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId	uuid.UUID	`gorm:"primaryKey"`
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
	FormFactor string `gorm:"type:varchar(40);not null"`
}