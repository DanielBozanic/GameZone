package model

import "time"

type NewsArticle struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	UnpublishedTitle       string    `gorm:"type:varchar(250);not null"`
	UnpublishedDescription *string   `gorm:"type:varchar(500);default:null"`
	UnpublishedContent     string    `gorm:"type:LONGBLOB;not null"`
	PublishedTitle       string    `gorm:"type:varchar(250);default:null"`
	PublishedDescription *string   `gorm:"type:varchar(500);default:null"`
	PublishedContent      string    `gorm:"type:LONGBLOB;default:null"`
	DateTime    time.Time `gorm:"not null"`
	IsSent      *bool     `gorm:"type:boolean;default:false;not null"`
	Archived    *bool     `gorm:"type:boolean;default:false;not null"`
}