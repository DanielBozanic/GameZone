package model

import "time"

type ProductComment struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	ProductName string    `gorm:"type:varchar(100);index:idx_name,unique;not null"`
	Username    string    `gorm:"type:varchar(120);index:idx_name,unique;not null"`
	Comment     string    `gorm:"type:varchar(490);not null"`
	Rating      Rating    `gorm:"not null"`
	Archived    *bool     `gorm:"type:boolean;index:idx_name,unique;default:false;not null"`
	DateTime    time.Time `gorm:"not null"`
}