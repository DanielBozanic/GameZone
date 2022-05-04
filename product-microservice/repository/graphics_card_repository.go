package repository

import (
	"product/dto/filter"
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type graphicsCardRepository struct {
	Database *gorm.DB
}

type IGraphicsCardRepository interface {
	GetAll(page int, pageSize int) ([] model.GraphicsCard)
	GetById(id uuid.UUID) (model.GraphicsCard, error)
	SearchByName(page int, pageSize int, name string) ([]model.GraphicsCard, error)
	Filter(page int, pageSize int, filter filter.GraphicsCardFilter) ([]model.GraphicsCard, error)
	GetManufacturers() []string
	GetChipManufacturers() []string
	GetMemorySizes() []string
	GetMemoryTypes() []string
	GetModelNames() []string
	Create(graphicsCard model.GraphicsCard) error
	Update(graphicsCard model.GraphicsCard) error
	Delete(graphicsCard model.GraphicsCard) error
}

func NewGraphicsCardRepository(DB *gorm.DB) IGraphicsCardRepository {
	return &graphicsCardRepository{Database: DB}
}

func (graphicsCardRepo *graphicsCardRepository) GetAll(page int, pageSize int) []model.GraphicsCard {
	var graphicsCards []model.GraphicsCard
	offset := (page - 1) * pageSize
	graphicsCardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&graphicsCards)
	return graphicsCards
}

func (graphicsCardRepo *graphicsCardRepository) GetById(id uuid.UUID) (model.GraphicsCard, error) {
	var graphicsCard model.GraphicsCard
	result := graphicsCardRepo.Database.Preload("Product").First(&graphicsCard, id)
	return graphicsCard, result.Error
}

func (graphicsCardRepo *graphicsCardRepository) SearchByName(page int, pageSize int, name string) ([]model.GraphicsCard, error) {
	var graphicsCards []model.GraphicsCard
	offset := (page - 1) * pageSize
	result := graphicsCardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&graphicsCards)
	return graphicsCards, result.Error
}

func (graphicsCardRepo *graphicsCardRepository) Filter(page int, pageSize int, filter filter.GraphicsCardFilter) ([]model.GraphicsCard, error) {
	var graphicsCards []model.GraphicsCard
	offset := (page - 1) * pageSize
	result := graphicsCardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(chip_manufacturer IN ? OR ?) AND 
				(memory_size IN ? OR ?) AND 
				(memory_type IN ? OR ?) AND
				(model_name IN ? OR ?)`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.ChipManufacturers,
				len(filter.ChipManufacturers) == 0,
				filter.MemorySizes,
				len(filter.MemorySizes) == 0,
				filter.MemoryTypes,
				len(filter.MemoryTypes) == 0,
				filter.ModelNames,
				len(filter.ModelNames) == 0).
		Find(&graphicsCards)
	return graphicsCards, result.Error
}

func (graphicsCardRepo *graphicsCardRepository) GetManufacturers() []string {
	var manufacturers []string
	graphicsCardRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Model(&model.GraphicsCard{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (graphicsCardRepo *graphicsCardRepository) GetChipManufacturers() []string {
	var chipManufacturers []string
	graphicsCardRepo.Database.
		Model(&model.GraphicsCard{}).
		Distinct().
		Pluck("chip_manufacturer", &chipManufacturers)
	return chipManufacturers
}

func (graphicsCardRepo *graphicsCardRepository) GetMemorySizes() []string {
	var memorySizes []string
	graphicsCardRepo.Database.
		Model(&model.GraphicsCard{}).
		Distinct().
		Pluck("memory_size", &memorySizes)
	return memorySizes
}

func (graphicsCardRepo *graphicsCardRepository) GetMemoryTypes() []string {
	var memoryTypes []string
	graphicsCardRepo.Database.
		Model(&model.GraphicsCard{}).
		Distinct().
		Pluck("memory_type", &memoryTypes)
	return memoryTypes
}

func (graphicsCardRepo *graphicsCardRepository) GetModelNames() []string {
	var modelNames []string
	graphicsCardRepo.Database.
		Model(&model.GraphicsCard{}).
		Distinct().
		Pluck("model_name", &modelNames)
	return modelNames
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