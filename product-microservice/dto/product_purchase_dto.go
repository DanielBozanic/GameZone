package dto

import (
	"product/model"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductPurchaseDTO struct {
	Id uuid.UUID
	UserId int
	ProductId uuid.UUID
	ProductName string
	ProductPrice decimal.Decimal
	ProductImage string;
	TotalPrice decimal.Decimal
	Amount uint
	PurchaseDate string
	DeliveryAddress string
	TypeOfPayment model.TypeOfPayment
}