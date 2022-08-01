package model

import "time"

type Report struct {
	Id                int       `gorm:"primaryKey;auto_increment"`
	UserId            int       `gorm:"not null"`
	Reason            string    `gorm:"type:varchar(100);not null"`
	Description       string    `gorm:"type:varchar(1000);default:null"`
	DateTime          time.Time `gorm:"not null"`
}