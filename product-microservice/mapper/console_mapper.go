package mapper

import (
	"product/dto"
	"product/model"
)


func ToConsole(consoleDTO dto.ConsoleDTO) (model.Console) {
	return model.Console {
		Product: model.Product(consoleDTO.Product),
		Platform: consoleDTO.Platform,
	}
}

func ToConsoleDTO(console model.Console) dto.ConsoleDTO {
	return dto.ConsoleDTO {
		Product: dto.ProductDTO(console.Product),
		Platform: console.Platform,
	}
}

func ToConsoleDTOs(consoles []model.Console) []dto.ConsoleDTO {
	consoleDTOs := make([]dto.ConsoleDTO, len(consoles))

	for i, itm := range consoles {
		consoleDTOs[i] = ToConsoleDTO(itm)
	}

	return consoleDTOs
}