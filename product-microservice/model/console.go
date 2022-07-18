package model

type Console struct {
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId int     `gorm:"primaryKey;auto_increment"`
	Platform  string  `gorm:"type:varchar(40);not null"`
}