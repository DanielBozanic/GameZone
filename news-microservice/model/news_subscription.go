package model

type NewsSubscription struct {
	Id    int    `gorm:"primaryKey;auto_increment"`
	Email string `gorm:"type:varchar(120);unique;not null"`
}