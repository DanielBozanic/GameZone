package repository

import (
	"product/dto/filter"
	"product/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type solidStateDriveRepository struct {
	Database *gorm.DB
}

type ISolidStateDriveRepository interface {
	GetAll(page int, pageSize int) ([] model.SolidStateDrive)
	GetNumberOfRecords() int64
	GetById(id int) (model.SolidStateDrive, error)
	SearchByName(page int, pageSize int, name string) ([]model.SolidStateDrive, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.SolidStateDriveFilter) ([]model.SolidStateDrive, error)
	GetNumberOfRecordsFilter(filter filter.SolidStateDriveFilter) int64
	GetCapacities() []string
	GetForms() []string
	GetManufacturers() []string
	GetMaxSequentialReads() []string
	GetMaxSequentialWrites() []string
	Create(solidStateDrive model.SolidStateDrive) error
	Update(solidStateDrive model.SolidStateDrive) error
}

func NewSolidStateDriveRepository(DB *gorm.DB) ISolidStateDriveRepository {
	return &solidStateDriveRepository{Database: DB}
}

func (solidStateDriveRepo *solidStateDriveRepository) GetAll(page int, pageSize int) []model.SolidStateDrive {
	var solidStateDrives []model.SolidStateDrive
	offset := (page - 1) * pageSize
	solidStateDriveRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Where("products.archived = false").
		Find(&solidStateDrives)
	return solidStateDrives
}

func (solidStateDriveRepo *solidStateDriveRepository) GetNumberOfRecords() int64 {
	var count int64
	solidStateDriveRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Where("products.archived = false").
		Model(&model.SolidStateDrive{}).
		Count(&count)
	return count
}

func (solidStateDriveRepo *solidStateDriveRepository) GetById(id int) (model.SolidStateDrive, error) {
	var solidStateDrive model.SolidStateDrive
	result := solidStateDriveRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Where("products.archived = false").
		First(&solidStateDrive, id)
	return solidStateDrive, result.Error
}

func (solidStateDriveRepo *solidStateDriveRepository) SearchByName(page int, pageSize int, name string) ([]model.SolidStateDrive, error) {
	var solidStateDrives []model.SolidStateDrive
	offset := (page - 1) * pageSize
	result := solidStateDriveRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Find(&solidStateDrives)
	return solidStateDrives, result.Error
}

func (solidStateDriveRepo *solidStateDriveRepository) GetNumberOfRecordsSearch(name string) int64 {
	var ssds []model.SolidStateDrive
	var count int64
	solidStateDriveRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Find(&ssds).
		Count(&count)
	return count
}

func (solidStateDriveRepo *solidStateDriveRepository) Filter(page int, pageSize int, filter filter.SolidStateDriveFilter) ([]model.SolidStateDrive, error) {
	var ssds []model.SolidStateDrive
	offset := (page - 1) * pageSize
	result := solidStateDriveRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Where(`(capacity IN ? OR ?) AND 
				(form IN ? OR ?) AND 
				(products.manufacturer IN ? OR ?) AND 
				(max_sequential_read IN ? OR ?) AND 
				(max_sequential_write IN ? OR ?) AND products.archived = false`,
			filter.Capacities, 
			len(filter.Capacities) == 0, 
			filter.Forms, 
			len(filter.Forms) == 0,
			filter.Manufacturers, 
			len(filter.Manufacturers) == 0,
			filter.MaxSequentialReads, 
			len(filter.MaxSequentialReads) == 0,
			filter.MaxSequentialWrites, 
			len(filter.MaxSequentialWrites) == 0).
		Find(&ssds)
	return ssds, result.Error
}

func (solidStateDriveRepo *solidStateDriveRepository) GetNumberOfRecordsFilter(filter filter.SolidStateDriveFilter) int64 {
	var ssds []model.SolidStateDrive
	var count int64
	solidStateDriveRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Where(`(capacity IN ? OR ?) AND 
				(form IN ? OR ?) AND 
				(products.manufacturer IN ? OR ?) AND 
				(max_sequential_read IN ? OR ?) AND 
				(max_sequential_write IN ? OR ?) AND products.archived = false`,
			filter.Capacities, 
			len(filter.Capacities) == 0, 
			filter.Forms, 
			len(filter.Forms) == 0,
			filter.Manufacturers, 
			len(filter.Manufacturers) == 0,
			filter.MaxSequentialReads, 
			len(filter.MaxSequentialReads) == 0,
			filter.MaxSequentialWrites, 
			len(filter.MaxSequentialWrites) == 0).
		Find(&ssds).
		Count(&count)
	return count
}

func (solidStateDriveRepo *solidStateDriveRepository) GetCapacities() []string {
	var capacities []string
	solidStateDriveRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Where("products.archived = false").
		Model(&model.SolidStateDrive{}).
		Distinct().
		Pluck("capacity", &capacities)
	return capacities
}

func (solidStateDriveRepo *solidStateDriveRepository) GetForms() []string {
	var forms []string
	solidStateDriveRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Where("products.archived = false").
		Model(&model.SolidStateDrive{}).
		Distinct().
		Pluck("form", &forms)
	return forms
}

func (solidStateDriveRepo *solidStateDriveRepository) GetManufacturers() []string {
	var manufacturers []string
	solidStateDriveRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Where("products.archived = false").
		Model(&model.SolidStateDrive{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (solidStateDriveRepo *solidStateDriveRepository) GetMaxSequentialReads() []string {
	var maxSequentialReads []string
	solidStateDriveRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Where("products.archived = false").
		Model(&model.SolidStateDrive{}).
		Distinct().
		Pluck("max_sequential_read", &maxSequentialReads)
	return maxSequentialReads
}

func (solidStateDriveRepo *solidStateDriveRepository) GetMaxSequentialWrites() []string {
	var maxSequentialWrites []string
	solidStateDriveRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Where("products.archived = false").
		Model(&model.SolidStateDrive{}).
		Distinct().
		Pluck("max_sequential_write", &maxSequentialWrites)
	return maxSequentialWrites
}

func (solidStateDriveRepo *solidStateDriveRepository) Create(solidStateDrive model.SolidStateDrive) error {
	result := solidStateDriveRepo.Database.Create(&solidStateDrive)
	return result.Error
}

func (solidStateDriveRepo *solidStateDriveRepository) Update(solidStateDrive model.SolidStateDrive) error {
	result := solidStateDriveRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&solidStateDrive)
	return result.Error
}