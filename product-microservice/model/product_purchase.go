package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductPurchase struct {
	Id  int  `gorm:"primaryKey;auto_increment"`
	ProductPurchaseDetail []ProductPurchaseDetail `gorm:"foreignKey:ProductPurchaseId"`
	UserId  int `gorm:"type:int;not null"`
	TotalPrice decimal.Decimal `gorm:"type:numeric;not null"`
	DeliveryAddress string `gorm:"type:varchar(50);not null"`
	City string `gorm:"type:varchar(50);not null"`
	MobilePhoneNumber string `gorm:"type:varchar(10);not null"`
	TypeOfPayment TypeOfPayment `gorm:"not null"`
	PurchaseDate time.Time `gorm:"not null"`
	IsPaidFor *bool `gorm:"default:false;not null"`
}