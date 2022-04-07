package mapper

import (
	"product/dto"
	"product/model"
)


func ToConsole(consoleDTO dto.ConsoleDTO) (model.Console) {
	return model.Console {
		Name: consoleDTO.Name, 
		Price: consoleDTO.Price, 
		Platform: consoleDTO.Platform,
		Amount: consoleDTO.Amount,
		Manufacturer: consoleDTO.Manufacturer,
	}
}

func ToConsoleDTO(console model.Console) dto.ConsoleDTO {
	return dto.ConsoleDTO {
		Id: console.Id, 
		Name: console.Name,
		Price: console.Price, 
		Platform: console.Platform,
		Manufacturer: console.Manufacturer,
		Amount: console.Amount,
	}
}

func ToConsoleDTOs(consoles []model.Console) []dto.ConsoleDTO {
	consoleDTOs := make([]dto.ConsoleDTO, len(consoles))

	for i, itm := range consoles {
		consoleDTOs[i] = ToConsoleDTO(itm)
	}

	return consoleDTOs
}