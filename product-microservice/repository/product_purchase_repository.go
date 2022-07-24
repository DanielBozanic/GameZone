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
	GetPurchaseHistory(userId int) []model.ProductPurchase
	GetProductPurchaseById(purchaseid int) (model.ProductPurchase, error)
	GetPaidProductPurchase(productId int, userId int) (model.ProductPurchase, error)
	GetUnpaidProductPurchase(productId int) (model.ProductPurchase, error)
	AddPurchase(purchase model.ProductPurchase) error
	UpdatePurchase(purchase model.ProductPurchase) error
	GetUserEmailsByProductId(productId int) []string
	GetProductAlertByProductIdAndEmail(email string, productId int) (model.ProductAlert, error)
	AddProductAlert(productAlert model.ProductAlert) error
	RemoveProductAlertByEmailAndProductId(email string, productId int) error
}

func NewProductPurchaseRepository(DB *gorm.DB) IProductPurchaseRepository {
	return &productPurchaseRepository{Database: DB}
}

func (productPurchaseRepository *productPurchaseRepository) GetPurchaseHistory(userId int) []model.ProductPurchase {
	var purchaseHistory []model.ProductPurchase
	productPurchaseRepository.Database.
		Where("user_id = ?", userId).
		Find(&purchaseHistory)
	return purchaseHistory
}

func (productPurchaseRepository *productPurchaseRepository) GetProductPurchaseById(purchaseId int) (model.ProductPurchase, error) {
	var productPurchase model.ProductPurchase
	result := productPurchaseRepository.Database.
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

func (productPurchaseRepository *productPurchaseRepository) GetUserEmailsByProductId(productId int) []string {
	var userEmails []string
	productPurchaseRepository.Database.
		Where("product_id = ?", productId).
		Model(&model.ProductAlert{}).
		Pluck("user_email", &userEmails)
	return userEmails
}

func (productPurchaseRepository *productPurchaseRepository) GetProductAlertByProductIdAndEmail(email string, productId int) (model.ProductAlert, error) {
	var productAlert model.ProductAlert
	result := productPurchaseRepository.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Where("user_email LIKE ? AND product_id = ?", email, productId).
		First(&productAlert)
	return productAlert, result.Error
}

func (productPurchaseRepository *productPurchaseRepository) AddProductAlert(productAlert model.ProductAlert) error {
	result := productPurchaseRepository.Database.Create(&productAlert)
	return result.Error
}

func (productPurchaseRepository *productPurchaseRepository)  RemoveProductAlertByEmailAndProductId(email string, productId int) error {
	result := productPurchaseRepository.Database.
		Where("user_email LIKE ? AND product_id = ?", email, productId).
		Delete(model.ProductAlert{})
	return result.Error
}