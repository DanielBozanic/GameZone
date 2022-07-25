package model

import "time"

type NewsArticle struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	Title       string    `gorm:"type:varchar(250);not null"`
	Description string   `gorm:"type:varchar(200);default:null"`
	UnpublishedContent     string    `gorm:"type:BLOB;not null"`
	PublishedContent      string    `gorm:"type:BLOB;default:null"`
	DateTime    time.Time `gorm:"not null"`
	IsSent      *bool     `gorm:"type:boolean;default:false;not null"`
	Archived    *bool     `gorm:"type:boolean;default:false;not null"`
}