package repository

import (
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type motherboardRepository struct {
	Database *gorm.DB
}

type IMotherboardRepository interface {
	GetAll(page int, pageSize int) ([] model.Motherboard)
	GetById(id uuid.UUID) (model.Motherboard, error)
	SearchByName(page int, pageSize int, name string) ([]model.Motherboard, error)
	Create(motherboard model.Motherboard) error
	Update(motherboard model.Motherboard) error
	Delete(motherboard model.Motherboard) error
}

func NewMotherboardRepository(DB *gorm.DB) IMotherboardRepository {
	return &motherboardRepository{Database: DB}
}

func (motherboardRepo *motherboardRepository) GetAll(page int, pageSize int) []model.Motherboard {
	var motherboards []model.Motherboard
	offset := (page - 1) * pageSize
	motherboardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&motherboards)
	return motherboards
}

func (motherboardRepo *motherboardRepository) GetById(id uuid.UUID) (model.Motherboard, error) {
	var motherboard model.Motherboard
	result := motherboardRepo.Database.Preload("Product").First(&motherboard, id)
	return motherboard, result.Error
}

func (motherboardRepo *motherboardRepository) SearchByName(page int, pageSize int, name string) ([]model.Motherboard, error) {
	var motherboards []model.Motherboard
	offset := (page - 1) * pageSize
	result := motherboardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = motherboards.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&motherboards)
	return motherboards, result.Error
}

func (motherboardRepo *motherboardRepository) Create(motherboard model.Motherboard) error {
	result := motherboardRepo.Database.Create(&motherboard)
	return result.Error
}

func (motherboardRepo *motherboardRepository) Update(motherboard model.Motherboard) error {
	result := motherboardRepo.Database.Save(&motherboard)
	return result.Error
}

func (motherboardRepo *motherboardRepository) Delete(motherboard model.Motherboard) error {
	result := motherboardRepo.Database.Delete(&motherboard)
	return result.Error
}