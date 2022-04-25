package model

import (
	"github.com/google/uuid"
)

type Mouse struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId	uuid.UUID	`gorm:"primaryKey"`
	Wireless bool `gorm:"type:bool;not null"`
	Sensor string `gorm:"type:varchar(30);not null"`
	DPI string `gorm:"type:varchar(40);not null"`
	PollingRate string `gorm:"type:varchar(30);not null"`
	Connection string `gorm:"type:varchar(30);not null"`
	Color string `gorm:"type:varchar(20);not null"`
	TrackingSpeed string `gorm:"type:varchar(30);not null"`
	Acceleration string `gorm:"type:varchar(30);not null"`
	Buttons uint `gorm:"not null"`
	Weight string `gorm:"type:varchar(20);not null"`
	Lifespan string `gorm:"type:varchar(30);not null"`
} 