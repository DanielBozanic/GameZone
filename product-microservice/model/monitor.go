package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Monitor struct {
	Id	uuid.UUID	`gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);unique;not null"`
	Size string `gorm:"type:varchar(40);not null"`
	AspectRatio string `gorm:"type:varchar(40);not null"`
	Resolution string `gorm:"type:varchar(40);not null"`
	ContrastRatio string `gorm:"type:varchar(40);not null"`
	ResponseTime string `gorm:"type:varchar(30);not null"`
	PanelType string `gorm:"type:varchar(30);not null"`
	ViewingAngle string `gorm:"type:varchar(30);not null"`
	Brightness string `gorm:"type:varchar(20);not null"`
	RefreshRate string `gorm:"type:varchar(20);not null"`
	Manufacturer string `gorm:"type:varchar(40);not null"`
	Price decimal.Decimal `gorm:"type:numeric;not null"`
	Amount uint `gorm:"not null"`
}