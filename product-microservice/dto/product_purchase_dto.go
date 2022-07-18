package dto

import (
	"product/model"

	"github.com/shopspring/decimal"
)

type ProductPurchaseDTO struct {
	Id int
	UserId int
	Product ProductDTO
	TotalPrice decimal.Decimal
	Amount uint
	PurchaseDate string
	DeliveryAddress string
	City string
	MobilePhoneNumber string
	TypeOfPayment model.TypeOfPayment
	IsPayedFor bool
}