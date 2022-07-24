package mapper

import (
	"news/dto"
	"news/model"
)


func ToNewsSubscription(newsSubscriptionDTO dto.NewsSubscriptionDTO) (model.NewsSubscription) {
	return model.NewsSubscription {
		Id: newsSubscriptionDTO.Id,
		Email: newsSubscriptionDTO.Email,
	}
}

func ToNewsSubscriptionDTO(newsSubscription model.NewsSubscription) dto.NewsSubscriptionDTO {
	return dto.NewsSubscriptionDTO {
		Id: newsSubscription.Id,
		Email: newsSubscription.Email,
	}
}

func ToNewsSubscriptionDTOs(newsSubscriptions []model.NewsSubscription) []dto.NewsSubscriptionDTO {
	newsSubscriptionsDTOs := make([]dto.NewsSubscriptionDTO, len(newsSubscriptions))

	for i, itm := range newsSubscriptions {
		newsSubscriptionsDTOs[i] = ToNewsSubscriptionDTO(itm)
	}

	return newsSubscriptionsDTOs
}