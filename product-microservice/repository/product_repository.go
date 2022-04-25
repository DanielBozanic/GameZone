package repository

import (
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productRepository struct {
	Database *gorm.DB
}

type IProductRepository interface {
	GetCurrentCart(userId int) []model.ProductPurchase
	GetPurchaseHistory(userId int) []model.ProductPurchase
	GetProductPurchaseById(purchaseId uuid.UUID) (model.ProductPurchase, error)
	GetProductById(productId uuid.UUID) (model.Product, error)
	AddPurchase(purchase model.ProductPurchase) error
	UpdatePurchase(purchase model.ProductPurchase) error
	RemoveProductFromCart(purchase model.ProductPurchase) error
}

func NewProductRepository(DB *gorm.DB) IProductRepository {
	return &productRepository{Database: DB}
}

func (productRepo *productRepository) GetCurrentCart(userId int) []model.ProductPurchase {
	var currentCart []model.ProductPurchase
	productRepo.Database.Find(&currentCart, "user_id = ? and purchase_date = null", userId)
	return currentCart
}

func (productRepo *productRepository) GetPurchaseHistory(userId int) []model.ProductPurchase {
	var purchaseHistory []model.ProductPurchase
	productRepo.Database.Find(&purchaseHistory, "user_id = ? and purchase_date != null", userId)
	return purchaseHistory
}

func (productRepo *productRepository) GetProductPurchaseById(purchaseId uuid.UUID) (model.ProductPurchase, error) {
	var productPurchase model.ProductPurchase
	result := productRepo.Database.First(&productPurchase, purchaseId)
	return productPurchase, result.Error
}

func (productRepo *productRepository) GetProductById(productId uuid.UUID) (model.Product, error) {
	var product model.Product
	result := productRepo.Database.First(&product, productId)
	return product, result.Error
}

func (productRepo *productRepository) AddPurchase(purchase model.ProductPurchase) error {
	result := productRepo.Database.Create(&purchase)
	return result.Error
}

func (productRepo *productRepository) UpdatePurchase(purchase model.ProductPurchase) error {
	result := productRepo.Database.Save(&purchase)
	return result.Error
}

func (productRepo *productRepository) RemoveProductFromCart(purchase model.ProductPurchase) error {
	result := productRepo.Database.Delete(&purchase)
	return result.Error
}