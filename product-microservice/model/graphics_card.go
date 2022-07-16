package model

import (
	"github.com/google/uuid"
)

type GraphicsCard struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId uuid.UUID `gorm:"primaryKey"`
	ChipManufacturer string `gorm:"type:varchar(40);not null"`
	ModelName string `gorm:"type:varchar(100);not null"`
	BusWidth *string  `gorm:"type:varchar(30)"`
	MemorySize *string `gorm:"type:varchar(30)"`
	MemoryType *string `gorm:"type:varchar(30)"`
	PCIInterface *string `gorm:"type:varchar(40)"`
	GPUSpeed *string `gorm:"type:varchar(20)"`
	CUDAStreamProcessors *uint
	Cooling *string `gorm:"type:varchar(20)"`
	HDMI *uint
	DisplayPort *uint
	TDP *string `gorm:"type:varchar(30)"`
	PowerConnector *string `gorm:"type:varchar(30)"`
	Dimensions *string `gorm:"type:varchar(40)"`
}