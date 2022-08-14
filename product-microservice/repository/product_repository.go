package repository

import (
	"product/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productRepository struct {
	Database *gorm.DB
}

type IProductRepository interface {
	GetProductById(productid int) (model.Product, error)
	UpdateProduct(product model.Product) error
	SearchByName(page int, pageSize int, name string) ([]model.Product, error)
	GetNumberOfRecordsSearch(name string) int64
	GetMainPageProducts() []model.Product
	IsProductOnMainPage(productId int) (model.Product, error)
	GetPopularProducts() []model.Product
	GetRecommendedProducts(userId int) []model.Product
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
		Order("price").
		Find(&products)
	return products, result.Error
}

func (productRepo *productRepository) GetNumberOfRecordsSearch(name string) int64 {
	var products []model.Product
	var count int64
	productRepo.Database.
		Where("name LIKE ? AND archived = false", "%" + name + "%").
		Find(&products).
		Count(&count)
	return count
}

func (productRepo *productRepository) GetMainPageProducts() []model.Product {
	var products []model.Product
	productRepo.Database.
		Preload("Image").
		Where("main_page = true AND archived = false").
		Find(&products).
		Limit(9)
	return products
}

func (productRepo *productRepository) IsProductOnMainPage(productId int) (model.Product, error) {
	var product model.Product
	result := productRepo.Database.
		Preload("Image").
		Where("id = ? AND main_page = true AND archived = false", productId).
		First(&product)
	return product, result.Error
}

func (productRepo *productRepository) GetPopularProducts() []model.Product {
	var productIds []int
	productRepo.Database.
		Model(model.ProductPurchase{}).
		Preload(clause.Associations).Preload("ProductPurchaseDetail." + clause.Associations).
		Select("product_purchase_details.product_id").
		Joins("JOIN product_purchase_details ON product_purchase_details.product_purchase_id = product_purchases.id").
		Joins("JOIN products ON product_purchase_details.product_id = products.id").
		Where("product_purchases.is_paid_for = true AND products.archived = false").
		Group("product_purchase_details.product_id").
		Order("SUM(product_purchase_details.product_quantity) DESC").
		Limit(3).
		Find(&productIds)

	var products []model.Product
	productRepo.Database.
		Model(model.Product{}).
		Preload(clause.Associations).Preload("Image." + clause.Associations).
		Where("id IN ? AND archived = false", productIds).
		Find(&products)
	return products
}

func (productRepo *productRepository) GetRecommendedProducts(userId int) []model.Product {
	var productId int
	productRepo.Database.
		Model(model.ProductPurchase{}).
		Select("product_purchase_details.product_id").
		Joins("JOIN product_purchase_details ON product_purchase_details.product_purchase_id = product_purchases.id").
		Joins("JOIN products ON product_purchase_details.product_id = products.id").
		Where("product_purchases.is_paid_for = true AND product_purchases.user_id = ?", userId).
		Group("product_purchase_details.product_id AND products.archived = false").
		Order("SUM(product_purchase_details.product_quantity) DESC").
		Limit(1).
		Find(&productId)

	var product model.Product
	productRepo.Database.
		Model(model.Product{}).
		Where("id = ? AND archived = false", productId).
		Find(&product)

	products := []model.Product{}
	productRepo.Database.
		Model(model.Product{}).
		Preload(clause.Associations).Preload("Image." + clause.Associations).
		Where("manufacturer LIKE ? AND archived = false", product.Manufacturer).
		Order("RAND()").
		Limit(3).
		Find(&products)
	return products
}