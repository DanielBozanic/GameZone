package service

import (
	"news/dto"
	"news/mapper"
	"news/model"
	"news/repository"
	"time"
)

type newsCommentService struct {
	INewsCommentRepository repository.INewsCommentRepository
}

type INewsCommentService interface {
	GetAll() []model.NewsComment
	GetById(id int) (model.NewsComment, error)
	GetByNewsArticle(newsArticleId int) []model.NewsComment
	GetByUserId(userId int) []model.NewsComment
	AddNewsComment(newsComment model.NewsComment, userData dto.UserData) string
	EditNewsCommment(newsCommentDTO dto.NewsCommentDTO) string
	DeleteNewsComment(id int) error
}

func NewNewsCommentService(newsCommentRepository repository.INewsCommentRepository) INewsCommentService {
	return &newsCommentService{INewsCommentRepository: newsCommentRepository}
}

func (newsCommentService *newsCommentService) GetAll() []model.NewsComment {
	return newsCommentService.INewsCommentRepository.GetAll()
}

func (newsCommentService *newsCommentService) GetById(id int) (model.NewsComment, error) {
	return newsCommentService.INewsCommentRepository.GetById(id)
}

func (newsCommentService *newsCommentService) GetByNewsArticle(newsArticleId int) []model.NewsComment {
	return newsCommentService.INewsCommentRepository.GetByNewsArticle(newsArticleId)
}

func (newsCommentService *newsCommentService) GetByUserId(userId int) []model.NewsComment {
	return newsCommentService.INewsCommentRepository.GetByUserId(userId)
}

func (newsCommentService *newsCommentService) AddNewsComment(newsComment model.NewsComment, userData dto.UserData) string {
	msg := ""
	newsComment.UserId = userData.Id
	newsComment.Username = userData.Username
	newsComment.DateTime = time.Now()
	err := newsCommentService.INewsCommentRepository.Create(newsComment)
	if err != nil {
		msg = err.Error()
	}
	return msg
}

func (newsCommentService *newsCommentService) EditNewsCommment(newsCommentDTO dto.NewsCommentDTO) string {
	msg := ""
	_, err := newsCommentService.INewsCommentRepository.GetById(newsCommentDTO.Id)
	if err != nil {
		return err.Error()
	}

	updatedNewsComment := mapper.ToNewsComment(newsCommentDTO)
	err = newsCommentService.INewsCommentRepository.Update(updatedNewsComment)
	if err != nil {
		msg = err.Error()
	}
	return msg
}

func (newsCommentService *newsCommentService) DeleteNewsComment(id int) error {
	newsComment, err := newsCommentService.INewsCommentRepository.GetById(id)
	if err != nil {
		return err
	}
	*newsComment.Archived = true
	return newsCommentService.INewsCommentRepository.Update(newsComment)
}