package mapper

import (
	"product/dto"
	"product/model"
)


func ToMotherboard(motherboardDTO dto.MotherboardDTO) (model.Motherboard) {
	return model.Motherboard {
		Name: motherboardDTO.Name,
		ProcessorType: motherboardDTO.ProcessorType,
		Socket: motherboardDTO.Socket,
		SupportedProcessors: motherboardDTO.SupportedProcessors,
		Chipset: motherboardDTO.Chipset,
		Memory: motherboardDTO.Memory,
		MultiGraphicsTechnology: motherboardDTO.MultiGraphicsTechnology,
		ExpansionSlots: motherboardDTO.ExpansionSlots,
		StorageInterface: motherboardDTO.StorageInterface,
		WirelessCommunicationModule: motherboardDTO.WirelessCommunicationModule,
		Audio: motherboardDTO.Audio,
		USB: motherboardDTO.USB,
		BackPanelConnectors: motherboardDTO.BackPanelConnectors,
		InternalConnectors: motherboardDTO.InternalConnectors,
		BIOS: motherboardDTO.BIOS,
		Manufacturer: motherboardDTO.Manufacturer,
		Price: motherboardDTO.Price,
		Amount: motherboardDTO.Amount,
	}
}

func ToMotherboardDTO(motherboard model.Motherboard) dto.MotherboardDTO {
	return dto.MotherboardDTO {
		Id: motherboard.Id, 
		Name: motherboard.Name,
		ProcessorType: motherboard.ProcessorType,
		Socket: motherboard.Socket,
		SupportedProcessors: motherboard.SupportedProcessors,
		Chipset: motherboard.Chipset,
		Memory: motherboard.Memory,
		MultiGraphicsTechnology: motherboard.MultiGraphicsTechnology,
		ExpansionSlots: motherboard.ExpansionSlots,
		StorageInterface: motherboard.StorageInterface,
		WirelessCommunicationModule: motherboard.WirelessCommunicationModule,
		Audio: motherboard.Audio,
		USB: motherboard.USB,
		BackPanelConnectors: motherboard.BackPanelConnectors,
		InternalConnectors: motherboard.InternalConnectors,
		BIOS: motherboard.BIOS,
		Manufacturer: motherboard.Manufacturer,
		Price: motherboard.Price,
		Amount: motherboard.Amount,
	}
}

func ToMotherboardDTOs(motherboards []model.Motherboard) []dto.MotherboardDTO {
	motherboardDTOs := make([]dto.MotherboardDTO, len(motherboards))

	for i, itm := range motherboards {
		motherboardDTOs[i] = ToMotherboardDTO(itm)
	}

	return motherboardDTOs
}