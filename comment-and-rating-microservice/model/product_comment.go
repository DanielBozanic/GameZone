package model

import "time"

type ProductComment struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	ProductId   int       `gorm:"not null"`
	UserId      int       `gorm:"not null"`
	Username    string    `gorm:"type:varchar(120);not null"`
	Comment     string    `gorm:"type:varchar(490);not null"`
	Rating      Rating    `gorm:"not null"`
	Archived    *bool     `gorm:"type:boolean;default:false;not null"`
	DateTime    time.Time `gorm:"not null"`
}