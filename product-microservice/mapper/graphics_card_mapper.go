package mapper

import (
	"product/dto"
	"product/model"
)


func ToGraphicsCard(graphicsCardDTO dto.GraphicsCardDTO) (model.GraphicsCard) {
	return model.GraphicsCard {
		Product: model.Product(graphicsCardDTO.Product),
		ModelName: graphicsCardDTO.ModelName,
		BusWidth: graphicsCardDTO.BusWidth,
		MemorySize: graphicsCardDTO.MemorySize,
		MemoryType: graphicsCardDTO.MemoryType,
		PCIInterface: graphicsCardDTO.PCIInterface,
		GPUSpeed: graphicsCardDTO.GPUSpeed,
		CUDAStreamProcessors: graphicsCardDTO.CUDAStreamProcessors,
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
		Product: dto.ProductDTO(graphicsCard.Product),
		ModelName: graphicsCard.ModelName,
		BusWidth: graphicsCard.BusWidth,
		MemorySize: graphicsCard.MemorySize,
		MemoryType: graphicsCard.MemoryType,
		PCIInterface: graphicsCard.PCIInterface,
		GPUSpeed: graphicsCard.GPUSpeed,
		CUDAStreamProcessors: graphicsCard.CUDAStreamProcessors,
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