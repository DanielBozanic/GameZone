package repository

import (
	"comment-and-rating/model"

	"gorm.io/gorm"
)

type productCommentRepository struct {
	Database *gorm.DB
}

type IProductCommentRepository interface {
	GetAll() []model.ProductComment
	GetById(id int) (model.ProductComment, error)
	GetByProductName(productName string) []model.ProductComment
	GetByUsername(username string) []model.ProductComment
	GetByProductNameAndUsername(productName string, username string) (model.ProductComment, error)
	Create(productComment model.ProductComment) error
	Update(productComment model.ProductComment) error
}

func NewProductCommentRepository(DB *gorm.DB) IProductCommentRepository {
	return &productCommentRepository{Database: DB}
}

func (productCommentRepo *productCommentRepository) GetAll() []model.ProductComment {
	var productComments []model.ProductComment
	productCommentRepo.Database.
		Where("archived = false").
		Find(&productComments)
	return productComments
}

func (productCommentRepo *productCommentRepository) GetById(id int) (model.ProductComment, error) {
	var productComment model.ProductComment
	result := productCommentRepo.Database.
		Where("archived = false").
		First(&productComment, id)
	return productComment, result.Error
}

func (productCommentRepo *productCommentRepository) GetByProductName(productName string) []model.ProductComment {
	var productComments []model.ProductComment
	productCommentRepo.Database.
		Where("product_name LIKE ? AND archived = false", productName).
		Find(&productComments)
	return productComments
}

func (productCommentRepo *productCommentRepository) GetByUsername(username string) []model.ProductComment {
	var productComments []model.ProductComment
	productCommentRepo.Database.
		Where("username LIKE ? AND archived = false", username).
		Find(&productComments)
	return productComments
}

func (productCommentRepo *productCommentRepository) GetByProductNameAndUsername(productName string, username string) (model.ProductComment, error) {
	var productComment model.ProductComment
	result := productCommentRepo.Database.
		Where("product_name LIKE ? AND username LIKE ? AND archived = false", productName, username).
		First(&productComment)
	return productComment, result.Error
}

func (productCommentRepo *productCommentRepository) Create(productComment model.ProductComment) error {
	result := productCommentRepo.Database.Create(&productComment)
	return result.Error
}

func (productCommentRepo *productCommentRepository) Update(productComment model.ProductComment) error {
	result := productCommentRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&productComment)
	return result.Error
}