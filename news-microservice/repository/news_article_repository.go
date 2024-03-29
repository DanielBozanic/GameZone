package repository

import (
	"news/model"

	"gorm.io/gorm"
)

type newsArticleRepository struct {
	Database *gorm.DB
}

type INewsArticleRepository interface {
	GetAll(page int, pageSize int) []model.NewsArticle
	GetNumberOfRecords() int64
	GetPublishedArticles(page int, pageSize int) []model.NewsArticle
	GetNumberOfRecordsPublishedArticles() int64
	GetById(id int) (model.NewsArticle, error)
	GetUnsentPublishedArticles() []model.NewsArticle
	Create(newsArticle model.NewsArticle) model.NewsArticle
	Update(newsArticle model.NewsArticle) error
}

func NewNewsArticleRepository(DB *gorm.DB) INewsArticleRepository {
	return &newsArticleRepository{Database: DB}
}

func (newsArticleRepo *newsArticleRepository) GetAll(page int, pageSize int) []model.NewsArticle {
	var newsArticles []model.NewsArticle
	offset := (page - 1) * pageSize
	newsArticleRepo.Database.
		Offset(offset).Limit(pageSize).
		Where("archived = false").
		Order("date_time DESC").
		Find(&newsArticles)
	return newsArticles
}

func (newsArticleRepo *newsArticleRepository) GetNumberOfRecords() int64 {
	var count int64
	newsArticleRepo.Database.
		Where("archived = false").
		Model(&model.NewsArticle{}).
		Count(&count)
	return count
}

func (newsArticleRepo *newsArticleRepository) GetPublishedArticles(page int, pageSize int) []model.NewsArticle {
	var newsArticles []model.NewsArticle
	offset := (page - 1) * pageSize
	newsArticleRepo.Database.
		Offset(offset).Limit(pageSize).
		Where("archived = false AND published_content IS NOT NULL").
		Order("date_time DESC").
		Find(&newsArticles)
	return newsArticles
}

func (newsArticleRepo *newsArticleRepository) GetNumberOfRecordsPublishedArticles() int64 {
	var count int64
	newsArticleRepo.Database.
		Where("archived = false AND published_content IS NOT NULL").
		Model(&model.NewsArticle{}).
		Count(&count)
	return count
}

func (newsArticleRepo *newsArticleRepository) GetById(id int) (model.NewsArticle, error) {
	var newsArticle model.NewsArticle
	result := newsArticleRepo.Database.
		Where("archived = false").
		First(&newsArticle, id)
	return newsArticle, result.Error
}

func (newsArticleRepo *newsArticleRepository) GetUnsentPublishedArticles() []model.NewsArticle {
	var newsArticles []model.NewsArticle
	newsArticleRepo.Database.
		Where("archived = false AND published_title IS NOT NULL AND published_content IS NOT NULL AND is_sent = false").
		Find(&newsArticles)
	return newsArticles
}

func (newsArticleRepo *newsArticleRepository) Create(newsArticle model.NewsArticle) model.NewsArticle {
	newsArticleRepo.Database.Create(&newsArticle)
	return newsArticle
}

func (newsArticleRepo *newsArticleRepository) Update(newsArticle model.NewsArticle) error {
	result := newsArticleRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&newsArticle)
	return result.Error
}