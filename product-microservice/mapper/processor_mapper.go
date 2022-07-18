package mapper

import (
	"product/dto"
	"product/model"
)


func ToProcessor(processorDTO dto.ProcessorDTO) (model.Processor) {
	return model.Processor {
		Product: ToProduct(processorDTO.Product),
		Type: processorDTO.Type,
		Socket: processorDTO.Socket,
		NumberOfCores: processorDTO.NumberOfCores,
		Threads: processorDTO.Threads,
		IntegratedGraphics: processorDTO.IntegratedGraphics,
		BaseClockRate: processorDTO.BaseClockRate,
		TurboClockRate: processorDTO.TurboClockRate,
		TDP: processorDTO.TDP,
	}
}

func ToProcessorDTO(processor model.Processor) dto.ProcessorDTO {
	return dto.ProcessorDTO {
		Product: ToProductDTO(processor.Product),
		Type: processor.Type,
		Socket: processor.Socket,
		NumberOfCores: processor.NumberOfCores,
		Threads: processor.Threads,
		IntegratedGraphics: processor.IntegratedGraphics,
		BaseClockRate: processor.BaseClockRate,
		TurboClockRate: processor.TurboClockRate,
		TDP: processor.TDP,
	}
}

func ToProcessorDTOs(processors []model.Processor) []dto.ProcessorDTO {
	processorDTOs := make([]dto.ProcessorDTO, len(processors))

	for i, itm := range processors {
		processorDTOs[i] = ToProcessorDTO(itm)
	}

	return processorDTOs
}