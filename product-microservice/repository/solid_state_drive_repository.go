package repository

import (
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type solidStateDriveRepository struct {
	Database *gorm.DB
}

type ISolidStateDriveRepository interface {
	GetAll() ([] model.SolidStateDrive)
	GetById(id uuid.UUID) (model.SolidStateDrive, error)
	GetByName(name string) (model.SolidStateDrive, error)
	Create(solidStateDrive model.SolidStateDrive) error
	Update(solidStateDrive model.SolidStateDrive) error
	Delete(solidStateDrive model.SolidStateDrive) error
}

func NewSolidStateDriveRepository(DB *gorm.DB) ISolidStateDriveRepository {
	return &solidStateDriveRepository{Database: DB}
}

func (solidStateDriveRepo *solidStateDriveRepository) GetAll() []model.SolidStateDrive {
	var solidStateDrives []model.SolidStateDrive
	solidStateDriveRepo.Database.Find(&solidStateDrives)
	return solidStateDrives
}

func (solidStateDriveRepo *solidStateDriveRepository) GetById(id uuid.UUID) (model.SolidStateDrive, error) {
	var solidStateDrive model.SolidStateDrive
	result := solidStateDriveRepo.Database.First(&solidStateDrive, id)
	return solidStateDrive, result.Error
}

func (solidStateDriveRepo *solidStateDriveRepository) GetByName(name string) (model.SolidStateDrive, error) {
	var solidStateDrive model.SolidStateDrive
	result := solidStateDriveRepo.Database.Find(&solidStateDrive, "name = ?", name)
	return solidStateDrive, result.Error
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