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
	GetAll() ([] model.Mouse)
	GetById(id uuid.UUID) (model.Mouse, error)
	GetByName(name string) (model.Mouse, error)
	Create(mouse model.Mouse) error
	Update(mouse model.Mouse) error
	Delete(mouse model.Mouse) error
}

func NewMouseRepository(DB *gorm.DB) IMouseRepository {
	return &mouseRepository{Database: DB}
}

func (mouseRepo *mouseRepository) GetAll() []model.Mouse {
	var mouses []model.Mouse
	mouseRepo.Database.Preload("Product").Find(&mouses)
	return mouses
}

func (mouseRepo *mouseRepository) GetById(id uuid.UUID) (model.Mouse, error) {
	var mouse model.Mouse
	result := mouseRepo.Database.Preload("Product").First(&mouse, id)
	return mouse, result.Error
}

func (mouseRepo *mouseRepository) GetByName(name string) (model.Mouse, error) {
	var mouse model.Mouse
	result := mouseRepo.Database.Preload("Product").Find(&mouse, "name = ?", name)
	return mouse, result.Error
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