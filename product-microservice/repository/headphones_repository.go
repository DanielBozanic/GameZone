package repository

import (
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type headphonesRepository struct {
	Database *gorm.DB
}

type IHeadphonesRepository interface {
	GetAll(page int, pageSize int) ([] model.Headphones)
	GetById(id uuid.UUID) (model.Headphones, error)
	SearchByName(page int, pageSize int, name string) ([]model.Headphones, error)
	Create(headphones model.Headphones) error
	Update(headphones model.Headphones) error
	Delete(headphones model.Headphones) error
}

func NewHeadphonesRepository(DB *gorm.DB) IHeadphonesRepository {
	return &headphonesRepository{Database: DB}
}

func (headphonesRepo *headphonesRepository) GetAll(page int, pageSize int) []model.Headphones {
	var headphones []model.Headphones
	offset := (page - 1) * pageSize
	headphonesRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&headphones)
	return headphones
}

func (headphonesRepo *headphonesRepository) GetById(id uuid.UUID) (model.Headphones, error) {
	var headphones model.Headphones
	result := headphonesRepo.Database.Preload("Product").First(&headphones, id)
	return headphones, result.Error
}

func (headphonesRepo *headphonesRepository) SearchByName(page int, pageSize int, name string) ([]model.Headphones, error) {
	var headphones []model.Headphones
	offset := (page - 1) * pageSize
	result := headphonesRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = headphones.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&headphones)
	return headphones, result.Error
}

func (headphonesRepo *headphonesRepository) Create(headphones model.Headphones) error {
	result := headphonesRepo.Database.Create(&headphones)
	return result.Error
}

func (headphonesRepo *headphonesRepository) Update(headphones model.Headphones) error {
	result := headphonesRepo.Database.Save(&headphones)
	return result.Error
}

func (headphonesRepo *headphonesRepository) Delete(headphones model.Headphones) error {
	result := headphonesRepo.Database.Delete(&headphones)
	return result.Error
}