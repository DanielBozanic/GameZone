package model

type PowerSupplyUnit struct {
	Product    Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId  int     `gorm:"primaryKey;auto_increment"`
	Power      string  `gorm:"type:varchar(40);not null"`
	Type       string  `gorm:"type:varchar(40);not null"`
	FormFactor *string `gorm:"type:varchar(40)"`
}