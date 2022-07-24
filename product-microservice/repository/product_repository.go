package repository

import (
	"product/model"

	"gorm.io/gorm"
)

type productRepository struct {
	Database *gorm.DB
}

type IProductRepository interface {
	GetProductById(productid int) (model.Product, error)
	UpdateProduct(product model.Product) error
	SearchByName(page int, pageSize int, name string) ([]model.Product, error)
	GetNumberOfRecordsSearch(name string) int64
}

func NewProductRepository(DB *gorm.DB) IProductRepository {
	return &productRepository{Database: DB}
}

func (productRepo *productRepository) GetProductById(productId int) (model.Product, error) {
	var product model.Product
	result := productRepo.Database.
		Preload("Image").
		Where("archived = false").
		First(&product, productId)
	return product, result.Error
}

func (productRepo *productRepository) UpdateProduct(product model.Product) error {
	result := productRepo.Database.Updates(&product)
	return result.Error
}

func (productRepo *productRepository) SearchByName(page int, pageSize int, name string) ([]model.Product, error) {
	var products []model.Product
	offset := (page - 1) * pageSize
	result := productRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Image").
		Where("name LIKE ? AND archived = false", "%" + name + "%").
		Find(&products)
	return products, result.Error
}

func (productRepo *productRepository) GetNumberOfRecordsSearch(name string) int64 {
	var products []model.Product
	var count int64
	productRepo.Database.
		Preload("Image").
		Where("name LIKE ? AND archived = false", "%" + name + "%").
		Find(&products).
		Count(&count)
	return count
}