package repository

import (
	"product/dto"
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type solidStateDriveRepository struct {
	Database *gorm.DB
}

type ISolidStateDriveRepository interface {
	GetAll(page int, pageSize int) ([] model.SolidStateDrive)
	GetById(id uuid.UUID) (model.SolidStateDrive, error)
	SearchByName(page int, pageSize int, name string) ([]model.SolidStateDrive, error)
	Filter(page int, pageSize int, filter dto.SolidStateDriveFilter) ([]model.SolidStateDrive, error)
	GetCapacities() []string
	GetForms() []string
	GetManufacturers() []string
	GetMaxSequentialReads() []string
	GetMaxSequentialWrites() []string
	Create(solidStateDrive model.SolidStateDrive) error
	Update(solidStateDrive model.SolidStateDrive) error
	Delete(solidStateDrive model.SolidStateDrive) error
}

func NewSolidStateDriveRepository(DB *gorm.DB) ISolidStateDriveRepository {
	return &solidStateDriveRepository{Database: DB}
}

func (solidStateDriveRepo *solidStateDriveRepository) GetAll(page int, pageSize int) []model.SolidStateDrive {
	var solidStateDrives []model.SolidStateDrive
	offset := (page - 1) * pageSize
	solidStateDriveRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&solidStateDrives)
	return solidStateDrives
}

func (solidStateDriveRepo *solidStateDriveRepository) GetById(id uuid.UUID) (model.SolidStateDrive, error) {
	var solidStateDrive model.SolidStateDrive
	result := solidStateDriveRepo.Database.Preload("Product").First(&solidStateDrive, id)
	return solidStateDrive, result.Error
}

func (solidStateDriveRepo *solidStateDriveRepository) SearchByName(page int, pageSize int, name string) ([]model.SolidStateDrive, error) {
	var solidStateDrives []model.SolidStateDrive
	offset := (page - 1) * pageSize
	result := solidStateDriveRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&solidStateDrives)
	return solidStateDrives, result.Error
}

func (solidStateDriveRepo *solidStateDriveRepository) Filter(page int, pageSize int, filter dto.SolidStateDriveFilter) ([]model.SolidStateDrive, error) {
	var ssds []model.SolidStateDrive
	offset := (page - 1) * pageSize
	result := solidStateDriveRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Where(`(capacity IN ? OR ?) AND 
				(form IN ? OR ?) AND 
				(manufacturer IN ? OR ?) AND 
				(max_sequential_read IN ? OR ?) AND 
				(max_sequential_write IN ? OR ?)`,
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

func (solidStateDriveRepo *solidStateDriveRepository) GetCapacities() []string {
	var capacities []string
	solidStateDriveRepo.Database.
		Model(&model.SolidStateDrive{}).
		Distinct().
		Pluck("capacity", &capacities)
	return capacities
}

func (solidStateDriveRepo *solidStateDriveRepository) GetForms() []string {
	var forms []string
	solidStateDriveRepo.Database.
		Model(&model.SolidStateDrive{}).
		Distinct().
		Pluck("form", &forms)
	return forms
}

func (solidStateDriveRepo *solidStateDriveRepository) GetManufacturers() []string {
	var manufacturers []string
	solidStateDriveRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = solid_state_drives.product_id").
		Model(&model.SolidStateDrive{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (solidStateDriveRepo *solidStateDriveRepository) GetMaxSequentialReads() []string {
	var maxSequentialReads []string
	solidStateDriveRepo.Database.
		Model(&model.SolidStateDrive{}).
		Distinct().
		Pluck("max_sequential_read", &maxSequentialReads)
	return maxSequentialReads
}

func (solidStateDriveRepo *solidStateDriveRepository) GetMaxSequentialWrites() []string {
	var maxSequentialWrites []string
	solidStateDriveRepo.Database.
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
	result := solidStateDriveRepo.Database.Save(&solidStateDrive)
	return result.Error
}

func (solidStateDriveRepo *solidStateDriveRepository) Delete(solidStateDrive model.SolidStateDrive) error {
	result := solidStateDriveRepo.Database.Delete(&solidStateDrive)
	return result.Error
}