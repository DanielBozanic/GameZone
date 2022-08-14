package repository

import (
	"product/dto/filter"
	"product/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type processorRepository struct {
	Database *gorm.DB
}

type IProcessorRepository interface {
	GetAll(page int, pageSize int) ([] model.Processor)
	GetNumberOfRecords() int64
	GetById(id int) (model.Processor, error)
	SearchByName(page int, pageSize int, name string) ([]model.Processor, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.ProcessorFilter) ([]model.Processor, error)
	GetNumberOfRecordsFilter(filter filter.ProcessorFilter) int64
	GetManufacturers() []string
	GetTypes() []string
	GetSockets() []string
	GetNumberOfCores() []uint
	GetThreads() []uint
	Create(processor model.Processor) error
	Update(processor model.Processor) error
}

func NewProcessorRepository(DB *gorm.DB) IProcessorRepository {
	return &processorRepository{Database: DB}
}

func (processorRepo *processorRepository) GetAll(page int, pageSize int) []model.Processor {
	var processors []model.Processor
	offset := (page - 1) * pageSize
	processorRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = processors.product_id").
		Where("products.archived = false").
		Order("products.price").
		Find(&processors)
	return processors
}

func (processorRepo *processorRepository) GetNumberOfRecords() int64 {
	var count int64
	processorRepo.Database.
		Model(&model.Processor{}).
		Joins("JOIN products ON products.id = processors.product_id").
		Where("products.archived = false").
		Count(&count)
	return count
}

func (processorRepo *processorRepository) GetById(id int) (model.Processor, error) {
	var processor model.Processor
	result := processorRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = processors.product_id").
		Where("products.archived = false").
		First(&processor, id)
	return processor, result.Error
}

func (processorRepo *processorRepository) SearchByName(page int, pageSize int, name string) ([]model.Processor, error) {
	var processors []model.Processor
	offset := (page - 1) * pageSize
	result := processorRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = processors.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Order("products.price").
		Find(&processors)
	return processors, result.Error
}

func (processorRepo *processorRepository) GetNumberOfRecordsSearch(name string) int64 {
	var processors []model.Processor
	var count int64
	processorRepo.Database.
		Joins("JOIN products ON products.id = processors.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Find(&processors).
		Count(&count)
	return count
}

func (processorRepo *processorRepository) Filter(page int, pageSize int, filter filter.ProcessorFilter) ([]model.Processor, error) {
	var processors []model.Processor
	offset := (page - 1) * pageSize
	result := processorRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = processors.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(processors.type IN ? OR ?) AND 
				(processors.socket IN ? OR ?) AND 
				(processors.number_of_cores IN ? OR ?) AND 
				(processors.threads IN ? OR ?) AND products.archived = false`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.Types,
				len(filter.Types) == 0,
				filter.Sockets,
				len(filter.Sockets) == 0,
				filter.NumberOfCores,
				len(filter.NumberOfCores) == 0,
				filter.Threads,
				len(filter.Threads) == 0).
		Order("products.price").
		Find(&processors)
	return processors, result.Error
}

func (processorRepo *processorRepository) GetNumberOfRecordsFilter(filter filter.ProcessorFilter) int64 {
	var processors []model.Processor
	var count int64
	processorRepo.Database.
		Joins("JOIN products ON products.id = processors.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(processors.type IN ? OR ?) AND 
				(processors.socket IN ? OR ?) AND 
				(processors.number_of_cores IN ? OR ?) AND 
				(processors.threads IN ? OR ?) AND products.archived = false`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.Types,
				len(filter.Types) == 0,
				filter.Sockets,
				len(filter.Sockets) == 0,
				filter.NumberOfCores,
				len(filter.NumberOfCores) == 0,
				filter.Threads,
				len(filter.Threads) == 0).
		Find(&processors).
		Count(&count)
	return count
}

func (processorRepo *processorRepository) GetManufacturers() []string {
	var manufacturers []string
	processorRepo.Database.
		Joins("JOIN products ON products.id = processors.product_id").
		Where("products.archived = false").
		Order("products.manufacturer * 1 ASC, products.manufacturer ASC").
		Model(&model.Processor{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (processorRepo *processorRepository) GetTypes() []string {
	var types []string
	processorRepo.Database.
		Joins("JOIN products ON products.id = processors.product_id").
		Where("products.archived = false").
		Order("processors.type * 1 ASC, processors.type ASC").
		Model(&model.Processor{}).
		Distinct().
		Pluck("processors.type", &types)
	return types
}

func (processorRepo *processorRepository) GetSockets() []string {
	var sockets []string
	processorRepo.Database.
		Joins("JOIN products ON products.id = processors.product_id").
		Where("products.archived = false").
		Order("processors.socket").
		Model(&model.Processor{}).
		Distinct().
		Pluck("processors.socket", &sockets)
	return sockets
}


func (processorRepo *processorRepository) GetNumberOfCores() []uint {
	var numberOfCores []uint
	processorRepo.Database.
		Joins("JOIN products ON products.id = processors.product_id").
		Where("products.archived = false").
		Order("processors.number_of_cores * 1 ASC, processors.number_of_cores ASC").
		Model(&model.Processor{}).
		Distinct().
		Pluck("processors.number_of_cores", &numberOfCores)
	return numberOfCores
}

func (processorRepo *processorRepository) GetThreads() []uint {
	var threads []uint
	processorRepo.Database.
		Joins("JOIN products ON products.id = processors.product_id").
		Where("products.archived = false").
		Order("processors.threads * 1 ASC, processors.threads ASC").
		Model(&model.Processor{}).
		Distinct().
		Pluck("processors.threads", &threads)
	return threads
}

func (processorRepo *processorRepository) Create(processor model.Processor) error {
	result := processorRepo.Database.Create(&processor)
	return result.Error
}

func (processorRepo *processorRepository) Update(processor model.Processor) error {
	result := processorRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&processor)
	return result.Error
}