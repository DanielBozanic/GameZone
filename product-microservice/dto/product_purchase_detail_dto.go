package dto

import "github.com/shopspring/decimal"

type ProductPurchaseDetailDTO struct {
	Id int
	ProductId int
	ProductPurchaseId int
	ProductName  string
	ProductPrice decimal.Decimal
	ProductQuantity uint
}