package model

import (
	"github.com/google/uuid"
)

type Console struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId uuid.UUID `gorm:"primaryKey"`
	Platform string `gorm:"type:varchar(40);not null"`
}