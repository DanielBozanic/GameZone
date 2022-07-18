package model

type File struct {
	Id      int    `gorm:"primaryKey;auto_increment"`
	Name    string `gorm:"type:varchar(200);not null"`
	Type    string `gorm:"type:varchar(50);not null"`
	Size    int    `gorm:"not null"`
	Content string `gorm:"type:LONGBLOB;not null"`
}