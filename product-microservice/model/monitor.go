package model

import (
	"github.com/google/uuid"
)

type Monitor struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId	uuid.UUID	`gorm:"primaryKey"`
	Resolution string `gorm:"type:varchar(40);not null"`
	RefreshRate string `gorm:"type:varchar(20);not null"`
	Size string `gorm:"type:varchar(40);not null"`
	AspectRatio *string `gorm:"type:varchar(40)"`
	ContrastRatio *string `gorm:"type:varchar(40)"`
	ResponseTime *string `gorm:"type:varchar(30)"`
	PanelType *string `gorm:"type:varchar(30)"`
	ViewingAngle *string `gorm:"type:varchar(30)"`
	Brightness *string `gorm:"type:varchar(20)"`
}