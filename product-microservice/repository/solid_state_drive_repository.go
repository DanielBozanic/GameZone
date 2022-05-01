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
	GetAll(page int, pageSize int) ([] model.SolidStateDrive)
	GetById(id uuid.UUID) (model.SolidStateDrive, error)
	SearchByName(page int, pageSize int, name string) ([]model.SolidStateDrive, error)
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