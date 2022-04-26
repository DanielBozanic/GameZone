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
	GetAll() ([] model.Ram)
	GetById(id uuid.UUID) (model.Ram, error)
	GetByName(name string) (model.Ram, error)
	Create(ram model.Ram) error
	Update(ram model.Ram) error
	Delete(ram model.Ram) error
}

func NewRamRepository(DB *gorm.DB) IRamRepository {
	return &ramRepository{Database: DB}
}

func (ramRepo *ramRepository) GetAll() []model.Ram {
	var Rams []model.Ram
	ramRepo.Database.Preload("Product").Find(&Rams)
	return Rams
}

func (ramRepo *ramRepository) GetById(id uuid.UUID) (model.Ram, error) {
	var ram model.Ram
	result := ramRepo.Database.Preload("Product").First(&ram, id)
	return ram, result.Error
}

func (ramRepo *ramRepository) GetByName(name string) (model.Ram, error) {
	var ram model.Ram
	result := ramRepo.Database.Preload("Product").Find(&ram, "name = ?", name)
	return ram, result.Error
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