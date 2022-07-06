package repository

import (
	"product/dto/filter"
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type headphonesRepository struct {
	Database *gorm.DB
}

type IHeadphonesRepository interface {
	GetAll(page int, pageSize int) ([] model.Headphones)
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.Headphones, error)
	SearchByName(page int, pageSize int, name string) ([]model.Headphones, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.HeadphonesFilter) ([]model.Headphones, error)
	GetNumberOfRecordsFilter(filter filter.HeadphonesFilter) int64
	GetManufacturers() []string
	GetConnectionTypes() []string
	Create(headphones model.Headphones) error
	Update(headphones model.Headphones) error
	Delete(headphones model.Headphones) error
}

func NewHeadphonesRepository(DB *gorm.DB) IHeadphonesRepository {
	return &headphonesRepository{Database: DB}
}

func (headphonesRepo *headphonesRepository) GetAll(page int, pageSize int) []model.Headphones {
	var headphones []model.Headphones
	offset := (page - 1) * pageSize
	headphonesRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&headphones)
	return headphones
}

func (headphonesRepo *headphonesRepository) GetNumberOfRecords() int64 {
	var count int64
	headphonesRepo.Database.Model(&model.Headphones{}).Count(&count)
	return count
}

func (headphonesRepo *headphonesRepository) GetById(id uuid.UUID) (model.Headphones, error) {
	var headphones model.Headphones
	result := headphonesRepo.Database.Preload("Product").First(&headphones, id)
	return headphones, result.Error
}

func (headphonesRepo *headphonesRepository) SearchByName(page int, pageSize int, name string) ([]model.Headphones, error) {
	var headphones []model.Headphones
	offset := (page - 1) * pageSize
	result := headphonesRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = headphones.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&headphones)
	return headphones, result.Error
}

func (headphonesRepo *headphonesRepository) GetNumberOfRecordsSearch(name string) int64 {
	var headphones []model.Headphones
	var count int64
	headphonesRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = headphones.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&headphones).
		Count(&count)
	return count
}

func (headphonesRepo *headphonesRepository) Filter(page int, pageSize int, filter filter.HeadphonesFilter) ([]model.Headphones, error) {
	var headphones []model.Headphones
	offset := (page - 1) * pageSize
	result := headphonesRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = headphones.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(wireless IN ? OR ?) AND 
				(connection_type IN ? OR ?) AND
				(microphone IN ? OR ?)`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.Wireless,
				len(filter.Wireless) == 0,
				filter.ConnectionTypes,
				len(filter.ConnectionTypes) == 0,
				filter.Microphone,
				len(filter.Microphone) == 0).
		Find(&headphones)
	return headphones, result.Error
}

func (headphonesRepo *headphonesRepository) GetNumberOfRecordsFilter(filter filter.HeadphonesFilter) int64 {
	var headphones []model.Headphones
	var count int64
	headphonesRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = headphones.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(wireless IN ? OR ?) AND 
				(connection_type IN ? OR ?) AND
				(microphone IN ? OR ?)`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.Wireless,
				len(filter.Wireless) == 0,
				filter.ConnectionTypes,
				len(filter.ConnectionTypes) == 0,
				filter.Microphone,
				len(filter.Microphone) == 0).
		Find(&headphones).
		Count(&count)
	return count
}

func (headphonesRepo *headphonesRepository) GetManufacturers() []string {
	var manufacturers []string
	headphonesRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = headphones.product_id").
		Model(&model.Headphones{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (headphonesRepo *headphonesRepository) GetConnectionTypes() []string {
	var connectionTypes []string
	headphonesRepo.Database.
		Model(&model.Headphones{}).
		Distinct().
		Pluck("connection_type", &connectionTypes)
	return connectionTypes
}

func (headphonesRepo *headphonesRepository) Create(headphones model.Headphones) error {
	result := headphonesRepo.Database.Create(&headphones)
	return result.Error
}

func (headphonesRepo *headphonesRepository) Update(headphones model.Headphones) error {
	result := headphonesRepo.Database.Save(&headphones)
	return result.Error
}

func (headphonesRepo *headphonesRepository) Delete(headphones model.Headphones) error {
	result := headphonesRepo.Database.Delete(&headphones)
	return result.Error
}