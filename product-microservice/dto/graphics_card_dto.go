package dto

type GraphicsCardDTO struct {
	Product          ProductDTO
	ChipManufacturer string
	ModelName        string
	BusWidth         *string
	MemorySize       *string
	MemoryType       *string
	PCIInterface     *string
	GPUSpeed         *string
	StreamProcessors *uint
	Cooling          *string
	HDMI             *uint
	DisplayPort      *uint
	TDP              *string
	PowerConnector   *string
	Dimensions       *string
}