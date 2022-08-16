package model

import "time"

type ContactMessage struct {
	Id           int       `gorm:"primaryKey;auto_increment"`
	UserId       int       `gorm:"not null"`
	Username     string    `gorm:"type:varchar(120);not null"`
	Subject      string    `gorm:"type:varchar(70);not null"`
	Message      string    `gorm:"type:varchar(1000);not null"`
	Answer       string    `gorm:"type:varchar(1000);default:null"`
	DateTime     time.Time `gorm:"not null"`
}