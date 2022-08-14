package dto

type MouseDTO struct {
	Product       ProductDTO
	Wireless      *bool
	Sensor        *string
	DPI           *string
	PollingRate   *string
	Connection    string
	Color         *string
	TrackingSpeed *string
	Acceleration  *string
	Weight        *string
	Lifespan      *string
}