package repository

import (
	"product/dto/filter"
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type mouseRepository struct {
	Database *gorm.DB
}

type IMouseRepository interface {
	GetAll(page int, pageSize int) ([] model.Mouse)
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.Mouse, error)
	SearchByName(page int, pageSize int, name string) ([]model.Mouse, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.MouseFilter) ([]model.Mouse, error)
	GetNumberOfRecordsFilter(filter filter.MouseFilter) int64
	GetManufacturers() []string
	GetDPIs() []string
	GetConnections() []string
	Create(mouse model.Mouse) error
	Update(mouse model.Mouse) error
	Delete(mouse model.Mouse) error
}

func NewMouseRepository(DB *gorm.DB) IMouseRepository {
	return &mouseRepository{Database: DB}
}

func (mouseRepo *mouseRepository) GetAll(page int, pageSize int) []model.Mouse {
	var mice []model.Mouse
	offset := (page - 1) * pageSize
	mouseRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&mice)
	return mice
}

func (mouseRepo *mouseRepository) GetNumberOfRecords() int64 {
	var count int64
	mouseRepo.Database.Model(&model.Mouse{}).Count(&count)
	return count
}

func (mouseRepo *mouseRepository) GetById(id uuid.UUID) (model.Mouse, error) {
	var mouse model.Mouse
	result := mouseRepo.Database.Preload("Product").First(&mouse, id)
	return mouse, result.Error
}

func (mouseRepo *mouseRepository) SearchByName(page int, pageSize int, name string) ([]model.Mouse, error) {
	var mice []model.Mouse
	offset := (page - 1) * pageSize
	result := mouseRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = mice.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&mice)
	return mice, result.Error
}

func (mouseRepo *mouseRepository) GetNumberOfRecordsSearch(name string) int64 {
	var count int64
	mouseRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = mice.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Count(&count)
	return count
}

func (mouseRepo *mouseRepository) Filter(page int, pageSize int, filter filter.MouseFilter) ([]model.Mouse, error) {
	var mice []model.Mouse
	offset := (page - 1) * pageSize
	result := mouseRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = mice.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(dpi IN ? OR ?) AND 
				(wireless IN ? OR ?) AND 
				(connection IN ? OR ?)`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.DPIs,
				len(filter.DPIs) == 0,
				filter.Wireless,
				len(filter.Wireless) == 0,
				filter.Connections,
				len(filter.Connections) == 0).
		Find(&mice)
	return mice, result.Error
}

func (mouseRepo *mouseRepository) GetNumberOfRecordsFilter(filter filter.MouseFilter) int64 {
	var count int64
	mouseRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = mice.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(dpi IN ? OR ?) AND 
				(wireless IN ? OR ?) AND 
				(connection IN ? OR ?)`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.DPIs,
				len(filter.DPIs) == 0,
				filter.Wireless,
				len(filter.Wireless) == 0,
				filter.Connections,
				len(filter.Connections) == 0).
		Count(&count)
	return count
}

func (mouseRepo *mouseRepository) GetManufacturers() []string {
	var manufacturers []string
	mouseRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = mice.product_id").
		Model(&model.Mouse{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (mouseRepo *mouseRepository) GetDPIs() []string {
	var dpis []string
	mouseRepo.Database.
		Model(&model.Mouse{}).
		Distinct().
		Pluck("dpi", &dpis)
	return dpis
}

func (mouseRepo *mouseRepository) GetConnections() []string {
	var connections []string
	mouseRepo.Database.
		Model(&model.Mouse{}).
		Distinct().
		Pluck("connection", &connections)
	return connections
}

func (mouseRepo *mouseRepository) Create(mouse model.Mouse) error {
	result := mouseRepo.Database.Create(&mouse)
	return result.Error
}

func (mouseRepo *mouseRepository) Update(mouse model.Mouse) error {
	result := mouseRepo.Database.Save(&mouse)
	return result.Error
}

func (mouseRepo *mouseRepository) Delete(mouse model.Mouse) error {
	result := mouseRepo.Database.Delete(&mouse)
	return result.Error
}