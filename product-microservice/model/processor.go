package model

import (
	"github.com/google/uuid"
)

type Processor struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId	uuid.UUID	`gorm:"primaryKey"`
	Type string `gorm:"type:varchar(40);not null"`
	Socket  string  `gorm:"type:varchar(30);not null"`
	NumberOfCores *uint
	Threads *uint
	TDP *string `gorm:"type:varchar(30)"`
	IntegratedGraphics *string `gorm:"type:varchar(30)"`
	BaseClockRate *string `gorm:"type:varchar(40)"`
	TurboClockRate *string `gorm:"type:varchar(40)"`
}