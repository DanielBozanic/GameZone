package service

import (
	"comment-and-rating/dto"
	"comment-and-rating/mapper"
	"comment-and-rating/model"
	"comment-and-rating/repository"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)


type productCommentService struct {
	IProductCommentRepository repository.IProductCommentRepository
}

type IProductCommentService interface {
	GetAll() []model.ProductComment
	GetById(id int) (model.ProductComment, error)
	GetByProductId(productId int) []dto.ProductCommentDTO
	GetByUserId(userId int) []model.ProductComment
	GetByProductAndUser(productId int, userId int) (model.ProductComment, error)
	AddComment(productComment model.ProductComment, userId int) string
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

func (productCommentService *productCommentService) GetByProductId(productId int) []dto.ProductCommentDTO {
    productCommentDTOs := []dto.ProductCommentDTO{}
	productComments := productCommentService.IProductCommentRepository.GetByProductId(productId)
	for index, productComment := range productComments {
		req, err := http.NewRequest("GET", "http://localhost:5000/api/users/getById?userId=" +  strconv.Itoa(productComment.UserId), nil)
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
		productCommentDTO := mapper.ToProductCommentDTO(productComment)
		productCommentDTO.Username = username
		productCommentDTOs = append(productCommentDTOs, productCommentDTO)
	}
	return productCommentDTOs
}

func (productCommentService *productCommentService) GetByUserId(userId int) []model.ProductComment {
	return productCommentService.IProductCommentRepository.GetByUserId(userId)
}

func (productCommentService *productCommentService) GetByProductAndUser(productId int, userId int) (model.ProductComment, error) {
	return productCommentService.IProductCommentRepository.GetByProductAndUser(productId, userId)
}

func (productCommentService *productCommentService) AddComment(productComment model.ProductComment, userId int) string {
	msg := ""
	productComment.UserId = userId
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