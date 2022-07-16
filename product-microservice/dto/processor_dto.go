package dto

type ProcessorDTO struct {
	Product            ProductDTO
	Type               string
	Socket             string
	NumberOfCores      *uint
	Threads            *uint
	TDP                *string
	IntegratedGraphics *string
	BaseClockRate      *string
	TurboClockRate     *string
}