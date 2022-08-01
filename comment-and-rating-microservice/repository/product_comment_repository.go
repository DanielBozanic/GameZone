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
	GetByProductId(productId int) []model.ProductComment
	GetByUserId(userId int) []model.ProductComment
	GetByProductAndUser(productId int, userId int) (model.ProductComment, error)
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

func (productCommentRepo *productCommentRepository) GetByProductId(productId int) []model.ProductComment {
	var productComments []model.ProductComment
	productCommentRepo.Database.
		Where("product_id = ? AND archived = false", productId).
		Order("date_time DESC").
		Find(&productComments)
	return productComments
}

func (productCommentRepo *productCommentRepository) GetByUserId(userId int) []model.ProductComment {
	var productComments []model.ProductComment
	productCommentRepo.Database.
		Where("user_id = ? AND archived = false", userId).
		Find(&productComments)
	return productComments
}

func (productCommentRepo *productCommentRepository) GetByProductAndUser(productId int, userId int) (model.ProductComment, error) {
	var productComment model.ProductComment
	result := productCommentRepo.Database.
		Where("product_id = ? AND user_id = ? AND archived = false", productId, userId).
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