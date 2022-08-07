package mapper

import (
	"product/dto"
	"product/model"
)


func ToGraphicsCard(graphicsCardDTO dto.GraphicsCardDTO) (model.GraphicsCard) {
	return model.GraphicsCard {
		Product: ToProduct(graphicsCardDTO.Product),
		ChipManufacturer: graphicsCardDTO.ChipManufacturer,
		ModelName: graphicsCardDTO.ModelName,
		BusWidth: graphicsCardDTO.BusWidth,
		MemorySize: graphicsCardDTO.MemorySize,
		MemoryType: graphicsCardDTO.MemoryType,
		PCIInterface: graphicsCardDTO.PCIInterface,
		GPUSpeed: graphicsCardDTO.GPUSpeed,
		StreamProcessors: graphicsCardDTO.StreamProcessors,
		Cooling: graphicsCardDTO.Cooling,
		HDMI: graphicsCardDTO.HDMI,
		DisplayPort: graphicsCardDTO.DisplayPort,
		TDP: graphicsCardDTO.TDP,
		PowerConnector: graphicsCardDTO.PowerConnector,
		Dimensions: graphicsCardDTO.Dimensions,
	}
}

func ToGraphicsCardDTO(graphicsCard model.GraphicsCard) dto.GraphicsCardDTO {
	return dto.GraphicsCardDTO {
		Product: ToProductDTO(graphicsCard.Product),
		ChipManufacturer: graphicsCard.ChipManufacturer,
		ModelName: graphicsCard.ModelName,
		BusWidth: graphicsCard.BusWidth,
		MemorySize: graphicsCard.MemorySize,
		MemoryType: graphicsCard.MemoryType,
		PCIInterface: graphicsCard.PCIInterface,
		GPUSpeed: graphicsCard.GPUSpeed,
		StreamProcessors: graphicsCard.StreamProcessors,
		Cooling: graphicsCard.Cooling,
		HDMI: graphicsCard.HDMI,
		DisplayPort: graphicsCard.DisplayPort,
		TDP: graphicsCard.TDP,
		PowerConnector: graphicsCard.PowerConnector,
		Dimensions: graphicsCard.Dimensions,
	}
}

func ToGraphicsCardDTOs(graphicsCards []model.GraphicsCard) []dto.GraphicsCardDTO {
	graphicsCardDTOs := make([]dto.GraphicsCardDTO, len(graphicsCards))

	for i, itm := range graphicsCards {
		graphicsCardDTOs[i] = ToGraphicsCardDTO(itm)
	}

	return graphicsCardDTOs
}