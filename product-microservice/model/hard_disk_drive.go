package model

type HardDiskDrive struct {
	Product      Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId    int     `gorm:"primaryKey;auto_increment"`
	Capacity     string  `gorm:"type:varchar(30);not null"`
	DiskSpeed    *string `gorm:"type:varchar(30)"`
	Interface    *string `gorm:"type:varchar(30)"`
	TransferRate *string `gorm:"type:varchar(30)"`
	Form         *string `gorm:"type:varchar(30)"`
}