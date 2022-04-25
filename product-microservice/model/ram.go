package model

import (
	"github.com/google/uuid"
)

type Ram struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId	uuid.UUID	`gorm:"primaryKey"`
	MemoryType string `gorm:"type:varchar(30);not null"`
	Capacity string `gorm:"type:varchar(30);not null"`
	Speed string `gorm:"type:varchar(30);not null"`
	Voltage string `gorm:"type:varchar(30);not null"`
	Latency string `gorm:"type:varchar(30);not null"`
}