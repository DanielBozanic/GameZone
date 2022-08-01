package repository

import (
	"contact-and-report/model"
	"time"

	"gorm.io/gorm"
)

type banRepository struct {
	Database *gorm.DB
}

type IBanRepository interface {
	GetAll() ([] model.Ban)
	GetUserBanHistory(userId int) []model.Ban
	IsUserBanned(userId int) (model.Ban, error)
	Create(ban model.Ban) error
}

func NewBanRepository(DB *gorm.DB) IBanRepository {
	return &banRepository{Database: DB}
}

func (banRepo *banRepository) GetAll() []model.Ban {
	var bans []model.Ban
	banRepo.Database.
		Find(&bans)
	return bans
}

func (banRepo *banRepository) GetUserBanHistory(userId int) []model.Ban {
	var bans []model.Ban
	banRepo.Database.
		Where("user_id = ?", userId).
		Order("expiration_date DESC").
		Find(&bans)
	return bans
}

func (banRepo *banRepository) IsUserBanned(userId int) (model.Ban, error) {
	var ban model.Ban
	currentDateTime := time.Now()
	result := banRepo.Database.
		Where("user_id = ? AND expiration_date > ?", userId, currentDateTime).
		First(&ban)
	return ban, result.Error
}

func (banRepo *banRepository) Create(ban model.Ban) error {
	result := banRepo.Database.Create(&ban)
	return result.Error
}