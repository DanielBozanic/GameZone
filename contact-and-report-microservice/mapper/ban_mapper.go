package mapper

import (
	"contact-and-report/dto"
	"contact-and-report/model"
)


func ToBan(banDTO dto.BanDTO) (model.Ban) {
	return model.Ban {
		Id: banDTO.Id,
		UserId: banDTO.UserId,
		Reason: banDTO.Reason,
		Description: banDTO.Description,
		ExpirationDate: banDTO.ExpirationDate,
	}
}

func ToBanDTO(ban model.Ban) dto.BanDTO {
	return dto.BanDTO {
		Id: ban.Id,
		UserId: ban.UserId,
		Reason: ban.Reason,
		Description: ban.Description,
		ExpirationDate: ban.ExpirationDate,
	}
}

func ToBanDTOs(bans []model.Ban) []dto.BanDTO {
	banDTOs := make([]dto.BanDTO, len(bans))

	for i, itm := range bans {
		banDTOs[i] = ToBanDTO(itm)
	}

	return banDTOs
}