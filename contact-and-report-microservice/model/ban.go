package model

import "time"

type Ban struct {
	Id             int       `gorm:"primaryKey;auto_increment"`
	UserId         int       `gorm:"not null"`
	Reason         Reason    `gorm:"not null"`
	ExpirationDate time.Time `gorm:"not null"`
}