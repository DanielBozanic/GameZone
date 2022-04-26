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
	GetAll() ([] model.Headphones)
	GetById(id uuid.UUID) (model.Headphones, error)
	GetByName(name string) (model.Headphones, error)
	Create(headphones model.Headphones) error
	Update(headphones model.Headphones) error
	Delete(headphones model.Headphones) error
}

func NewHeadphonesRepository(DB *gorm.DB) IHeadphonesRepository {
	return &headphonesRepository{Database: DB}
}

func (headphonesRepo *headphonesRepository) GetAll() []model.Headphones {
	var headphones []model.Headphones
	headphonesRepo.Database.Preload("Product").Find(&headphones)
	return headphones
}

func (headphonesRepo *headphonesRepository) GetById(id uuid.UUID) (model.Headphones, error) {
	var headphones model.Headphones
	result := headphonesRepo.Database.Preload("Product").First(&headphones, id)
	return headphones, result.Error
}

func (headphonesRepo *headphonesRepository) GetByName(name string) (model.Headphones, error) {
	var headphones model.Headphones
	result := headphonesRepo.Database.Preload("Product").Find(&headphones, "name = ?", name)
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