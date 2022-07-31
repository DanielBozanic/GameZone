package model

import "time"

type ContactMessage struct {
	Id           int       `gorm:"primaryKey;auto_increment"`
	UserId       int       `gorm:"not null"`
	UserQuestion string    `gorm:"type:BLOB;not null"`
	Answer       string    `gorm:"type:BLOB;default:null"`
	DateTime     time.Time `gorm:"not null"`
}