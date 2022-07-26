package model

import "time"

type NewsComment struct {
	Id            int `gorm:"primaryKey;auto_increment"`
	NewsArticleId int
	NewsArticle   NewsArticle `gorm:"foreignKey:NewsArticleId"`
	Comment       string      `gorm:"type:varchar(490);not null"`
	UserId        int         `gorm:"not null"`
	DateTime      time.Time   `gorm:"not null"`
	Archived      *bool       `gorm:"type:boolean;default:false;not null"`
}