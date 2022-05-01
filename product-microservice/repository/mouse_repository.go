package repository

import (
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type mouseRepository struct {
	Database *gorm.DB
}

type IMouseRepository interface {
	GetAll(page int, pageSize int) ([] model.Mouse)
	GetById(id uuid.UUID) (model.Mouse, error)
	SearchByName(page int, pageSize int, name string) ([]model.Mouse, error)
	Create(mouse model.Mouse) error
	Update(mouse model.Mouse) error
	Delete(mouse model.Mouse) error
}

func NewMouseRepository(DB *gorm.DB) IMouseRepository {
	return &mouseRepository{Database: DB}
}

func (mouseRepo *mouseRepository) GetAll(page int, pageSize int) []model.Mouse {
	var mice []model.Mouse
	offset := (page - 1) * pageSize
	mouseRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&mice)
	return mice
}

func (mouseRepo *mouseRepository) GetById(id uuid.UUID) (model.Mouse, error) {
	var mouse model.Mouse
	result := mouseRepo.Database.Preload("Product").First(&mouse, id)
	return mouse, result.Error
}

func (mouseRepo *mouseRepository) SearchByName(page int, pageSize int, name string) ([]model.Mouse, error) {
	var mice []model.Mouse
	offset := (page - 1) * pageSize
	result := mouseRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = mice.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&mice)
	return mice, result.Error
}

func (mouseRepo *mouseRepository) Create(mouse model.Mouse) error {
	result := mouseRepo.Database.Create(&mouse)
	return result.Error
}

func (mouseRepo *mouseRepository) Update(mouse model.Mouse) error {
	result := mouseRepo.Database.Save(&mouse)
	return result.Error
}

func (mouseRepo *mouseRepository) Delete(mouse model.Mouse) error {
	result := mouseRepo.Database.Delete(&mouse)
	return result.Error
}