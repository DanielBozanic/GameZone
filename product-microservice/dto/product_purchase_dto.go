package dto

import (
	"product/model"
	"time"

	"github.com/shopspring/decimal"
)

type ProductPurchaseDTO struct {
	Id int
	UserId int
	ProductPurchaseDetail []ProductPurchaseDetailDTO
	TotalPrice decimal.Decimal
	PurchaseDate time.Time
	DeliveryAddress string
	City string
	MobilePhoneNumber string
	TypeOfPayment model.TypeOfPayment
	IsPaidFor *bool
}