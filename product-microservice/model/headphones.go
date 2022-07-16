package model

import (
	"github.com/google/uuid"
)

type Headphones struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId uuid.UUID	`gorm:"primaryKey"`
	ConnectionType string `gorm:"type:varchar(40);not null"`
	Wireless bool `gorm:"type:bool;not null"`
	Microphone bool `gorm:"type:bool;not null"`
	VirtualSurroundEncoding *string `gorm:"type:varchar(30)"`
	Sensitivity *string `gorm:"type:varchar(30)"`
	DriverSize *string `gorm:"type:varchar(20)"`
	Color *string `gorm:"type:varchar(20)"`
	Weight *string `gorm:"type:varchar(20)"`
	FrequencyResponse *string `gorm:"type:varchar(30)"`
}