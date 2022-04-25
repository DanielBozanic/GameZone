package model

import (
	"github.com/google/uuid"
)

type Keyboard struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId uuid.UUID `gorm:"primaryKey"`
	Wireless bool `gorm:"type:bool;not null"`
	KeyboardConnector string `gorm:"type:varchar(30);not null"`
	KeyType string `gorm:"type:varchar(30);not null"`
	LetterLayout string `gorm:"type:varchar(20);not null"`
	KeyboardColor string `gorm:"type:varchar(20);not null"`
}