package mapper

import (
	"product/dto"
	"product/model"
)


func ToPowerSupplyUnit(psuDTO dto.PowerSupplyUnitDTO) (model.PowerSupplyUnit) {
	return model.PowerSupplyUnit {
		Product: ToProduct(psuDTO.Product),
		Power: psuDTO.Power,
		Type: psuDTO.Type,
		FormFactor: psuDTO.FormFactor,
	}
}

func ToPowerSupplyUnitDTO(psu model.PowerSupplyUnit) dto.PowerSupplyUnitDTO {
	return dto.PowerSupplyUnitDTO {
		Product: ToProductDTO(psu.Product),
		Power: psu.Power,
		Type: psu.Type,
		FormFactor: psu.FormFactor,
	}
}

func ToPowerSupplyUnitDTOs(psus []model.PowerSupplyUnit) []dto.PowerSupplyUnitDTO {
	psuDTOs := make([]dto.PowerSupplyUnitDTO, len(psus))

	for i, itm := range psus {
		psuDTOs[i] = ToPowerSupplyUnitDTO(itm)
	}

	return psuDTOs
}