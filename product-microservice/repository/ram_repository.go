package repository

import (
	"product/dto/filter"
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ramRepository struct {
	Database *gorm.DB
}

type IRamRepository interface {
	GetAll(page int, pageSize int) ([] model.Ram)
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.Ram, error)
	SearchByName(page int, pageSize int, name string) ([]model.Ram, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.RAMFilter) ([]model.Ram, error)
	GetNumberOfRecordsFilter(filter filter.RAMFilter) int64
	GetManufacturers() []string
	GetCapacities() []string
	GetMemoryTypes() []string
	GetSpeeds() []string
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

func (ramRepo *ramRepository) GetNumberOfRecords() int64 {
	var count int64
	ramRepo.Database.Model(&model.Ram{}).Count(&count)
	return count
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

func (ramRepo *ramRepository) GetNumberOfRecordsSearch(name string) int64 {
	var rams []model.Ram
	var count int64
	ramRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = rams.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&rams).
		Count(&count)
	return count
}

func (ramRepo *ramRepository) Filter(page int, pageSize int, filter filter.RAMFilter) ([]model.Ram, error) {
	var rams []model.Ram
	offset := (page - 1) * pageSize
	result := ramRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = rams.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(capacity IN ? OR ?) AND 
				(memory_type IN ? OR ?) AND 
				(speed IN ? OR ?)`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.Capacities,
				len(filter.Capacities) == 0,
				filter.MemoryTypes,
				len(filter.MemoryTypes) == 0,
				filter.Speeds, 
				len(filter.Speeds)).
		Find(&rams)
	return rams, result.Error
}

func (ramRepo *ramRepository) GetNumberOfRecordsFilter(filter filter.RAMFilter) int64 {
	var rams []model.Ram
	var count int64
	ramRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = rams.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(capacity IN ? OR ?) AND 
				(memory_type IN ? OR ?) AND 
				(speed IN ? OR ?)`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.Capacities,
				len(filter.Capacities) == 0,
				filter.MemoryTypes,
				len(filter.MemoryTypes) == 0,
				filter.Speeds, 
				len(filter.Speeds)).
		Find(&rams).
		Count(&count)
	return count
}

func (ramRepo *ramRepository) GetManufacturers() []string {
	var manufacturers []string
	ramRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = rams.product_id").
		Model(&model.Ram{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (ramRepo *ramRepository) GetCapacities() []string {
	var capacities []string
	ramRepo.Database.
		Model(&model.Ram{}).
		Distinct().
		Pluck("capacity", &capacities)
	return capacities
}

func (ramRepo *ramRepository) GetMemoryTypes() []string {
	var memoryTypes []string
	ramRepo.Database.
		Model(&model.Ram{}).
		Distinct().
		Pluck("memory_type", &memoryTypes)
	return memoryTypes
}


func (ramRepo *ramRepository) GetSpeeds() []string {
	var speeds []string
	ramRepo.Database.
		Model(&model.Ram{}).
		Distinct().
		Pluck("speed", &speeds)
	return speeds
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