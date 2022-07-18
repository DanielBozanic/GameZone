package mapper

import (
	"product/dto"
	"product/model"
)

func ToSolidStateDrive(ssdDTO dto.SolidStateDriveDTO) (model.SolidStateDrive) {
	return model.SolidStateDrive {
		Product: ToProduct(ssdDTO.Product),
		Interface: ssdDTO.Interface,
		MaxSequentialRead: ssdDTO.MaxSequentialRead,
		MaxSequentialWrite: ssdDTO.MaxSequentialWrite,
		Form: ssdDTO.Form,
		Dimensions: ssdDTO.Dimensions,
		Capacity: ssdDTO.Capacity,
	}
}

func ToSolidStateDriveDTO(ssd model.SolidStateDrive) dto.SolidStateDriveDTO {
	return dto.SolidStateDriveDTO {
		Product: ToProductDTO(ssd.Product),
		Interface: ssd.Interface,
		MaxSequentialRead: ssd.MaxSequentialRead,
		MaxSequentialWrite: ssd.MaxSequentialWrite,
		Form: ssd.Form,
		Dimensions: ssd.Dimensions,
		Capacity: ssd.Capacity,
	}
}

func ToSolidStateDriveDTOs(ssds []model.SolidStateDrive) []dto.SolidStateDriveDTO {
	ssdDTOs := make([]dto.SolidStateDriveDTO, len(ssds))

	for i, itm := range ssds {
		ssdDTOs[i] = ToSolidStateDriveDTO(itm)
	}

	return ssdDTOs
}