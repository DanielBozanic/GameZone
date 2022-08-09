package service

import (
	"errors"
	"news/dto"
	"news/mapper"
	"news/model"
	"news/repository"
	"time"

	"gorm.io/gorm"
)

type newsArticleService struct {
	INewsArticleRepository repository.INewsArticleRepository
}

type INewsArticleService interface {
	GetAll(page int, pageSize int) []model.NewsArticle
	GetNumberOfRecords() int64
	GetPublishedArticles(page int, pageSize int) []model.NewsArticle
	GetNumberOfRecordsPublishedArticles() int64
	GetById(id int) (model.NewsArticle, error)
	AddNewsArticle(newsArticle model.NewsArticle) model.NewsArticle
	EditNewsArticle(newsArticleDTO dto.NewsArticleDTO) string
	DeleteNewsArticle(id int) error
	PublishNewsArticle(newsArticleDTO dto.NewsArticleDTO) string
}

func NewNewsArticleService(newsArticleRepository repository.INewsArticleRepository) INewsArticleService {
	return &newsArticleService{INewsArticleRepository: newsArticleRepository}
}

func (newsArticleService *newsArticleService) GetAll(page int, pageSize int) []model.NewsArticle {
	return newsArticleService.INewsArticleRepository.GetAll(page, pageSize)
}

func (newsArticleService *newsArticleService) GetNumberOfRecords() int64 {
	return newsArticleService.INewsArticleRepository.GetNumberOfRecords()
}

func (newsArticleService *newsArticleService) GetPublishedArticles(page int, pageSize int) []model.NewsArticle {
	return newsArticleService.INewsArticleRepository.GetPublishedArticles(page, pageSize);
}

func (newsArticleService *newsArticleService) GetNumberOfRecordsPublishedArticles() int64 {
	return newsArticleService.INewsArticleRepository.GetNumberOfRecordsPublishedArticles()
}

func (newsArticleService *newsArticleService) GetById(id int) (model.NewsArticle, error) {
	return newsArticleService.INewsArticleRepository.GetById(id)
}

func (newsArticleService *newsArticleService) AddNewsArticle(newsArticle model.NewsArticle) model.NewsArticle {
	newsArticle.DateTime = time.Now()
	createdArticle := newsArticleService.INewsArticleRepository.Create(newsArticle)
	return createdArticle
}

func (newsArticleService *newsArticleService) EditNewsArticle(newsArticleDTO dto.NewsArticleDTO) string {
	msg := ""
	_, err := newsArticleService.INewsArticleRepository.GetById(newsArticleDTO.Id)
	if err != nil {
		return err.Error()
	}

	updatedNewsArticle := mapper.ToNewsArticle(newsArticleDTO)
	updatedNewsArticle.DateTime = time.Now()
	err = newsArticleService.INewsArticleRepository.Update(updatedNewsArticle)
	if err != nil {
		msg = err.Error()
	}
	return msg
}

func (newsArticleService *newsArticleService) DeleteNewsArticle(id int) error {
	newsArticle, err := newsArticleService.INewsArticleRepository.GetById(id)
	if err != nil {
		return err
	}
	*newsArticle.Archived = true
	return newsArticleService.INewsArticleRepository.Update(newsArticle)
}

func (newsArticleService *newsArticleService) PublishNewsArticle(newsArticleDTO dto.NewsArticleDTO) string {
	msg := ""
	newsArticle, err := newsArticleService.INewsArticleRepository.GetById(newsArticleDTO.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newsArticle = newsArticleService.AddNewsArticle(mapper.ToNewsArticle(newsArticleDTO))
	} else {
		newsArticle = mapper.ToNewsArticle(newsArticleDTO)
		newsArticle.DateTime = time.Now()
	}

	newsArticle.PublishedTitle = newsArticleDTO.UnpublishedTitle
	newsArticle.PublishedDescription = newsArticleDTO.UnpublishedDescription
	newsArticle.PublishedContent = newsArticleDTO.UnpublishedContent
	err = newsArticleService.INewsArticleRepository.Update(newsArticle)
	if err != nil {
		msg = err.Error()
	}
	return msg
}