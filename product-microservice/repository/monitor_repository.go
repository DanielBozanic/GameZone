package repository

import (
	"product/dto/filter"
	"product/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type monitorRepository struct {
	Database *gorm.DB
}

type IMonitorRepository interface {
	GetAll(page int, pageSize int) ([] model.Monitor)
	GetNumberOfRecords() int64
	GetById(id int) (model.Monitor, error)
	SearchByName(page int, pageSize int, name string) ([]model.Monitor, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.MonitorFilter) ([]model.Monitor, error)
	GetNumberOfRecordsFilter(filter filter.MonitorFilter) int64
	GetManufacturers() []string
	GetAspectRatios() []string
	GetResolutions() []string
	GetRefreshRates() []string
	Create(monitor model.Monitor) error
	Update(monitor model.Monitor) error
}

func NewMonitorRepository(DB *gorm.DB) IMonitorRepository {
	return &monitorRepository{Database: DB}
}

func (monitorRepo *monitorRepository) GetAll(page int, pageSize int) []model.Monitor {
	var monitors []model.Monitor
	offset := (page - 1) * pageSize
	monitorRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = monitors.product_id").
		Where("products.archived = false").
		Order("products.price").
		Find(&monitors)
	return monitors
}

func (monitorRepo *monitorRepository) GetNumberOfRecords() int64 {
	var count int64
	monitorRepo.Database.
		Joins("JOIN products ON products.id = monitors.product_id").
		Where("products.archived = false").
		Model(&model.Monitor{}).
		Count(&count)
	return count
}

func (monitorRepo *monitorRepository) GetById(id int) (model.Monitor, error) {
	var monitor model.Monitor
	result := monitorRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = monitors.product_id").
		Where("products.archived = false").
		First(&monitor, id)
	return monitor, result.Error
}

func (monitorRepo *monitorRepository) SearchByName(page int, pageSize int, name string) ([]model.Monitor, error) {
	var monitors []model.Monitor
	offset := (page - 1) * pageSize
	result := monitorRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = monitors.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Order("products.price").
		Find(&monitors)
	return monitors, result.Error
}

func (monitorRepo *monitorRepository) GetNumberOfRecordsSearch(name string) int64 {
	var monitors []model.Monitor
	var count int64
	monitorRepo.Database.
		Joins("JOIN products ON products.id = monitors.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Find(&monitors).
		Count(&count)
	return count
}

func (monitorRepo *monitorRepository) Filter(page int, pageSize int, filter filter.MonitorFilter) ([]model.Monitor, error) {
	var monitors []model.Monitor
	offset := (page - 1) * pageSize
	result := monitorRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = monitors.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(monitors.aspect_ratio IN ? OR ?) AND 
				(monitors.resolution IN ? OR ?) AND 
				(monitors.refresh_rate IN ? OR ?) AND products.archived = false`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.AspectRatios,
				len(filter.AspectRatios) == 0,
				filter.Resolutions,
				len(filter.Resolutions) == 0,
				filter.RefreshRates,
				len(filter.RefreshRates) == 0).
		Order("products.price").
		Find(&monitors)
	return monitors, result.Error
}

func (monitorRepo *monitorRepository) GetNumberOfRecordsFilter(filter filter.MonitorFilter) int64 {
	var monitors []model.Monitor
	var count int64
	monitorRepo.Database.
		Joins("JOIN products ON products.id = monitors.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(monitors.aspect_ratio IN ? OR ?) AND 
				(monitors.resolution IN ? OR ?) AND 
				(monitors.refresh_rate IN ? OR ?) AND products.archived = false`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.AspectRatios,
				len(filter.AspectRatios) == 0,
				filter.Resolutions,
				len(filter.Resolutions) == 0,
				filter.RefreshRates,
				len(filter.RefreshRates) == 0).
		Find(&monitors).
		Count(&count)
	return count
}

func (monitorRepo *monitorRepository) GetManufacturers() []string {
	var manufacturers []string
	monitorRepo.Database.
		Joins("JOIN products ON products.id = monitors.product_id").
		Where("products.archived = false").
		Order("products.manufacturer * 1 ASC, products.manufacturer ASC").
		Model(&model.Monitor{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (monitorRepo *monitorRepository) GetAspectRatios() []string {
	var aspectRatios []string
	monitorRepo.Database.
		Joins("JOIN products ON products.id = monitors.product_id").
		Where("products.archived = false").
		Order("monitors.aspect_ratio * 1 ASC, monitors.aspect_ratio ASC").
		Model(&model.Monitor{}).
		Distinct().
		Pluck("monitors.aspect_ratio", &aspectRatios)
	return aspectRatios
}

func (monitorRepo *monitorRepository) GetResolutions() []string {
	var resolutions []string
	monitorRepo.Database.
		Joins("JOIN products ON products.id = monitors.product_id").
		Where("products.archived = false").
		Order("monitors.resolution * 1 ASC, monitors.resolution ASC").
		Model(&model.Monitor{}).
		Distinct().
		Pluck("monitors.resolution", &resolutions)
	return resolutions
}


func (monitorRepo *monitorRepository) GetRefreshRates() []string {
	var refreshRates []string
	monitorRepo.Database.
		Joins("JOIN products ON products.id = monitors.product_id").
		Where("products.archived = false").
		Order("monitors.refresh_rate * 1 ASC, monitors.refresh_rate ASC").
		Model(&model.Monitor{}).
		Distinct().
		Pluck("monitors.refresh_rate", &refreshRates)
	return refreshRates
}

func (monitorRepo *monitorRepository) Create(monitor model.Monitor) error {
	result := monitorRepo.Database.Create(&monitor)
	return result.Error
}

func (monitorRepo *monitorRepository) Update(monitor model.Monitor) error {
	result := monitorRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&monitor)
	return result.Error
}