package mapper

import (
	"product/dto"
	"product/model"
)

func ToHardDiskDrive(hddDTO dto.HardDiskDriveDTO) (model.HardDiskDrive) {
	return model.HardDiskDrive {
		Name: hddDTO.Name,
		Interface: hddDTO.Interface,
		Form: hddDTO.Form,
		DiskSpeed: hddDTO.DiskSpeed,
		TransferRate: hddDTO.TransferRate,
		Manufacturer: hddDTO.Manufacturer,
		Capacity: hddDTO.Capacity,
		Price: hddDTO.Price,
		Amount: hddDTO.Amount,
	}
}

func ToHardDiskDriveDTO(hdd model.HardDiskDrive) dto.HardDiskDriveDTO {
	return dto.HardDiskDriveDTO {
		Id: hdd.Id, 
		Name: hdd.Name,
		Interface: hdd.Interface,
		Form: hdd.Form,
		DiskSpeed: hdd.DiskSpeed,
		TransferRate: hdd.TransferRate,
		Manufacturer: hdd.Manufacturer,
		Capacity: hdd.Capacity,
		Price: hdd.Price,
		Amount: hdd.Amount,
	}
}

func ToHardDiskDriveDTOs(hdds []model.HardDiskDrive) []dto.HardDiskDriveDTO {
	hddDTOs := make([]dto.HardDiskDriveDTO, len(hdds))

	for i, itm := range hdds {
		hddDTOs[i] = ToHardDiskDriveDTO(itm)
	}

	return hddDTOs
}