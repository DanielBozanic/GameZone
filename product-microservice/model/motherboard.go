package model

import (
	"github.com/google/uuid"
)

type Motherboard struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId	uuid.UUID	`gorm:"primaryKey"`
	ProcessorType string `gorm:"type:varchar(30);not null"`
	Socket  string  `gorm:"type:varchar(30);not null"`
	SupportedProcessors *string `gorm:"type:varchar(600)"`
	Chipset *string  `gorm:"type:varchar(30)"`
	Memory  *string `gorm:"type:varchar(1000)"`
	ExpansionSlots *string `gorm:"type:varchar(400)"`
	StorageInterface *string `gorm:"type:varchar(200)"`
	Audio *string `gorm:"type:varchar(200)"`
	USB *string `gorm:"type:varchar(600)"`
	BackPanelConnectors *string `gorm:"type:varchar(1000)"`
	InternalConnectors *string `gorm:"type:varchar(1000)"`
	BIOS *string `gorm:"type:varchar(400)"`
	FormFactor *string `gorm:"type:varchar(40)"`
}