package service

import (
	"encoding/json"
	"net/http"
	"news/dto"
	"news/mapper"
	"news/model"
	"news/repository"
	"strconv"
	"time"
)

type newsCommentService struct {
	INewsCommentRepository repository.INewsCommentRepository
}

type INewsCommentService interface {
	GetAll() []model.NewsComment
	GetById(id int) (model.NewsComment, error)
	GetByNewsArticle(newsArticleId int) []dto.NewsCommentDTO
	AddNewsComment(newsComment model.NewsComment, userId int) string
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

func (newsCommentService *newsCommentService) GetByNewsArticle(newsArticleId int) []dto.NewsCommentDTO {
	newsCommentDTOs := []dto.NewsCommentDTO{}
	newsComments := newsCommentService.INewsCommentRepository.GetByNewsArticle(newsArticleId)
	for index, newsComment := range newsComments {
		req, err := http.NewRequest("GET", "http://localhost:5000/api/users/getById?userId=" +  strconv.Itoa(newsComment.UserId), nil)
		client := &http.Client{}
		resp, err := client.Do(req)

		username := ""
		var target map[string]interface{}
		if err != nil {
			username = "Unknown user " + strconv.Itoa(index)
		} else if resp.StatusCode != http.StatusOK {
			username = "Unknown user " + strconv.Itoa(index)
			defer resp.Body.Close()
		} else {
			json.NewDecoder(resp.Body).Decode(&target)
			username = target["user"].(map[string]interface{})["user_name"].(string)
			defer resp.Body.Close()
		}
		newsCommentDTO := mapper.ToNewsCommentDTO(newsComment)
		newsCommentDTO.Username = username
		newsCommentDTOs = append(newsCommentDTOs, newsCommentDTO)
	}
	return newsCommentDTOs
}

func (newsCommentService *newsCommentService) AddNewsComment(newsComment model.NewsComment, userId int) string {
	msg := ""
	newsComment.UserId = userId
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
	updatedNewsComment.DateTime = time.Now()
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