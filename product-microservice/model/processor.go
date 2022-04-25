package model

import (
	"github.com/google/uuid"
)

type Processor struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId	uuid.UUID	`gorm:"primaryKey"`
	Type string `gorm:"type:varchar(40);not null"`
	Socket  string  `gorm:"type:varchar(30);not null"`
	NumberOfCores uint `gorm:"not null"`
	Threads uint `gorm:"type:varchar(20);not null"`
	TDP string `gorm:"type:varchar(30);not null"`
	IntegratedGraphics string `gorm:"type:varchar(30);default=None"`
	BaseClockRate string `gorm:"type:varchar(40);not null"`
	TurboClockRate string `gorm:"type:varchar(40)"`
}