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
	GetAll() ([] model.HardDiskDrive)
	GetById(id uuid.UUID) (model.HardDiskDrive, error)
	GetByName(name string) (model.HardDiskDrive, error)
	Create(hardDiskDrive model.HardDiskDrive) error
	Update(hardDiskDrive model.HardDiskDrive) error
	Delete(hardDiskDrive model.HardDiskDrive) error
}

func NewHardDiskDriveRepository(DB *gorm.DB) IHardDiskDriveRepository {
	return &hardDiskDriveRepository{Database: DB}
}

func (hardDiskDriveRepo *hardDiskDriveRepository) GetAll() []model.HardDiskDrive {
	var hardDiskDrives []model.HardDiskDrive
	hardDiskDriveRepo.Database.Find(&hardDiskDrives)
	return hardDiskDrives
}

func (hardDiskDriveRepo *hardDiskDriveRepository) GetById(id uuid.UUID) (model.HardDiskDrive, error) {
	var hardDiskDrive model.HardDiskDrive
	result := hardDiskDriveRepo.Database.First(&hardDiskDrive, id)
	return hardDiskDrive, result.Error
}

func (hardDiskDriveRepo *hardDiskDriveRepository) GetByName(name string) (model.HardDiskDrive, error) {
	var hardDiskDrive model.HardDiskDrive
	result := hardDiskDriveRepo.Database.Find(&hardDiskDrive, "name = ?", name)
	return hardDiskDrive, result.Error
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