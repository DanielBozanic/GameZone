package model

import (
	"github.com/google/uuid"
)

type PowerSupplyUnit struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId	uuid.UUID	`gorm:"primaryKey"`
	PowerRating string `gorm:"type:varchar(40);not null"`
	Type string `gorm:"type:varchar(40);not null"`
	FormFactor string `gorm:"type:varchar(40);not null"`
}