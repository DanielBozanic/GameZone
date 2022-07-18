package model

type SolidStateDrive struct {
	Product            Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId          int     `gorm:"primaryKey;auto_increment"`
	Capacity           string  `gorm:"type:varchar(30);not null"`
	Form               string  `gorm:"type:varchar(30);not null"`
	Interface          *string `gorm:"type:varchar(30)"`
	MaxSequentialRead  *string `gorm:"type:varchar(30)"`
	MaxSequentialWrite *string `gorm:"type:varchar(30)"`
	Dimensions         *string `gorm:"type:varchar(40)"`
}