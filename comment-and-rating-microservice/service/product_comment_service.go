package service

import (
	"comment-and-rating/dto"
	"comment-and-rating/mapper"
	"comment-and-rating/model"
	"comment-and-rating/repository"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)


type productCommentService struct {
	IProductCommentRepository repository.IProductCommentRepository
}

type IProductCommentService interface {
	GetAll() []model.ProductComment
	GetById(id int) (model.ProductComment, error)
	GetByProductName(productName string) []model.ProductComment
	GetByUsername(username string) []model.ProductComment
	GetByProductNameAndUsername(productName string, username string) (model.ProductComment, error)
	AddComment(productComment model.ProductComment, username string) string
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

func (productCommentService *productCommentService) GetByProductName(productName string) []model.ProductComment {
	return productCommentService.IProductCommentRepository.GetByProductName(productName)
}

func (productCommentService *productCommentService) GetByUsername(username string) []model.ProductComment {
	return productCommentService.IProductCommentRepository.GetByUsername(username)
}

func (productCommentService *productCommentService) GetByProductNameAndUsername(productName string, username string) (model.ProductComment, error) {
	return productCommentService.IProductCommentRepository.GetByProductNameAndUsername(productName, username)
}

func (productCommentService *productCommentService) AddComment(productComment model.ProductComment, username string) string {
	msg := ""
	productComment.Username = username
	productComment.DateTime = time.Now()
	err := productCommentService.IProductCommentRepository.Create(productComment)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
		msg = "You already left a comment and rating on this product"
	}
	return msg
}

func (productCommentService *productCommentService) EditComment(productCommentDTO dto.ProductCommentDTO) string {
	msg := ""
	_, err := productCommentService.IProductCommentRepository.GetById(productCommentDTO.Id)
	if err != nil {
		return err.Error()
	}

	updatedProductComment := mapper.ToProductComment(productCommentDTO)
	updatedProductComment.DateTime = time.Now()
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