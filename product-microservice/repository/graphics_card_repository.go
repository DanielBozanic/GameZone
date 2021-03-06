package repository

import (
	"product/dto/filter"
	"product/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type graphicsCardRepository struct {
	Database *gorm.DB
}

type IGraphicsCardRepository interface {
	GetAll(page int, pageSize int) ([] model.GraphicsCard)
	GetNumberOfRecords() int64
	GetById(id int) (model.GraphicsCard, error)
	SearchByName(page int, pageSize int, name string) ([]model.GraphicsCard, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.GraphicsCardFilter) ([]model.GraphicsCard, error)
	GetNumberOfRecordsFilter(filter filter.GraphicsCardFilter) int64
	GetManufacturers() []string
	GetChipManufacturers() []string
	GetMemorySizes() []string
	GetMemoryTypes() []string
	GetModelNames() []string
	Create(graphicsCard model.GraphicsCard) error
	Update(graphicsCard model.GraphicsCard) error
}

func NewGraphicsCardRepository(DB *gorm.DB) IGraphicsCardRepository {
	return &graphicsCardRepository{Database: DB}
}

func (graphicsCardRepo *graphicsCardRepository) GetAll(page int, pageSize int) []model.GraphicsCard {
	var graphicsCards []model.GraphicsCard
	offset := (page - 1) * pageSize
	graphicsCardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Where("products.archived = false").
		Find(&graphicsCards)
	return graphicsCards
}

func (graphicsCardRepo *graphicsCardRepository) GetNumberOfRecords() int64 {
	var count int64
	graphicsCardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Where("products.archived = false").
		Model(&model.GraphicsCard{}).
		Count(&count)
	return count
}

func (graphicsCardRepo *graphicsCardRepository) GetById(id int) (model.GraphicsCard, error) {
	var graphicsCard model.GraphicsCard
	result := graphicsCardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Where("products.archived = false").
		First(&graphicsCard, id)
	return graphicsCard, result.Error
}

func (graphicsCardRepo *graphicsCardRepository) SearchByName(page int, pageSize int, name string) ([]model.GraphicsCard, error) {
	var graphicsCards []model.GraphicsCard
	offset := (page - 1) * pageSize
	result := graphicsCardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Find(&graphicsCards)
	return graphicsCards, result.Error
}

func (graphicsCardRepo *graphicsCardRepository) GetNumberOfRecordsSearch(name string) int64 {
	var graphicsCards []model.GraphicsCard
	var count int64
	graphicsCardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Find(&graphicsCards).
		Count(&count)
	return count
}

func (graphicsCardRepo *graphicsCardRepository) Filter(page int, pageSize int, filter filter.GraphicsCardFilter) ([]model.GraphicsCard, error) {
	var graphicsCards []model.GraphicsCard
	offset := (page - 1) * pageSize
	result := graphicsCardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(chip_manufacturer IN ? OR ?) AND 
				(memory_size IN ? OR ?) AND 
				(memory_type IN ? OR ?) AND
				(model_name IN ? OR ?) AND products.archived = false`,
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

func (graphicsCardRepo *graphicsCardRepository) GetNumberOfRecordsFilter(filter filter.GraphicsCardFilter) int64 {
	var graphicsCards []model.GraphicsCard
	var count int64
	graphicsCardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(chip_manufacturer IN ? OR ?) AND 
				(memory_size IN ? OR ?) AND 
				(memory_type IN ? OR ?) AND
				(model_name IN ? OR ?) AND products.archived = false`,
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
		Find(&graphicsCards).
		Count(&count)
	return count
}

func (graphicsCardRepo *graphicsCardRepository) GetManufacturers() []string {
	var manufacturers []string
	graphicsCardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Where("products.archived = false").
		Model(&model.GraphicsCard{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (graphicsCardRepo *graphicsCardRepository) GetChipManufacturers() []string {
	var chipManufacturers []string
	graphicsCardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Where("products.archived = false").
		Model(&model.GraphicsCard{}).
		Distinct().
		Pluck("chip_manufacturer", &chipManufacturers)
	return chipManufacturers
}

func (graphicsCardRepo *graphicsCardRepository) GetMemorySizes() []string {
	var memorySizes []string
	graphicsCardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Where("products.archived = false").
		Model(&model.GraphicsCard{}).
		Distinct().
		Pluck("memory_size", &memorySizes)
	return memorySizes
}

func (graphicsCardRepo *graphicsCardRepository) GetMemoryTypes() []string {
	var memoryTypes []string
	graphicsCardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Where("products.archived = false").
		Model(&model.GraphicsCard{}).
		Distinct().
		Pluck("memory_type", &memoryTypes)
	return memoryTypes
}

func (graphicsCardRepo *graphicsCardRepository) GetModelNames() []string {
	var modelNames []string
	graphicsCardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = graphics_cards.product_id").
		Where("products.archived = false").
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
	result := graphicsCardRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&console)
	return result.Error
}