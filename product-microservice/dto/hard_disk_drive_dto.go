package dto

type HardDiskDriveDTO struct {
	Product      ProductDTO
	Capacity     string
	DiskSpeed    *string
	Interface    *string
	TransferRate *string
	Form         *string
}