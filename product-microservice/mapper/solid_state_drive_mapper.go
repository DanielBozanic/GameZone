package mapper

import (
	"product/dto"
	"product/model"
)

func ToSolidStateDrive(ssdDTO dto.SolidStateDriveDTO) (model.SolidStateDrive) {
	return model.SolidStateDrive {
		Name: ssdDTO.Name,
		Interface: ssdDTO.Interface,
		MaxSequentialRead: ssdDTO.MaxSequentialRead,
		MaxSequentialWrite: ssdDTO.MaxSequentialWrite,
		Form: ssdDTO.Form,
		Dimensions: ssdDTO.Dimensions,
		Manufacturer: ssdDTO.Manufacturer,
		Capacity: ssdDTO.Capacity,
		Price: ssdDTO.Price,
		Amount: ssdDTO.Amount,
	}
}

func ToSolidStateDriveDTO(ssd model.SolidStateDrive) dto.SolidStateDriveDTO {
	return dto.SolidStateDriveDTO {
		Id: ssd.Id, 
		Name: ssd.Name,
		Interface: ssd.Interface,
		MaxSequentialRead: ssd.MaxSequentialRead,
		MaxSequentialWrite: ssd.MaxSequentialWrite,
		Form: ssd.Form,
		Dimensions: ssd.Dimensions,
		Manufacturer: ssd.Manufacturer,
		Capacity: ssd.Capacity,
		Price: ssd.Price,
		Amount: ssd.Amount,
	}
}

func ToSolidStateDriveDTOs(ssds []model.SolidStateDrive) []dto.SolidStateDriveDTO {
	ssdDTOs := make([]dto.SolidStateDriveDTO, len(ssds))

	for i, itm := range ssds {
		ssdDTOs[i] = ToSolidStateDriveDTO(itm)
	}

	return ssdDTOs
}