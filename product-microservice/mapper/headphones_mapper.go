package mapper

import (
	"product/dto"
	"product/model"
)


func ToHeadphones(headphonesDTO dto.HeadphonesDTO) (model.Headphones) {
	return model.Headphones {
		Product: ToProduct(headphonesDTO.Product),
		Sensitivity: headphonesDTO.Sensitivity,
		ConnectionType: headphonesDTO.ConnectionType,
		Connection: headphonesDTO.Connection,
		DriverSize: headphonesDTO.DriverSize,
		Microphone: headphonesDTO.Microphone,
		Color: headphonesDTO.Color,
		Weight: headphonesDTO.Weight,
		FrequencyResponse: headphonesDTO.FrequencyResponse,
	}
}

func ToHeadphonesDTO(headphone model.Headphones) dto.HeadphonesDTO {
	return dto.HeadphonesDTO {
		Product: ToProductDTO(headphone.Product),
		Sensitivity: headphone.Sensitivity,
		ConnectionType: headphone.ConnectionType,
		Connection: headphone.Connection,
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