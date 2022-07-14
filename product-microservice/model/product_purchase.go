package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductPurchase struct {
	Id	uuid.UUID	`gorm:"primaryKey"`
	ProductId uuid.UUID `gorm:"index;ForeignKey:Id"`
	Product Product
	UserId  int `gorm:"type:int;not null"`
	Amount uint `gorm:"not null"`
	TotalPrice decimal.Decimal `gorm:"type:numeric;not null"`
	DeliveryAddress string `gorm:"type:varchar(50)"`
	City string `gorm:"type:varchar(50)"`
	MobilePhoneNumber string `gorm:"type:varchar(10)"`
	TypeOfPayment TypeOfPayment
	PurchaseDate time.Time `gorm:"default:null"`
	IsPayedFor bool `gorm:"default:false"`
}