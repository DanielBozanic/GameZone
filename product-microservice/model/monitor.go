package model

import (
	"github.com/google/uuid"
)

type Monitor struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId	uuid.UUID	`gorm:"primaryKey"`
	Size string `gorm:"type:varchar(40);not null"`
	AspectRatio string `gorm:"type:varchar(40);not null"`
	Resolution string `gorm:"type:varchar(40);not null"`
	ContrastRatio string `gorm:"type:varchar(40);not null"`
	ResponseTime string `gorm:"type:varchar(30);not null"`
	PanelType string `gorm:"type:varchar(30);not null"`
	ViewingAngle string `gorm:"type:varchar(30);not null"`
	Brightness string `gorm:"type:varchar(20);not null"`
	RefreshRate string `gorm:"type:varchar(20);not null"`
}