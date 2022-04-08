package mapper

import (
	"product/dto"
	"product/model"
)


func ToProcessor(processorDTO dto.ProcessorDTO) (model.Processor) {
	return model.Processor {
		Name: processorDTO.Name,
		Type: processorDTO.Type,
		Manufacturer: processorDTO.Manufacturer,
		Socket: processorDTO.Socket,
		NumberOfCores: processorDTO.NumberOfCores,
		Threads: processorDTO.Threads,
		IntegratedGraphics: processorDTO.IntegratedGraphics,
		BaseClockRate: processorDTO.BaseClockRate,
		TurboClockRate: processorDTO.TurboClockRate,
		TDP: processorDTO.TDP,
		Price: processorDTO.Price,
		Amount: processorDTO.Amount,
	}
}

func ToProcessorDTO(processor model.Processor) dto.ProcessorDTO {
	return dto.ProcessorDTO {
		Id: processor.Id, 
		Type: processor.Type,
		Manufacturer: processor.Manufacturer,
		Socket: processor.Socket,
		NumberOfCores: processor.NumberOfCores,
		Threads: processor.Threads,
		IntegratedGraphics: processor.IntegratedGraphics,
		BaseClockRate: processor.BaseClockRate,
		TurboClockRate: processor.TurboClockRate,
		TDP: processor.TDP,
		Price: processor.Price,
		Amount: processor.Amount,
	}
}

func ToProcessorDTOs(processors []model.Processor) []dto.ProcessorDTO {
	processorDTOs := make([]dto.ProcessorDTO, len(processors))

	for i, itm := range processors {
		processorDTOs[i] = ToProcessorDTO(itm)
	}

	return processorDTOs
}