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
	Subscribe(email string) string
	Unsubscribe(email string) string
	IsUserSubscribed(email string) bool
}

func NewNewsSubscriptionService(newsSubscriptionRepository repository.INewsSubscriptionRepository) INewsSubscriptionService {
	return &newsSubscriptionService{INewsSubscriptionRepository: newsSubscriptionRepository}
}

func (newsSubscriptionService *newsSubscriptionService) Subscribe(email string) string {
	msg := ""
	var newsSubscription model.NewsSubscription
	_, err := newsSubscriptionService.INewsSubscriptionRepository.IsUserSubscribed(email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "You are already subscribed to receive emails."
	}
	newsSubscription.Email = email
	err = newsSubscriptionService.INewsSubscriptionRepository.Create(newsSubscription)
	if err != nil {
		msg = err.Error()
	}
	return msg
}

func (newsSubscriptionService *newsSubscriptionService) Unsubscribe(email string) string {
	msg := ""
	newsSubscription, err := newsSubscriptionService.INewsSubscriptionRepository.IsUserSubscribed(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "You are not subscribed to receive emails."
	}
	
	err = newsSubscriptionService.INewsSubscriptionRepository.Delete(newsSubscription)
	if err != nil {
		msg = err.Error()
	}
	return msg
}

func (newsSubscriptionService *newsSubscriptionService) IsUserSubscribed(email string) bool {
	_, err := newsSubscriptionService.INewsSubscriptionRepository.IsUserSubscribed(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	} else {
		return true
	}
}