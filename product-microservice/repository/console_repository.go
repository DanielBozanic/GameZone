package repository

import (
	"product/dto/filter"
	"product/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type consoleRepository struct {
	Database *gorm.DB
}

type IConsoleRepository interface {
	GetAll(page int, pageSize int) ([] model.Console)
	GetNumberOfRecords() int64
	GetById(id int) (model.Console, error)
	SearchByName(page int, pageSize int, name string) ([]model.Console, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.ConsoleFilter) ([]model.Console, error)
	GetNumberOfRecordsFilter(filter filter.ConsoleFilter) int64
	GetPlatforms() []string
	Create(videoGame model.Console) error
	Update(videoGame model.Console) error
}

func NewConsoleRepository(DB *gorm.DB) IConsoleRepository {
	return &consoleRepository{Database: DB}
}

func (consoleRepo *consoleRepository) GetAll(page int, pageSize int) []model.Console {
	var consoles []model.Console
	offset := (page - 1) * pageSize
	consoleRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = consoles.product_id").
		Where("products.archived = false").
		Order("products.price").
		Find(&consoles)
	return consoles
}

func (consoleRepo *consoleRepository) GetNumberOfRecords() int64 {
	var count int64
	consoleRepo.Database.
		Joins("JOIN products ON products.id = consoles.product_id").
		Where("products.archived = false").
		Model(&model.Console{}).
		Count(&count)
	return count
}

func (consoleRepo *consoleRepository) GetById(id int) (model.Console, error) {
	var console model.Console
	result := consoleRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = consoles.product_id").
		Where("products.archived = false").
		First(&console, id)
	return console, result.Error
}

func (consoleRepo *consoleRepository) SearchByName(page int, pageSize int, name string) ([]model.Console, error) {
	var consoles []model.Console
	offset := (page - 1) * pageSize
	result := consoleRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = consoles.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Order("products.price").
		Find(&consoles)
	return consoles, result.Error
}

func (consoleRepo *consoleRepository) GetNumberOfRecordsSearch(name string) int64 {
	var consoles []model.Console
	var count int64
	consoleRepo.Database.
		Joins("JOIN products ON products.id = consoles.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Find(&consoles).
		Count(&count)
	return count
}

func (consoleRepo *consoleRepository) Filter(page int, pageSize int, filter filter.ConsoleFilter) ([]model.Console, error) {
	var consoles []model.Console
	offset := (page - 1) * pageSize
	result := consoleRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = consoles.product_id").
		Where(`(consoles.platform IN ? OR ?) AND products.archived = false`,
				filter.Platforms,
				len(filter.Platforms) == 0).
		Order("products.price").
		Find(&consoles)
	return consoles, result.Error
}

func (consoleRepo *consoleRepository) GetNumberOfRecordsFilter(filter filter.ConsoleFilter) int64 {
	var consoles []model.Console
	var count int64
	consoleRepo.Database.
		Joins("JOIN products ON products.id = consoles.product_id").
		Where(`(consoles.platform IN ? OR ?) AND products.archived = false`,
				filter.Platforms,
				len(filter.Platforms) == 0).
		Find(&consoles).
		Count(&count)
	return count
}

func (consoleRepo *consoleRepository) GetPlatforms() []string {
	var platforms []string
	consoleRepo.Database.
		Joins("JOIN products ON products.id = consoles.product_id").
		Where("products.archived = false").
		Order("consoles.platform * 1 ASC, consoles.platform ASC").
		Model(&model.Console{}).
		Distinct().
		Pluck("consoles.platform", &platforms)
	return platforms
}

func (consoleRepo *consoleRepository) Create(console model.Console) error {
	result := consoleRepo.Database.Create(&console)
	return result.Error
}

func (consoleRepo *consoleRepository) Update(console model.Console) error {
	result := consoleRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&console)
	return result.Error
}