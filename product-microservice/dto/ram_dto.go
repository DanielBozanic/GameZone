package dto

type RamDTO struct {
	Product    ProductDTO
	MemoryType string
	Capacity   string
	Speed      *string
	Voltage    *string
	Latency    *string
}