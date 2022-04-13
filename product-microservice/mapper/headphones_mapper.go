package mapper

import (
	"product/dto"
	"product/model"
)


func ToHeadphones(headphonesDTO dto.HeadphonesDTO) (model.Headphones) {
	return model.Headphones {
		Name: headphonesDTO.Name, 
		Price: headphonesDTO.Price,
		Wireless: headphonesDTO.Wireless,
		Description: headphonesDTO.Description,
		VirtualSurroundEncoding: headphonesDTO.VirtualSurroundEncoding,
		Sensitivity: headphonesDTO.Sensitivity,
		ConnectionType: headphonesDTO.ConnectionType,
		DriverSize: headphonesDTO.DriverSize,
		Microphone: headphonesDTO.Microphone,
		Color: headphonesDTO.Color,
		Weight: headphonesDTO.Weight,
		FrequencyResponse: headphonesDTO.FrequencyResponse,
		Amount: headphonesDTO.Amount,
		Manufacturer: headphonesDTO.Manufacturer,
	}
}

func ToHeadphonesDTO(headphone model.Headphones) dto.HeadphonesDTO {
	return dto.HeadphonesDTO {
		Id: headphone.Id, 
		Name: headphone.Name, 
		Price: headphone.Price,
		Wireless: headphone.Wireless,
		Description: headphone.Description,
		VirtualSurroundEncoding: headphone.VirtualSurroundEncoding,
		Sensitivity: headphone.Sensitivity,
		ConnectionType: headphone.ConnectionType,
		DriverSize: headphone.DriverSize,
		Microphone: headphone.Microphone,
		Color: headphone.Color,
		Weight: headphone.Weight,
		FrequencyResponse: headphone.FrequencyResponse,
		Amount: headphone.Amount,
		Manufacturer: headphone.Manufacturer,
	}
}

func ToHeadphonesDTOs(headphones []model.Headphones) []dto.HeadphonesDTO {
	headphonesDTOs := make([]dto.HeadphonesDTO, len(headphones))

	for i, itm := range headphones {
		headphonesDTOs[i] = ToHeadphonesDTO(itm)
	}

	return headphonesDTOs
}