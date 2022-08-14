package model

type Motherboard struct {
	Product             Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId           int     `gorm:"primaryKey;auto_increment"`
	ProcessorType       string  `gorm:"type:varchar(30);not null"`
	Socket              string  `gorm:"type:varchar(30);not null"`
	SupportedProcessors *string `gorm:"type:varchar(1000)"`
	Chipset             *string `gorm:"type:varchar(30)"`
	Memory              *string `gorm:"type:varchar(1000)"`
	ExpansionSlots      *string `gorm:"type:varchar(1000)"`
	StorageInterface    *string `gorm:"type:varchar(1000)"`
	Audio               *string `gorm:"type:varchar(1000)"`
	USB                 *string `gorm:"type:varchar(1000)"`
	BackPanelConnectors *string `gorm:"type:varchar(1000)"`
	InternalConnectors  *string `gorm:"type:varchar(1000)"`
	BIOS                *string `gorm:"type:varchar(1000)"`
	FormFactor          *string `gorm:"type:varchar(40)"`
}