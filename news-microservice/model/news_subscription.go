package model

type NewsSubscription struct {
	Id     int `gorm:"primaryKey;auto_increment"`
	UserId int `gorm:"unique;not null"`
}