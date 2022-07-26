package repository

import (
	"news/model"

	"gorm.io/gorm"
)

type newsSubscriptionRepository struct {
	Database *gorm.DB
}

type INewsSubscriptionRepository interface {
	Create(newsSubscription model.NewsSubscription) error
	Delete(newsSubscription model.NewsSubscription) error
	IsUserSubscribed(userId int) (model.NewsSubscription, error)
}

func NewNewsSubscriptionRepository(DB *gorm.DB) INewsSubscriptionRepository {
	return &newsSubscriptionRepository{Database: DB}
}

func (newsSubscriptionRepo *newsSubscriptionRepository) Create(newsSubscription model.NewsSubscription) error {
	result := newsSubscriptionRepo.Database.Create(&newsSubscription)
	return result.Error
}

func (newsSubscriptionRepo *newsSubscriptionRepository) Delete(newsSubscription model.NewsSubscription) error {
	result := newsSubscriptionRepo.Database.Delete(&newsSubscription)
	return result.Error
}

func (newsSubscriptionRepo *newsSubscriptionRepository) IsUserSubscribed(userId int) (model.NewsSubscription, error) {
	var newsSubscription model.NewsSubscription
	result := newsSubscriptionRepo.Database.
		Where("user_id = ?", userId).
		First(&newsSubscription)
	return newsSubscription, result.Error
}