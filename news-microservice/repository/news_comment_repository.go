package repository

import (
	"news/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type newsCommentRepository struct {
	Database *gorm.DB
}

type INewsCommentRepository interface {
	GetAll() []model.NewsComment
	GetById(id int) (model.NewsComment, error)
	GetByUserId(userId int) []model.NewsComment
	GetByNewsArticle(newsArticleId int) []model.NewsComment
	GetByUserIdAndNewsArticle(userId int, newsArticleId int) []model.NewsComment
	Create(newsComment model.NewsComment) error
	Update(newsComment model.NewsComment) error
}

func NewNewsCommentRepository(DB *gorm.DB) INewsCommentRepository {
	return &newsCommentRepository{Database: DB}
}

func (newsCommentRepo *newsCommentRepository) GetAll() []model.NewsComment {
	var newsComments []model.NewsComment
	newsCommentRepo.Database.
		Preload(clause.Associations).Preload("NewsArticle." + clause.Associations).
		Joins("JOIN news_articles ON news_articles.id = news_comments.news_article_id").
		Where("news_articles.archived = false AND news_comments.archived = false").
		Find(&newsComments)
	return newsComments
}

func (newsCommentRepo *newsCommentRepository) GetById(id int) (model.NewsComment, error) {
	var newsComment model.NewsComment
	result := newsCommentRepo.Database.
		Preload(clause.Associations).Preload("NewsArticle." + clause.Associations).
		Joins("JOIN news_articles ON news_articles.id = news_comments.news_article_id").
		Where("news_articles.archived = false AND news_comments.archived = false").
		First(&newsComment, id)
	return newsComment, result.Error
}

func (newsCommentRepo *newsCommentRepository) GetByUserId(userId int) []model.NewsComment {
	var newsComments []model.NewsComment
	newsCommentRepo.Database.
		Preload(clause.Associations).Preload("NewsArticle." + clause.Associations).
		Where("news_comments.user_id = ?", userId).
		Find(&newsComments)
	return newsComments
}

func (newsCommentRepo *newsCommentRepository) GetByNewsArticle(newsArticleId int) []model.NewsComment {
	var newsComments []model.NewsComment
	newsCommentRepo.Database.
		Preload(clause.Associations).Preload("NewsArticle." + clause.Associations).
		Joins("JOIN news_articles ON news_articles.id = news_comments.news_article_id").
		Where("news_articles.archived = false AND news_comments.archived = false AND news_articles.id = ?", newsArticleId).
		Order("date_time DESC").
		Find(&newsComments)
	return newsComments
}

func (newsCommentRepo *newsCommentRepository) GetByUserIdAndNewsArticle(userId int, newsArticleId int) []model.NewsComment {
	var newsComments []model.NewsComment
	newsCommentRepo.Database.
		Preload(clause.Associations).Preload("NewsArticle." + clause.Associations).
		Joins("JOIN news_articles ON news_articles.id = news_comments.news_article_id").
		Where("news_articles.archived = false AND news_comments.archived = false AND news_comments.user_id = ? AND news_articles.id = ?", userId, newsArticleId).
		Find(&newsComments)
	return newsComments
}

func (newsCommentRepo *newsCommentRepository) Create(newsComment model.NewsComment) error {
	result := newsCommentRepo.Database.Create(&newsComment)
	return result.Error
}

func (newsCommentRepo *newsCommentRepository) Update(newsComment model.NewsComment) error {
	result := newsCommentRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&newsComment)
	return result.Error
}