package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type VideoGame struct {
	gorm.Model
	Id	uuid.UUID	`gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);unique;not null"`
	Price decimal.Decimal `gorm:"type:numeric;not null"`
	Digital bool `gorm:"type:bool;not null"`
	Platform string `gorm:"type:varchar(30);not null"`
	Publisher string `gorm:"type:varchar(50);not null"`
	Rating uint 
	Genre string `gorm:"type:varchar(50);not null"`
	ReleaseDate time.Time
	Amount uint
  }