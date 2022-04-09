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
	GetAll() ([] model.Motherboard)
	GetById(id uuid.UUID) (model.Motherboard, error)
	GetByName(name string) (model.Motherboard, error)
	Create(motherboard model.Motherboard) error
	Update(motherboard model.Motherboard) error
	Delete(motherboard model.Motherboard) error
}

func NewMotherboardRepository(DB *gorm.DB) IMotherboardRepository {
	return &motherboardRepository{Database: DB}
}

func (motherboardRepo *motherboardRepository) GetAll() []model.Motherboard {
	var motherboards []model.Motherboard
	motherboardRepo.Database.Find(&motherboards)
	return motherboards
}

func (motherboardRepo *motherboardRepository) GetById(id uuid.UUID) (model.Motherboard, error) {
	var motherboard model.Motherboard
	result := motherboardRepo.Database.First(&motherboard, id)
	return motherboard, result.Error
}

func (motherboardRepo *motherboardRepository) GetByName(name string) (model.Motherboard, error) {
	var motherboard model.Motherboard
	result := motherboardRepo.Database.Find(&motherboard, "name = ?", name)
	return motherboard, result.Error
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