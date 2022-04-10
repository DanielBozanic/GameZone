package mapper

import (
	"product/dto"
	"product/model"
)


func ToPowerSupplyUnit(psuDTO dto.PowerSupplyUnitDTO) (model.PowerSupplyUnit) {
	return model.PowerSupplyUnit {
		Name: psuDTO.Name, 
		Price: psuDTO.Price, 
		PowerRating: psuDTO.PowerRating,
		Type: psuDTO.Type,
		FormFactor: psuDTO.FormFactor,
		Amount: psuDTO.Amount,
		Manufacturer: psuDTO.Manufacturer,
	}
}

func ToPowerSupplyUnitDTO(psu model.PowerSupplyUnit) dto.PowerSupplyUnitDTO {
	return dto.PowerSupplyUnitDTO {
		Id: psu.Id, 
		Name: psu.Name, 
		Price: psu.Price, 
		PowerRating: psu.PowerRating,
		Type: psu.Type,
		FormFactor: psu.FormFactor,
		Amount: psu.Amount,
		Manufacturer: psu.Manufacturer,
	}
}

func ToPowerSupplyUnitDTOs(psus []model.PowerSupplyUnit) []dto.PowerSupplyUnitDTO {
	psuDTOs := make([]dto.PowerSupplyUnitDTO, len(psus))

	for i, itm := range psus {
		psuDTOs[i] = ToPowerSupplyUnitDTO(itm)
	}

	return psuDTOs
}