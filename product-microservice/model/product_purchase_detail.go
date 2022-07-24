package model

import "github.com/shopspring/decimal"

type ProductPurchaseDetail struct {
	Id  int  `gorm:"primaryKey;auto_increment"`
	ProductPurchaseId int
	ProductId int
	ProductName  string `gorm:"type:varchar(100);not null"`
	ProductPrice decimal.Decimal `gorm:"type:numeric;not null"`
	ProductQuantity uint `gorm:"not null"`
}