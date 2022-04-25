package mapper

import (
	"product/dto"
	"product/model"
)

func ToRam(ramDTO dto.RamDTO) (model.Ram) {
	return model.Ram {
		Product: model.Product(ramDTO.Product),
		MemoryType: ramDTO.MemoryType,
		Capacity: ramDTO.Capacity,
		Speed: ramDTO.Speed,
		Voltage: ramDTO.Voltage,
		Latency: ramDTO.Latency,
	}
}

func ToRamDTO(ram model.Ram) dto.RamDTO {
	return dto.RamDTO {
		Product: dto.ProductDTO(ram.Product),
		MemoryType: ram.MemoryType,
		Capacity: ram.Capacity,
		Speed: ram.Speed,
		Voltage: ram.Voltage,
		Latency: ram.Latency,
	}
}

func ToRamDTOs(processors []model.Ram) []dto.RamDTO {
	ramDTOs := make([]dto.RamDTO, len(processors))

	for i, itm := range processors {
		ramDTOs[i] = ToRamDTO(itm)
	}

	return ramDTOs
}