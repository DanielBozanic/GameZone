package service

import (
	"errors"
	"news/model"
	"news/repository"

	"gorm.io/gorm"
)

type newsSubscriptionService struct {
	INewsSubscriptionRepository repository.INewsSubscriptionRepository
}

type INewsSubscriptionService interface {
	Subscribe(userId int) string
	Unsubscribe(userId int) string
	IsUserSubscribed(userId int) bool
}

func NewNewsSubscriptionService(newsSubscriptionRepository repository.INewsSubscriptionRepository) INewsSubscriptionService {
	return &newsSubscriptionService{INewsSubscriptionRepository: newsSubscriptionRepository}
}

func (newsSubscriptionService *newsSubscriptionService) Subscribe(userId int) string {
	msg := ""
	var newsSubscription model.NewsSubscription
	_, err := newsSubscriptionService.INewsSubscriptionRepository.IsUserSubscribed(userId)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "You are already subscribed to receive emails."
	}
	newsSubscription.UserId = userId
	err = newsSubscriptionService.INewsSubscriptionRepository.Create(newsSubscription)
	if err != nil {
		msg = err.Error()
	}
	return msg
}

func (newsSubscriptionService *newsSubscriptionService) Unsubscribe(userId int) string {
	msg := ""
	newsSubscription, err := newsSubscriptionService.INewsSubscriptionRepository.IsUserSubscribed(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "You are not subscribed to receive emails."
	}
	
	err = newsSubscriptionService.INewsSubscriptionRepository.Delete(newsSubscription)
	if err != nil {
		msg = err.Error()
	}
	return msg
}

func (newsSubscriptionService *newsSubscriptionService) IsUserSubscribed(userId int) bool {
	_, err := newsSubscriptionService.INewsSubscriptionRepository.IsUserSubscribed(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	} else {
		return true
	}
}