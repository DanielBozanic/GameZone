package dto

type HeadphonesDTO struct {
	Product           ProductDTO
	Sensitivity       *string
	ConnectionType    string
	Connection        string
	DriverSize        *string
	Microphone        *bool
	Color             *string
	Weight            *string
	FrequencyResponse *string
}