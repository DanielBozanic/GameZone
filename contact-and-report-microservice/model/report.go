package model

import "time"

type Report struct {
	Id                int       `gorm:"primaryKey;auto_increment"`
	UserId    int       `gorm:"not null"`
	Reason            Reason    `gorm:"not null"`
	ReasonDescription string    `gorm:"type:BLOB;default:null"`
	DateTime          time.Time `gorm:"not null"`
}