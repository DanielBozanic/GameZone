package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type VideoGame struct {
	Id	uuid.UUID	`gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);index:idx_name,unique;not null"`
	Price decimal.Decimal `gorm:"type:numeric;not null"`
	Digital bool `gorm:"type:bool;index:idx_name,unique;not null"`
	Platform string `gorm:"type:varchar(40);index:idx_name,unique;not null"`
	Publisher string `gorm:"type:varchar(50);not null"`
	Rating uint `gorm:"not null"`
	Genre string `gorm:"type:varchar(50); not null"`
	ReleaseDate time.Time `gorm:"not null"`
	Amount uint
}