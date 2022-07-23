package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductPurchase struct {
	Id  int  `gorm:"primaryKey;auto_increment"`
	ProductId int
	Product Product `gorm:"foreignKey:ProductId"`
	UserId  int `gorm:"type:int;not null"`
	Amount uint `gorm:"not null"`
	TotalPrice decimal.Decimal `gorm:"type:numeric;not null"`
	DeliveryAddress string `gorm:"type:varchar(50)"`
	City string `gorm:"type:varchar(50)"`
	MobilePhoneNumber string `gorm:"type:varchar(10)"`
	TypeOfPayment TypeOfPayment
	PurchaseDate time.Time `gorm:"default:null"`
	IsPaidFor *bool `gorm:"default:false"`
}