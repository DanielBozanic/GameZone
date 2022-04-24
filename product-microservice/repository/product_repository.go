package repository

import (
	"product/model"

	"gorm.io/gorm"
)

type productRepository struct {
	Database *gorm.DB
}

type IProductRepository interface {
	AddPurchase(purchase model.ProductPurchase) error
	GetCurrentCart(userId int) []model.ProductPurchase
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