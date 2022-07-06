package repository

import (
	"product/dto/filter"
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type consoleRepository struct {
	Database *gorm.DB
}

type IConsoleRepository interface {
	GetAll(page int, pageSize int) ([] model.Console)
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.Console, error)
	SearchByName(page int, pageSize int, name string) ([]model.Console, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.ConsoleFilter) ([]model.Console, error)
	GetNumberOfRecordsFilter(filter filter.ConsoleFilter) int64
	GetPlatforms() []string
	Create(videoGame model.Console) error
	Update(videoGame model.Console) error
	Delete(videoGame model.Console) error
}

func NewConsoleRepository(DB *gorm.DB) IConsoleRepository {
	return &consoleRepository{Database: DB}
}

func (consoleRepo *consoleRepository) GetAll(page int, pageSize int) []model.Console {
	var consoles []model.Console
	offset := (page - 1) * pageSize
	consoleRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&consoles)
	return consoles
}

func (consoleRepo *consoleRepository) GetNumberOfRecords() int64 {
	var count int64
	consoleRepo.Database.Model(&model.Console{}).Count(&count)
	return count
}

func (consoleRepo *consoleRepository) GetById(id uuid.UUID) (model.Console, error) {
	var console model.Console
	result := consoleRepo.Database.Preload("Product").First(&console, id)
	return console, result.Error
}

func (consoleRepo *consoleRepository) SearchByName(page int, pageSize int, name string) ([]model.Console, error) {
	var consoles []model.Console
	offset := (page - 1) * pageSize
	result := consoleRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = consoles.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&consoles)
	return consoles, result.Error
}

func (consoleRepo *consoleRepository) GetNumberOfRecordsSearch(name string) int64 {
	var consoles []model.Console
	var count int64
	consoleRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = consoles.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&consoles).
		Count(&count)
	return count
}

func (consoleRepo *consoleRepository) Filter(page int, pageSize int, filter filter.ConsoleFilter) ([]model.Console, error) {
	var consoles []model.Console
	offset := (page - 1) * pageSize
	result := consoleRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = consoles.product_id").
		Where(`(platform IN ? OR ?)`,
				filter.Platforms,
				len(filter.Platforms) == 0).
		Find(&consoles)
	return consoles, result.Error
}

func (consoleRepo *consoleRepository) GetNumberOfRecordsFilter(filter filter.ConsoleFilter) int64 {
	var consoles []model.Console
	var count int64
	consoleRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = consoles.product_id").
		Where(`(platform IN ? OR ?)`,
				filter.Platforms,
				len(filter.Platforms) == 0).
		Find(&consoles).
		Count(&count)
	return count
}

func (consoleRepo *consoleRepository) GetPlatforms() []string {
	var platforms []string
	consoleRepo.Database.
		Model(&model.Ram{}).
		Distinct().
		Pluck("platform", &platforms)
	return platforms
}

func (consoleRepo *consoleRepository) Create(console model.Console) error {
	result := consoleRepo.Database.Create(&console)
	return result.Error
}

func (consoleRepo *consoleRepository) Update(console model.Console) error {
	result := consoleRepo.Database.Save(&console)
	return result.Error
}

func (consoleRepo *consoleRepository) Delete(console model.Console) error {
	result := consoleRepo.Database.Delete(&console)
	return result.Error
}