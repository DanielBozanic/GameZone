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
	GetCurrentCart(userId int) []model.ProductPurchase
	GetAllDigitalItemsFromCart(userId int) []model.ProductPurchase
	GetPurchaseHistory(userId int) []model.ProductPurchase
	GetProductPurchaseById(purchaseid int) (model.ProductPurchase, error)
	GetProductPurchaseFromCart(productName string, userId int) (model.ProductPurchase, error)
	GetProductById(productid int) (model.Product, error)
	SearchByName(page int, pageSize int, name string) ([]model.Product, error)
	GetNumberOfRecordsSearch(name string) int64
	AddPurchase(purchase model.ProductPurchase) error
	UpdatePurchase(purchase model.ProductPurchase) error
	RemoveProductFromCart(purchase model.ProductPurchase) error
	UpdateProduct(product model.Product) error
}

func NewProductRepository(DB *gorm.DB) IProductRepository {
	return &productRepository{Database: DB}
}

func (productRepo *productRepository) GetCurrentCart(userId int) []model.ProductPurchase {
	var currentCart []model.ProductPurchase
	productRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Where("user_id = ? AND purchase_date IS NULL", userId).
		Find(&currentCart)
	return currentCart
}

func (productRepo *productRepository) GetAllDigitalItemsFromCart(userId int) []model.ProductPurchase {
	var allDigitalItems []model.ProductPurchase
	productRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = product_purchases.product_id").
		Joins("JOIN video_games ON video_games.product_id = products.id").
		Where("products.type = 13 AND video_games.Digital = true AND user_id = ? AND purchase_date IS NULL", userId).
		Find(&allDigitalItems)
	return allDigitalItems
}

func (productRepo *productRepository) GetPurchaseHistory(userId int) []model.ProductPurchase {
	var purchaseHistory []model.ProductPurchase
	productRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Where("user_id = ? AND purchase_date IS NOT null", userId).
		Find(&purchaseHistory)
	return purchaseHistory
}

func (productRepo *productRepository) GetProductPurchaseById(purchaseId int) (model.ProductPurchase, error) {
	var productPurchase model.ProductPurchase
	result := productRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		First(&productPurchase, purchaseId)
	return productPurchase, result.Error
}

func (productRepo *productRepository) GetProductById(productId int) (model.Product, error) {
	var product model.Product
	result := productRepo.Database.
		Preload("Image").
		Where("archived = false").
		First(&product, productId)
	return product, result.Error
}

func (productRepo *productRepository) GetProductPurchaseFromCart(productName string, userId int) (model.ProductPurchase, error) {
	var productPurchase model.ProductPurchase
	result := productRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = product_purchases.product_id").
		First(&productPurchase, "products.name LIKE ? AND user_id = ? AND purchase_date IS NULL", productName, userId)
	return productPurchase, result.Error
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

func (productRepo *productRepository) AddPurchase(purchase model.ProductPurchase) error {
	result := productRepo.Database.Create(&purchase)
	return result.Error
}

func (productRepo *productRepository) UpdatePurchase(purchase model.ProductPurchase) error {
	result := productRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&purchase)
	return result.Error
}

func (productRepo *productRepository) RemoveProductFromCart(purchase model.ProductPurchase) error {
	result := productRepo.Database.Delete(&purchase)
	return result.Error
}

func (productRepo *productRepository) UpdateProduct(product model.Product) error {
	result := productRepo.Database.Updates(&product)
	return result.Error
}