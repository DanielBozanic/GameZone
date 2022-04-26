package repository

import (
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type graphicsCardRepository struct {
	Database *gorm.DB
}

type IGraphicsCardRepository interface {
	GetAll() ([] model.GraphicsCard)
	GetById(id uuid.UUID) (model.GraphicsCard, error)
	GetByName(name string) (model.GraphicsCard, error)
	Create(graphicsCard model.GraphicsCard) error
	Update(graphicsCard model.GraphicsCard) error
	Delete(graphicsCard model.GraphicsCard) error
}

func NewGraphicsCardRepository(DB *gorm.DB) IGraphicsCardRepository {
	return &graphicsCardRepository{Database: DB}
}

func (graphicsCardRepo *graphicsCardRepository) GetAll() []model.GraphicsCard {
	var graphicsCards []model.GraphicsCard
	graphicsCardRepo.Database.Preload("Product").Find(&graphicsCards)
	return graphicsCards
}

func (graphicsCardRepo *graphicsCardRepository) GetById(id uuid.UUID) (model.GraphicsCard, error) {
	var graphicsCard model.GraphicsCard
	result := graphicsCardRepo.Database.Preload("Product").First(&graphicsCard, id)
	return graphicsCard, result.Error
}

func (graphicsCardRepo *graphicsCardRepository) GetByName(name string) (model.GraphicsCard, error) {
	var graphicsCard model.GraphicsCard
	result := graphicsCardRepo.Database.Preload("Product").Find(&graphicsCard, "name = ?", name)
	return graphicsCard, result.Error
}

func (graphicsCardRepo *graphicsCardRepository) Create(console model.GraphicsCard) error {
	result := graphicsCardRepo.Database.Create(&console)
	return result.Error
}

func (graphicsCardRepo *graphicsCardRepository) Update(console model.GraphicsCard) error {
	result := graphicsCardRepo.Database.Save(&console)
	return result.Error
}

func (graphicsCardRepo *graphicsCardRepository) Delete(console model.GraphicsCard) error {
	result := graphicsCardRepo.Database.Delete(&console)
	return result.Error
}