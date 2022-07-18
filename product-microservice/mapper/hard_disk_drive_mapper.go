package mapper

import (
	"product/dto"
	"product/model"
)

func ToHardDiskDrive(hddDTO dto.HardDiskDriveDTO) (model.HardDiskDrive) {
	return model.HardDiskDrive {
		Product: ToProduct(hddDTO.Product),
		Interface: hddDTO.Interface,
		Form: hddDTO.Form,
		DiskSpeed: hddDTO.DiskSpeed,
		TransferRate: hddDTO.TransferRate,
		Capacity: hddDTO.Capacity,
	}
}

func ToHardDiskDriveDTO(hdd model.HardDiskDrive) dto.HardDiskDriveDTO {
	return dto.HardDiskDriveDTO {
		Product: ToProductDTO(hdd.Product),
		Interface: hdd.Interface,
		Form: hdd.Form,
		DiskSpeed: hdd.DiskSpeed,
		TransferRate: hdd.TransferRate,
		Capacity: hdd.Capacity,
	}
}

func ToHardDiskDriveDTOs(hdds []model.HardDiskDrive) []dto.HardDiskDriveDTO {
	hddDTOs := make([]dto.HardDiskDriveDTO, len(hdds))

	for i, itm := range hdds {
		hddDTOs[i] = ToHardDiskDriveDTO(itm)
	}

	return hddDTOs
}