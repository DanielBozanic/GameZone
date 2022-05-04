package model

import (
	"github.com/google/uuid"
)

type GraphicsCard struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId uuid.UUID `gorm:"primaryKey"`
	ChipManufacturer string `gorm:"type:varchar(40);not null"`
	ModelName string `gorm:"type:varchar(100);not null"`
	BusWidth string  `gorm:"type:varchar(30);not null"`
	MemorySize string `gorm:"type:varchar(30);not null"`
	MemoryType string `gorm:"type:varchar(30);not null"`
	PCIInterface string `gorm:"type:varchar(40);not null"`
	GPUSpeed string `gorm:"type:varchar(20);not null"`
	CUDAStreamProcessors uint `gorm:"not null"`
	Cooling string `gorm:"type:varchar(20);not null"`
	HDMI uint `gorm:"not null"`
	DisplayPort uint `gorm:"not null"`
	TDP string `gorm:"type:varchar(30);not null"`
	PowerConnector string `gorm:"type:varchar(30);not null"`
	Dimensions string `gorm:"type:varchar(40);not null"`
}