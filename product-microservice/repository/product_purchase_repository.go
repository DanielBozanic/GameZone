package repository

import (
	"product/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productPurchaseRepository struct {
	Database *gorm.DB
}

type IProductPurchaseRepository interface {
	GetPurchaseHistory(userId int, page int, pageSize int) []model.ProductPurchase
	GetNumberOfRecordsPurchaseHistory(userId int) int64
	GetProductPurchaseById(purchaseid int) (model.ProductPurchase, error)
	GetPaidProductPurchase(productId int, userId int) (model.ProductPurchase, error)
	GetUnpaidProductPurchase(productId int) (model.ProductPurchase, error)
	AddPurchase(purchase model.ProductPurchase) error
	UpdatePurchase(purchase model.ProductPurchase) error
	GetUserIdsByProductId(productId int) []int
	GetProductAlertByProductIdAndUserId(userId int, productId int) (model.ProductAlert, error)
	AddProductAlert(productAlert model.ProductAlert) error
	RemoveProductAlertByUserIdAndProductId(userId int, productId int) error
}

func NewProductPurchaseRepository(DB *gorm.DB) IProductPurchaseRepository {
	return &productPurchaseRepository{Database: DB}
}

func (productPurchaseRepository *productPurchaseRepository) GetPurchaseHistory(userId int, page int, pageSize int) []model.ProductPurchase {
	var purchaseHistory []model.ProductPurchase
	offset := (page - 1) * pageSize
	productPurchaseRepository.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("ProductPurchaseDetail." + clause.Associations).
		Where("user_id = ?", userId).
		Order("purchase_date DESC").
		Find(&purchaseHistory)
	return purchaseHistory
}

func (productPurchaseRepository *productPurchaseRepository) GetNumberOfRecordsPurchaseHistory(userId int) int64 {
	var count int64
	productPurchaseRepository.Database.
		Preload(clause.Associations).Preload("ProductPurchaseDetail." + clause.Associations).
		Where("user_id = ?", userId).
		Model(&model.ProductPurchase{}).
		Count(&count)
	return count
}

func (productPurchaseRepository *productPurchaseRepository) GetProductPurchaseById(purchaseId int) (model.ProductPurchase, error) {
	var productPurchase model.ProductPurchase
	result := productPurchaseRepository.Database.
		Preload(clause.Associations).Preload("ProductPurchaseDetail." + clause.Associations).
		First(&productPurchase, purchaseId)
	return productPurchase, result.Error
}

func (productPurchaseRepository *productPurchaseRepository) GetPaidProductPurchase(productId int, userId int) (model.ProductPurchase, error) {
	var productPurchase model.ProductPurchase
	result := productPurchaseRepository.Database.
		Joins("JOIN product_purchase_details ON product_purchase_details.product_purchase_id = product_purchases.id").
		Where("product_purchase_details.product_id = ? AND user_id = ? AND is_paid_for = true", productId, userId).
		First(&productPurchase)
	return productPurchase, result.Error
}

func (productPurchaseRepository *productPurchaseRepository) GetUnpaidProductPurchase(productId int) (model.ProductPurchase, error) {
	var productPurchase model.ProductPurchase
	result := productPurchaseRepository.Database.
		Joins("JOIN product_purchase_details ON product_purchase_details.product_purchase_id = product_purchases.id").
		Where("product_purchase_details.product_id = ? AND is_paid_for = false", productId).
		First(&productPurchase)
	return productPurchase, result.Error
}
 
func (productPurchaseRepository *productPurchaseRepository) AddPurchase(purchase model.ProductPurchase) error {
	result := productPurchaseRepository.Database.Create(&purchase)
	return result.Error
}

func (productPurchaseRepository *productPurchaseRepository) UpdatePurchase(purchase model.ProductPurchase) error {
	result := productPurchaseRepository.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&purchase)
	return result.Error
}

func (productPurchaseRepository *productPurchaseRepository) GetUserIdsByProductId(productId int) []int {
	var userIds []int
	productPurchaseRepository.Database.
		Where("product_id = ?", productId).
		Model(&model.ProductAlert{}).
		Pluck("user_id", &userIds)
	return userIds
}

func (productPurchaseRepository *productPurchaseRepository) GetProductAlertByProductIdAndUserId(userId int, productId int) (model.ProductAlert, error) {
	var productAlert model.ProductAlert
	result := productPurchaseRepository.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Where("user_id = ? AND product_id = ?", userId, productId).
		First(&productAlert)
	return productAlert, result.Error
}

func (productPurchaseRepository *productPurchaseRepository) AddProductAlert(productAlert model.ProductAlert) error {
	result := productPurchaseRepository.Database.Create(&productAlert)
	return result.Error
}

func (productPurchaseRepository *productPurchaseRepository) RemoveProductAlertByUserIdAndProductId(userId int, productId int) error {
	result := productPurchaseRepository.Database.
		Where("user_id = ? AND product_id = ?", userId, productId).
		Delete(model.ProductAlert{})
	return result.Error
}