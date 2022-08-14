package repository

import (
	"product/dto/filter"
	"product/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type headphonesRepository struct {
	Database *gorm.DB
}

type IHeadphonesRepository interface {
	GetAll(page int, pageSize int) ([] model.Headphones)
	GetNumberOfRecords() int64
	GetById(id int) (model.Headphones, error)
	SearchByName(page int, pageSize int, name string) ([]model.Headphones, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.HeadphonesFilter) ([]model.Headphones, error)
	GetNumberOfRecordsFilter(filter filter.HeadphonesFilter) int64
	GetManufacturers() []string
	GetConnectionTypes() []string
	Create(headphones model.Headphones) error
	Update(headphones model.Headphones) error
}

func NewHeadphonesRepository(DB *gorm.DB) IHeadphonesRepository {
	return &headphonesRepository{Database: DB}
}

func (headphonesRepo *headphonesRepository) GetAll(page int, pageSize int) []model.Headphones {
	var headphones []model.Headphones
	offset := (page - 1) * pageSize
	headphonesRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = headphones.product_id").
		Where("products.archived = false").
		Order("products.price").
		Find(&headphones)
	return headphones
}

func (headphonesRepo *headphonesRepository) GetNumberOfRecords() int64 {
	var count int64
	headphonesRepo.Database.
		Joins("JOIN products ON products.id = headphones.product_id").
		Where("products.archived = false").
		Model(&model.Headphones{}).
		Count(&count)
	return count
}

func (headphonesRepo *headphonesRepository) GetById(id int) (model.Headphones, error) {
	var headphones model.Headphones
	result := headphonesRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = headphones.product_id").
		Where("products.archived = false").
		First(&headphones, id)
	return headphones, result.Error
}

func (headphonesRepo *headphonesRepository) SearchByName(page int, pageSize int, name string) ([]model.Headphones, error) {
	var headphones []model.Headphones
	offset := (page - 1) * pageSize
	result := headphonesRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = headphones.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Order("products.price").
		Find(&headphones)
	return headphones, result.Error
}

func (headphonesRepo *headphonesRepository) GetNumberOfRecordsSearch(name string) int64 {
	var headphones []model.Headphones
	var count int64
	headphonesRepo.Database.
		Joins("JOIN products ON products.id = headphones.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Find(&headphones).
		Count(&count)
	return count
}

func (headphonesRepo *headphonesRepository) Filter(page int, pageSize int, filter filter.HeadphonesFilter) ([]model.Headphones, error) {
	var headphones []model.Headphones
	offset := (page - 1) * pageSize
	result := headphonesRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = headphones.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(headphones.connection_type IN ? OR ?) AND 
				products.archived = false`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.ConnectionTypes,
				len(filter.ConnectionTypes) == 0).
		Order("products.price").
		Find(&headphones)
	return headphones, result.Error
}

func (headphonesRepo *headphonesRepository) GetNumberOfRecordsFilter(filter filter.HeadphonesFilter) int64 {
	var headphones []model.Headphones
	var count int64
	headphonesRepo.Database.
		Joins("JOIN products ON products.id = headphones.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(headphones.connection_type IN ? OR ?) AND 
				products.archived = false`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.ConnectionTypes,
				len(filter.ConnectionTypes) == 0).
		Find(&headphones).
		Count(&count)
	return count
}

func (headphonesRepo *headphonesRepository) GetManufacturers() []string {
	var manufacturers []string
	headphonesRepo.Database.
		Joins("JOIN products ON products.id = headphones.product_id").
		Where("products.archived = false").
		Order("products.manufacturer * 1 ASC, products.manufacturer ASC").
		Model(&model.Headphones{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (headphonesRepo *headphonesRepository) GetConnectionTypes() []string {
	var connectionTypes []string
	headphonesRepo.Database.
		Joins("JOIN products ON products.id = headphones.product_id").
		Where("products.archived = false").
		Order("headphones.connection_type * 1 ASC, headphones.connection_type ASC").
		Model(&model.Headphones{}).
		Distinct().
		Pluck("headphones.connection_type", &connectionTypes)
	return connectionTypes
}

func (headphonesRepo *headphonesRepository) Create(headphones model.Headphones) error {
	result := headphonesRepo.Database.Create(&headphones)
	return result.Error
}

func (headphonesRepo *headphonesRepository) Update(headphones model.Headphones) error {
	result := headphonesRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&headphones)
	return result.Error
}