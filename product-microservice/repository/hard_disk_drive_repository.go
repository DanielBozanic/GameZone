package repository

import (
	"product/dto/filter"
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type hardDiskDriveRepository struct {
	Database *gorm.DB
}

type IHardDiskDriveRepository interface {
	GetAll(page int, pageSize int) ([] model.HardDiskDrive)
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.HardDiskDrive, error)
	SearchByName(page int, pageSize int, name string) ([]model.HardDiskDrive, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.HardDiskDriveFilter) ([]model.HardDiskDrive, error)
	GetNumberOfRecordsFilter(filter filter.HardDiskDriveFilter) int64
	GetCapacities() []string
	GetForms() []string
	GetManufacturers() []string
	GetDiskSpeeds() []string
	Create(hardDiskDrive model.HardDiskDrive) error
	Update(hardDiskDrive model.HardDiskDrive) error
	Delete(hardDiskDrive model.HardDiskDrive) error
}

func NewHardDiskDriveRepository(DB *gorm.DB) IHardDiskDriveRepository {
	return &hardDiskDriveRepository{Database: DB}
}

func (hardDiskDriveRepo *hardDiskDriveRepository) GetAll(page int, pageSize int) []model.HardDiskDrive {
	var hardDiskDrives []model.HardDiskDrive
	offset := (page - 1) * pageSize
	hardDiskDriveRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&hardDiskDrives)
	return hardDiskDrives
}

func (hardDiskDriveRepo *hardDiskDriveRepository) GetNumberOfRecords() int64 {
	var count int64
	hardDiskDriveRepo.Database.Model(&model.HardDiskDrive{}).Count(&count)
	return count
}

func (hardDiskDriveRepo *hardDiskDriveRepository) GetById(id uuid.UUID) (model.HardDiskDrive, error) {
	var hardDiskDrive model.HardDiskDrive
	result := hardDiskDriveRepo.Database.Preload("Product").First(&hardDiskDrive, id)
	return hardDiskDrive, result.Error
}

func (hardDiskDriveRepo *hardDiskDriveRepository) SearchByName(page int, pageSize int, name string) ([]model.HardDiskDrive, error) {
	var hardDiskDrives []model.HardDiskDrive
	offset := (page - 1) * pageSize
	result := hardDiskDriveRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = hard_disk_drives.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&hardDiskDrives)
	return hardDiskDrives, result.Error
}

func (hardDiskDriveRepo *hardDiskDriveRepository) GetNumberOfRecordsSearch(name string) int64 {
	var hardDiskDrives []model.HardDiskDrive
	var count int64
	hardDiskDriveRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = hard_disk_drives.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&hardDiskDrives).
		Count(&count)
	return count
}

func (hardDiskDriveRepo *hardDiskDriveRepository) Filter(page int, pageSize int, filter filter.HardDiskDriveFilter) ([]model.HardDiskDrive, error) {
	var hdds []model.HardDiskDrive
	offset := (page - 1) * pageSize
	result := hardDiskDriveRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = hard_disk_drives.product_id").
		Where(`(capacity IN ? OR ?) AND 
				(form IN ? OR ?) AND 
				(products.manufacturer IN ? OR ?) AND 
				(disk_speed IN ? OR ?)`,
			filter.Capacities, 
			len(filter.Capacities) == 0, 
			filter.Forms, 
			len(filter.Forms) == 0,
			filter.Manufacturers, 
			len(filter.Manufacturers) == 0,
			filter.DiskSpeeds, 
			len(filter.DiskSpeeds) == 0).
		Find(&hdds)
	return hdds, result.Error
}

func (hardDiskDriveRepo *hardDiskDriveRepository) GetNumberOfRecordsFilter(filter filter.HardDiskDriveFilter) int64 {
	var hdds []model.HardDiskDrive
	var count int64
	hardDiskDriveRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = hard_disk_drives.product_id").
		Where(`(capacity IN ? OR ?) AND 
				(form IN ? OR ?) AND 
				(products.manufacturer IN ? OR ?) AND 
				(disk_speed IN ? OR ?)`,
			filter.Capacities, 
			len(filter.Capacities) == 0, 
			filter.Forms, 
			len(filter.Forms) == 0,
			filter.Manufacturers, 
			len(filter.Manufacturers) == 0,
			filter.DiskSpeeds, 
			len(filter.DiskSpeeds) == 0).
		Find(&hdds).
		Count(&count)
	return count
}

func (hardDiskDriveRepo *hardDiskDriveRepository) GetCapacities() []string {
	var capacities []string
	hardDiskDriveRepo.Database.
		Model(&model.HardDiskDrive{}).
		Distinct().
		Pluck("capacity", &capacities)
	return capacities
}

func (hardDiskDriveRepo *hardDiskDriveRepository) GetForms() []string {
	var forms []string
	hardDiskDriveRepo.Database.
		Model(&model.HardDiskDrive{}).
		Distinct().
		Pluck("form", &forms)
	return forms
}

func (hardDiskDriveRepo *hardDiskDriveRepository) GetManufacturers() []string {
	var manufacturers []string
	hardDiskDriveRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = hard_disk_drives.product_id").
		Model(&model.HardDiskDrive{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (hardDiskDriveRepo *hardDiskDriveRepository) GetDiskSpeeds() []string {
	var diskSpeeds []string
	hardDiskDriveRepo.Database.
		Model(&model.HardDiskDrive{}).
		Distinct().
		Pluck("disk_speed", &diskSpeeds)
	return diskSpeeds
}

func (hardDiskDriveRepo *hardDiskDriveRepository) Create(hardDiskDrive model.HardDiskDrive) error {
	result := hardDiskDriveRepo.Database.Create(&hardDiskDrive)
	return result.Error
}

func (hardDiskDriveRepo *hardDiskDriveRepository) Update(hardDiskDrive model.HardDiskDrive) error {
	result := hardDiskDriveRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&hardDiskDrive)
	return result.Error
}

func (hardDiskDriveRepo *hardDiskDriveRepository) Delete(hardDiskDrive model.HardDiskDrive) error {
	result := hardDiskDriveRepo.Database.Delete(&hardDiskDrive)
	return result.Error
}