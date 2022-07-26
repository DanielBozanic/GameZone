package mapper

import (
	"news/dto"
	"news/model"
)


func ToNewsSubscription(newsSubscriptionDTO dto.NewsSubscriptionDTO) (model.NewsSubscription) {
	return model.NewsSubscription {
		Id: newsSubscriptionDTO.Id,
		UserId: newsSubscriptionDTO.UserId,
	}
}

func ToNewsSubscriptionDTO(newsSubscription model.NewsSubscription) dto.NewsSubscriptionDTO {
	return dto.NewsSubscriptionDTO {
		Id: newsSubscription.Id,
		UserId: newsSubscription.UserId,
	}
}

func ToNewsSubscriptionDTOs(newsSubscriptions []model.NewsSubscription) []dto.NewsSubscriptionDTO {
	newsSubscriptionsDTOs := make([]dto.NewsSubscriptionDTO, len(newsSubscriptions))

	for i, itm := range newsSubscriptions {
		newsSubscriptionsDTOs[i] = ToNewsSubscriptionDTO(itm)
	}

	return newsSubscriptionsDTOs
}