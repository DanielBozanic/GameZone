package repository

import (
	"product/dto/filter"
	"product/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type motherboardRepository struct {
	Database *gorm.DB
}

type IMotherboardRepository interface {
	GetAll(page int, pageSize int) ([] model.Motherboard)
	GetNumberOfRecords() int64
	GetById(id int) (model.Motherboard, error)
	SearchByName(page int, pageSize int, name string) ([]model.Motherboard, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.MotherboardFilter) ([]model.Motherboard, error)
	GetNumberOfRecordsFilter(filter filter.MotherboardFilter) int64
	GetManufacturers() []string
	GetProcessorTypes() []string
	GetSockets() []string
	GetFormFactors() []string
	Create(motherboard model.Motherboard) error
	Update(motherboard model.Motherboard) error
}

func NewMotherboardRepository(DB *gorm.DB) IMotherboardRepository {
	return &motherboardRepository{Database: DB}
}

func (motherboardRepo *motherboardRepository) GetAll(page int, pageSize int) []model.Motherboard {
	var motherboards []model.Motherboard
	offset := (page - 1) * pageSize
	motherboardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = motherboards.product_id").
		Where("products.archived = false").
		Find(&motherboards)
	return motherboards
}

func (motherboardRepo *motherboardRepository) GetNumberOfRecords() int64 {
	var count int64
	motherboardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = motherboards.product_id").
		Where("products.archived = false").
		Model(&model.Motherboard{}).
		Count(&count)
	return count
}

func (motherboardRepo *motherboardRepository) GetById(id int) (model.Motherboard, error) {
	var motherboard model.Motherboard
	result := motherboardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = motherboards.product_id").
		Where("products.archived = false").
		First(&motherboard, id)
	return motherboard, result.Error
}

func (motherboardRepo *motherboardRepository) SearchByName(page int, pageSize int, name string) ([]model.Motherboard, error) {
	var motherboards []model.Motherboard
	offset := (page - 1) * pageSize
	result := motherboardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = motherboards.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Find(&motherboards)
	return motherboards, result.Error
}

func (motherboardRepo *motherboardRepository) GetNumberOfRecordsSearch(name string) int64 {
	var motherboards []model.Motherboard
	var count int64
	motherboardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = motherboards.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Find(&motherboards).
		Count(&count)
	return count
}

func (motherboardRepo *motherboardRepository) Filter(page int, pageSize int, filter filter.MotherboardFilter) ([]model.Motherboard, error) {
	var motherboards []model.Motherboard
	offset := (page - 1) * pageSize
	result := motherboardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = motherboards.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(processor_type IN ? OR ?) AND 
				(socket IN ? OR ?) AND 
				(form_factor IN ? OR ?) AND products.archived = false`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.ProcessorTypes,
				len(filter.ProcessorTypes) == 0,
				filter.Sockets,
				len(filter.Sockets) == 0,
				filter.FormFactor,
				len(filter.FormFactor) == 0).
		Find(&motherboards)
	return motherboards, result.Error
}

func (motherboardRepo *motherboardRepository) GetNumberOfRecordsFilter(filter filter.MotherboardFilter) int64 {
	var motherboards []model.Motherboard
	var count int64
	motherboardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = motherboards.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(processor_type IN ? OR ?) AND 
				(socket IN ? OR ?) AND 
				(form_factor IN ? OR ?) AND products.archived = false`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.ProcessorTypes,
				len(filter.ProcessorTypes) == 0,
				filter.Sockets,
				len(filter.Sockets) == 0,
				filter.FormFactor,
				len(filter.FormFactor) == 0).
		Find(&motherboards).
		Count(&count)
	return count
}

func (motherboardRepo *motherboardRepository) GetManufacturers() []string {
	var manufacturers []string
	motherboardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = motherboards.product_id").
		Where("products.archived = false").
		Model(&model.Motherboard{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (motherboardRepo *motherboardRepository) GetProcessorTypes() []string {
	var processorTypes []string
	motherboardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = motherboards.product_id").
		Where("products.archived = false").
		Model(&model.Motherboard{}).
		Distinct().
		Pluck("processor_type", &processorTypes)
	return processorTypes
}

func (motherboardRepo *motherboardRepository) GetSockets() []string {
	var sockets []string
	motherboardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = motherboards.product_id").
		Where("products.archived = false").
		Model(&model.Motherboard{}).
		Distinct().
		Pluck("socket", &sockets)
	return sockets
}


func (motherboardRepo *motherboardRepository) GetFormFactors() []string {
	var formFactors []string
	motherboardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = motherboards.product_id").
		Where("products.archived = false").
		Model(&model.Motherboard{}).
		Distinct().
		Pluck("form_factor", &formFactors)
	return formFactors
}

func (motherboardRepo *motherboardRepository) Create(motherboard model.Motherboard) error {
	result := motherboardRepo.Database.Create(&motherboard)
	return result.Error
}

func (motherboardRepo *motherboardRepository) Update(motherboard model.Motherboard) error {
	result := motherboardRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&motherboard)
	return result.Error
}