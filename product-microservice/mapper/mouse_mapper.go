package mapper

import (
	"product/dto"
	"product/model"
)


func ToMouse(mouseDTO dto.MouseDTO) (model.Mouse) {
	return model.Mouse {
		Name: mouseDTO.Name, 
		Price: mouseDTO.Price,
		Wireless: mouseDTO.Wireless,
		Sensor: mouseDTO.Sensor,
		DPI: mouseDTO.DPI,
		PollingRate: mouseDTO.PollingRate,
		Connection: mouseDTO.Connection,
		Color: mouseDTO.Color,
		Acceleration: mouseDTO.Acceleration,
		Buttons: mouseDTO.Buttons,
		Weight: mouseDTO.Weight,
		Lifespan: mouseDTO.Lifespan,
		TrackingSpeed: mouseDTO.TrackingSpeed,
		Amount: mouseDTO.Amount,
		Manufacturer: mouseDTO.Manufacturer,
	}
}

func ToMouseDTO(mouse model.Mouse) dto.MouseDTO {
	return dto.MouseDTO {
		Id: mouse.Id, 
		Name: mouse.Name, 
		Price: mouse.Price,
		Wireless: mouse.Wireless,
		Sensor: mouse.Sensor,
		DPI: mouse.DPI,
		PollingRate: mouse.PollingRate,
		Connection: mouse.Connection,
		Color: mouse.Color,
		Acceleration: mouse.Acceleration,
		Buttons: mouse.Buttons,
		Weight: mouse.Weight,
		Lifespan: mouse.Lifespan,
		TrackingSpeed: mouse.TrackingSpeed,
		Amount: mouse.Amount,
		Manufacturer: mouse.Manufacturer,
	}
}

func ToMouseDTOs(mouses []model.Mouse) []dto.MouseDTO {
	mouseDTOs := make([]dto.MouseDTO, len(mouses))

	for i, itm := range mouses {
		mouseDTOs[i] = ToMouseDTO(itm)
	}

	return mouseDTOs
}