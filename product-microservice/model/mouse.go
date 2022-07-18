package model

type Mouse struct {
	Product       Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId     int     `gorm:"primaryKey;auto_increment"`
	Wireless      *bool   `gorm:"type:bool;not null"`
	Connection    string  `gorm:"type:varchar(30);not null"`
	Sensor        *string `gorm:"type:varchar(30)"`
	DPI           *string `gorm:"type:varchar(40)"`
	PollingRate   *string `gorm:"type:varchar(30)"`
	Color         *string `gorm:"type:varchar(20)"`
	TrackingSpeed *string `gorm:"type:varchar(30)"`
	Acceleration  *string `gorm:"type:varchar(30)"`
	Buttons       *uint
	Weight        *string `gorm:"type:varchar(20)"`
	Lifespan      *string `gorm:"type:varchar(30)"`
}