package model

import (
	"time"
)

type VideoGame struct {
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId  int    `gorm:"primaryKey;auto_increment"`
	Digital *bool `gorm:"type:bool;not null"`
	Platform string `gorm:"type:varchar(40);not null"`
	Rating uint `gorm:"not null"`
	Genre string `gorm:"type:varchar(50);not null"`
	ReleaseDate *time.Time
}