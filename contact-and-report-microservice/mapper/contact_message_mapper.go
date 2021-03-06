package mapper

import (
	"contact-and-report/dto"
	"contact-and-report/model"
)


func ToContactMessage(contactMessageDTO dto.ContactMessageDTO) (model.ContactMessage) {
	return model.ContactMessage {
		Id: contactMessageDTO.Id,
		UserId: contactMessageDTO.UserId,
		UserQuestion: contactMessageDTO.UserQuestion,
		Answer: contactMessageDTO.Answer,
		DateTime: contactMessageDTO.DateTime,
	}
}

func ToContactMessageDTO(contactMessage model.ContactMessage) dto.ContactMessageDTO {
	return dto.ContactMessageDTO {
		Id: contactMessage.Id,
		UserId: contactMessage.UserId,
		UserQuestion: contactMessage.UserQuestion,
		Answer: contactMessage.Answer,
		DateTime: contactMessage.DateTime,
	}
}

func ToContactMessageDTOs(contactMessages []model.ContactMessage) []dto.ContactMessageDTO {
	contactMessageDTOs := make([]dto.ContactMessageDTO, len(contactMessages))

	for i, itm := range contactMessages {
		contactMessageDTOs[i] = ToContactMessageDTO(itm)
	}

	return contactMessageDTOs
}