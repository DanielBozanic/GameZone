package mapper

import (
	"product/dto"
	"product/model"
)


func ToMouse(mouseDTO dto.MouseDTO) (model.Mouse) {
	return model.Mouse {
		Product: ToProduct(mouseDTO.Product),
		Wireless: mouseDTO.Wireless,
		Sensor: mouseDTO.Sensor,
		DPI: mouseDTO.DPI,
		PollingRate: mouseDTO.PollingRate,
		Connection: mouseDTO.Connection,
		Color: mouseDTO.Color,
		Acceleration: mouseDTO.Acceleration,
		Weight: mouseDTO.Weight,
		Lifespan: mouseDTO.Lifespan,
		TrackingSpeed: mouseDTO.TrackingSpeed,
	}
}

func ToMouseDTO(mouse model.Mouse) dto.MouseDTO {
	return dto.MouseDTO {
		Product: ToProductDTO(mouse.Product),
		Wireless: mouse.Wireless,
		Sensor: mouse.Sensor,
		DPI: mouse.DPI,
		PollingRate: mouse.PollingRate,
		Connection: mouse.Connection,
		Color: mouse.Color,
		Acceleration: mouse.Acceleration,
		Weight: mouse.Weight,
		Lifespan: mouse.Lifespan,
		TrackingSpeed: mouse.TrackingSpeed,
	}
}

func ToMouseDTOs(mouses []model.Mouse) []dto.MouseDTO {
	mouseDTOs := make([]dto.MouseDTO, len(mouses))

	for i, itm := range mouses {
		mouseDTOs[i] = ToMouseDTO(itm)
	}

	return mouseDTOs
}