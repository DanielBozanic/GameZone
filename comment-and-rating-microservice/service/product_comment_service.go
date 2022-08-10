package service

import (
	"comment-and-rating/dto"
	"comment-and-rating/mapper"
	"comment-and-rating/model"
	"comment-and-rating/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)


type productCommentService struct {
	IProductCommentRepository repository.IProductCommentRepository
}

type IProductCommentService interface {
	GetAll() []model.ProductComment
	GetById(id int) (model.ProductComment, error)
	GetByProductId(productId int) []model.ProductComment
	GetByUserId(userId int) []model.ProductComment
	AddComment(productComment model.ProductComment, userData dto.UserData) string
	EditComment(productCommentDTO dto.ProductCommentDTO) string
	DeleteComment(id int) error
}

func NewProductCommentService(productCommentRepository repository.IProductCommentRepository) IProductCommentService {
	return &productCommentService{IProductCommentRepository: productCommentRepository}
}

func (productCommentService *productCommentService) GetAll() []model.ProductComment {
	return productCommentService.IProductCommentRepository.GetAll()
}

func (productCommentService *productCommentService) GetById(id int) (model.ProductComment, error) {
	return productCommentService.IProductCommentRepository.GetById(id)
}

func (productCommentService *productCommentService) GetByProductId(productId int) []model.ProductComment {
    return productCommentService.IProductCommentRepository.GetByProductId(productId)
}

func (productCommentService *productCommentService) GetByUserId(userId int) []model.ProductComment {
	return productCommentService.IProductCommentRepository.GetByUserId(userId)
}

func (productCommentService *productCommentService) AddComment(productComment model.ProductComment, userData dto.UserData) string {
	_, err := productCommentService.IProductCommentRepository.CheckIfUserCommented(userData.Id)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "You already left a comment and rating on this product"
	}
	productComment.UserId = userData.Id
	productComment.Username = userData.Username
	productComment.DateTime = time.Now()
	productCommentService.IProductCommentRepository.Create(productComment)
	return ""
}

func (productCommentService *productCommentService) EditComment(productCommentDTO dto.ProductCommentDTO) string {
	msg := ""
	_, err := productCommentService.IProductCommentRepository.GetById(productCommentDTO.Id)
	if err != nil {
		return err.Error()
	}

	updatedProductComment := mapper.ToProductComment(productCommentDTO)
	err = productCommentService.IProductCommentRepository.Update(updatedProductComment)
	if err != nil {
		return err.Error()
	}
	return msg
}

func (productCommentService *productCommentService) DeleteComment(id int) error {
	productComment, err := productCommentService.IProductCommentRepository.GetById(id)
	if err != nil {
		return err
	}
	*productComment.Archived = true
	return productCommentService.IProductCommentRepository.Update(productComment)
}