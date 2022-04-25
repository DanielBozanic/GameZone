package model

import (
	"github.com/google/uuid"
)

type HardDiskDrive struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId uuid.UUID `gorm:"primaryKey"`
	Capacity string `gorm:"type:varchar(30);not null"`
	DiskSpeed string `gorm:"type:varchar(30);not null"`
	Interface string `gorm:"type:varchar(30);not null"`
	TransferRate string `gorm:"type:varchar(30);not null"`
	Form string `gorm:"type:varchar(30);not null"`
}