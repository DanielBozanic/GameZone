package model

type Ram struct {
	Product    Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId  int     `gorm:"primaryKey;auto_increment"`
	MemoryType string  `gorm:"type:varchar(30);not null"`
	Capacity   string  `gorm:"type:varchar(30);not null"`
	Speed      *string `gorm:"type:varchar(30)"`
	Voltage    *string `gorm:"type:varchar(30)"`
	Latency    *string `gorm:"type:varchar(30)"`
}