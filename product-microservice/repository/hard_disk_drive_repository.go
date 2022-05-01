package repository

import (
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type hardDiskDriveRepository struct {
	Database *gorm.DB
}

type IHardDiskDriveRepository interface {
	GetAll(page int, pageSize int) ([] model.HardDiskDrive)
	GetById(id uuid.UUID) (model.HardDiskDrive, error)
	SearchByName(page int, pageSize int, name string) ([]model.HardDiskDrive, error)
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

func (hardDiskDriveRepo *hardDiskDriveRepository) Create(hardDiskDrive model.HardDiskDrive) error {
	result := hardDiskDriveRepo.Database.Create(&hardDiskDrive)
	return result.Error
}

func (hardDiskDriveRepo *hardDiskDriveRepository) Update(hardDiskDrive model.HardDiskDrive) error {
	result := hardDiskDriveRepo.Database.Save(&hardDiskDrive)
	return result.Error
}

func (hardDiskDriveRepo *hardDiskDriveRepository) Delete(hardDiskDrive model.HardDiskDrive) error {
	result := hardDiskDriveRepo.Database.Delete(&hardDiskDrive)
	return result.Error
}