package model

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	Id  int    `gorm:"primaryKey;auto_increment"`
	Name string `gorm:"type:varchar(100);unique;not null"`
	Description string `gorm:"type:varchar(1500);not null"`
	Manufacturer string `gorm:"type:varchar(40);not null"`
	Price decimal.Decimal `gorm:"type:numeric;not null"`
	Amount uint `gorm:"not null"`
	Type Type `gorm:"not null"`
	Image File `gorm:"foreignKey:ImageId"`
	ImageId int
	MainPage *bool `gorm:"type:boolean;default:false;not null"`
	Archived *bool `gorm:"type:boolean;default:false;not null"`
}