package model

import (
	"github.com/google/uuid"
)

type Headphones struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId uuid.UUID	`gorm:"primaryKey"`
	Description string `gorm:"type:text;not null"`
	VirtualSurroundEncoding string `gorm:"type:varchar(30);not null"`
	Sensitivity string `gorm:"type:varchar(30);not null"`
	ConnectionType string `gorm:"type:varchar(40);not null"`
	Wireless bool `gorm:"type:bool;not null"`
	DriverSize string `gorm:"type:varchar(20);not null"`
	Microphone bool `gorm:"type:bool;not null"`
	Color string `gorm:"type:varchar(20);not null"`
	Weight string `gorm:"type:varchar(20);not null"`
	FrequencyResponse string `gorm:"type:varchar(30);not null"`
}