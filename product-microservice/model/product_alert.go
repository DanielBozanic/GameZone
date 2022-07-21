package model

type ProductAlert struct {
	Id        int     `gorm:"primaryKey;auto_increment"`
	ProductId int     `gorm:"index:idx_name,unique;not null"`
	Product   Product `gorm:"foreignKey:ProductId"`
	UserEmail string  `gorm:"type:varchar(120);index:idx_name,unique;not null"`
}