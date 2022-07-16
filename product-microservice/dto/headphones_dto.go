package dto

type HeadphonesDTO struct {
	Product                 ProductDTO
	VirtualSurroundEncoding *string
	Sensitivity             *string
	ConnectionType          string
	Wireless                bool
	DriverSize              *string
	Microphone              bool
	Color                   *string
	Weight                  *string
	FrequencyResponse       *string
}