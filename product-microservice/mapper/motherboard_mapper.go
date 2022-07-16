package mapper

import (
	"product/dto"
	"product/model"
)


func ToMotherboard(motherboardDTO dto.MotherboardDTO) (model.Motherboard) {
	return model.Motherboard {
		Product: model.Product(motherboardDTO.Product),
		ProcessorType: motherboardDTO.ProcessorType,
		Socket: motherboardDTO.Socket,
		SupportedProcessors: motherboardDTO.SupportedProcessors,
		Chipset: motherboardDTO.Chipset,
		Memory: motherboardDTO.Memory,
		ExpansionSlots: motherboardDTO.ExpansionSlots,
		StorageInterface: motherboardDTO.StorageInterface,
		Audio: motherboardDTO.Audio,
		USB: motherboardDTO.USB,
		BackPanelConnectors: motherboardDTO.BackPanelConnectors,
		InternalConnectors: motherboardDTO.InternalConnectors,
		BIOS: motherboardDTO.BIOS,
		FormFactor: motherboardDTO.FormFactor,
	}
}

func ToMotherboardDTO(motherboard model.Motherboard) dto.MotherboardDTO {
	return dto.MotherboardDTO {
		Product: dto.ProductDTO(motherboard.Product),
		ProcessorType: motherboard.ProcessorType,
		Socket: motherboard.Socket,
		SupportedProcessors: motherboard.SupportedProcessors,
		Chipset: motherboard.Chipset,
		Memory: motherboard.Memory,
		ExpansionSlots: motherboard.ExpansionSlots,
		StorageInterface: motherboard.StorageInterface,
		Audio: motherboard.Audio,
		USB: motherboard.USB,
		BackPanelConnectors: motherboard.BackPanelConnectors,
		InternalConnectors: motherboard.InternalConnectors,
		BIOS: motherboard.BIOS,
		FormFactor: motherboard.FormFactor,
	}
}

func ToMotherboardDTOs(motherboards []model.Motherboard) []dto.MotherboardDTO {
	motherboardDTOs := make([]dto.MotherboardDTO, len(motherboards))

	for i, itm := range motherboards {
		motherboardDTOs[i] = ToMotherboardDTO(itm)
	}

	return motherboardDTOs
}