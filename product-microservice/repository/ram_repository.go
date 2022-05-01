package repository

import (
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ramRepository struct {
	Database *gorm.DB
}

type IRamRepository interface {
	GetAll(page int, pageSize int) ([] model.Ram)
	GetById(id uuid.UUID) (model.Ram, error)
	SearchByName(page int, pageSize int, name string) ([]model.Ram, error)
	Create(ram model.Ram) error
	Update(ram model.Ram) error
	Delete(ram model.Ram) error
}

func NewRamRepository(DB *gorm.DB) IRamRepository {
	return &ramRepository{Database: DB}
}

func (ramRepo *ramRepository) GetAll(page int, pageSize int) []model.Ram {
	var rams []model.Ram
	offset := (page - 1) * pageSize
	ramRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&rams)
	return rams
}

func (ramRepo *ramRepository) GetById(id uuid.UUID) (model.Ram, error) {
	var ram model.Ram
	result := ramRepo.Database.Preload("Product").First(&ram, id)
	return ram, result.Error
}

func (ramRepo *ramRepository) SearchByName(page int, pageSize int, name string) ([]model.Ram, error) {
	var rams []model.Ram
	offset := (page - 1) * pageSize
	result := ramRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = rams.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&rams)
	return rams, result.Error
}

func (ramRepo *ramRepository) Create(ram model.Ram) error {
	result := ramRepo.Database.Create(&ram)
	return result.Error
}

func (ramRepo *ramRepository) Update(ram model.Ram) error {
	result := ramRepo.Database.Save(&ram)
	return result.Error
}

func (ramRepo *ramRepository) Delete(ram model.Ram) error {
	result := ramRepo.Database.Delete(&ram)
	return result.Error
}