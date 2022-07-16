package mapper

import (
	"product/dto"
	"product/model"
)


func ToHeadphones(headphonesDTO dto.HeadphonesDTO) (model.Headphones) {
	return model.Headphones {
		Product: model.Product(headphonesDTO.Product),
		Wireless: headphonesDTO.Wireless,
		VirtualSurroundEncoding: headphonesDTO.VirtualSurroundEncoding,
		Sensitivity: headphonesDTO.Sensitivity,
		ConnectionType: headphonesDTO.ConnectionType,
		DriverSize: headphonesDTO.DriverSize,
		Microphone: headphonesDTO.Microphone,
		Color: headphonesDTO.Color,
		Weight: headphonesDTO.Weight,
		FrequencyResponse: headphonesDTO.FrequencyResponse,
	}
}

func ToHeadphonesDTO(headphone model.Headphones) dto.HeadphonesDTO {
	return dto.HeadphonesDTO {
		Product: dto.ProductDTO(headphone.Product),
		Wireless: headphone.Wireless,
		VirtualSurroundEncoding: headphone.VirtualSurroundEncoding,
		Sensitivity: headphone.Sensitivity,
		ConnectionType: headphone.ConnectionType,
		DriverSize: headphone.DriverSize,
		Microphone: headphone.Microphone,
		Color: headphone.Color,
		Weight: headphone.Weight,
		FrequencyResponse: headphone.FrequencyResponse,
	}
}

func ToHeadphonesDTOs(headphones []model.Headphones) []dto.HeadphonesDTO {
	headphonesDTOs := make([]dto.HeadphonesDTO, len(headphones))

	for i, itm := range headphones {
		headphonesDTOs[i] = ToHeadphonesDTO(itm)
	}

	return headphonesDTOs
}