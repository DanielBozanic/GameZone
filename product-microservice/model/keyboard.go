package model

type Keyboard struct {
	Product           Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId         int     `gorm:"primaryKey;auto_increment"`
	Wireless          *bool   `gorm:"type:bool;not null"`
	KeyboardConnector string  `gorm:"type:varchar(30);not null"`
	KeyType           string  `gorm:"type:varchar(30);not null"`
	LetterLayout      *string `gorm:"type:varchar(20)"`
	KeyboardColor     *string `gorm:"type:varchar(20)"`
}