package dto

type MotherboardDTO struct {
	Product             ProductDTO
	ProcessorType       string
	Socket              string
	SupportedProcessors *string
	Chipset             *string
	Memory              *string
	ExpansionSlots      *string
	StorageInterface    *string
	Audio               *string
	USB                 *string
	BackPanelConnectors *string
	InternalConnectors  *string
	BIOS                *string
	FormFactor          *string
}