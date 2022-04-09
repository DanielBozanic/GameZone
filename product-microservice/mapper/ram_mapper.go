package mapper

import (
	"product/dto"
	"product/model"
)

func ToRam(ramDTO dto.RamDTO) (model.Ram) {
	return model.Ram {
		Name: ramDTO.Name,
		MemoryType: ramDTO.MemoryType,
		Manufacturer: ramDTO.Manufacturer,
		Capacity: ramDTO.Capacity,
		Speed: ramDTO.Speed,
		Voltage: ramDTO.Voltage,
		Latency: ramDTO.Latency,
		Price: ramDTO.Price,
		Amount: ramDTO.Amount,
	}
}

func ToRamDTO(ram model.Ram) dto.RamDTO {
	return dto.RamDTO {
		Id: ram.Id, 
		Name: ram.Name,
		MemoryType: ram.MemoryType,
		Manufacturer: ram.Manufacturer,
		Capacity: ram.Capacity,
		Speed: ram.Speed,
		Voltage: ram.Voltage,
		Latency: ram.Latency,
		Price: ram.Price,
		Amount: ram.Amount,
	}
}

func ToRamDTOs(processors []model.Ram) []dto.RamDTO {
	ramDTOs := make([]dto.RamDTO, len(processors))

	for i, itm := range processors {
		ramDTOs[i] = ToRamDTO(itm)
	}

	return ramDTOs
}