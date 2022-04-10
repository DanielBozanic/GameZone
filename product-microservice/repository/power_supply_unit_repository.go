package repository

import (
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type powerSupplyUnitRepository struct {
	Database *gorm.DB
}

type IPowerSupplyUnitRepository interface {
	GetAll() ([] model.PowerSupplyUnit)
	GetById(id uuid.UUID) (model.PowerSupplyUnit, error)
	GetByName(name string) (model.PowerSupplyUnit, error)
	Create(powerSupplyUnit model.PowerSupplyUnit) error
	Update(powerSupplyUnit model.PowerSupplyUnit) error
	Delete(powerSupplyUnit model.PowerSupplyUnit) error
}

func NewPowerSupplyUnitRepository(DB *gorm.DB) IPowerSupplyUnitRepository {
	return &powerSupplyUnitRepository{Database: DB}
}

func (powerSupplyUnitRepo *powerSupplyUnitRepository) GetAll() []model.PowerSupplyUnit {
	var powerSupplyUnits []model.PowerSupplyUnit
	powerSupplyUnitRepo.Database.Find(&powerSupplyUnits)
	return powerSupplyUnits
}

func (powerSupplyUnitRepo *powerSupplyUnitRepository) GetById(id uuid.UUID) (model.PowerSupplyUnit, error) {
	var powerSupplyUnit model.PowerSupplyUnit
	result := powerSupplyUnitRepo.Database.First(&powerSupplyUnit, id)
	return powerSupplyUnit, result.Error
}

func (powerSupplyUnitRepo *powerSupplyUnitRepository) GetByName(name string) (model.PowerSupplyUnit, error) {
	var powerSupplyUnit model.PowerSupplyUnit
	result := powerSupplyUnitRepo.Database.Find(&powerSupplyUnit, "name = ?", name)
	return powerSupplyUnit, result.Error
}

func (powerSupplyUnitRepo *powerSupplyUnitRepository) Create(powerSupplyUnit model.PowerSupplyUnit) error {
	result := powerSupplyUnitRepo.Database.Create(&powerSupplyUnit)
	return result.Error
}

func (powerSupplyUnitRepo *powerSupplyUnitRepository) Update(powerSupplyUnit model.PowerSupplyUnit) error {
	result := powerSupplyUnitRepo.Database.Save(&powerSupplyUnit)
	return result.Error
}

func (powerSupplyUnitRepo *powerSupplyUnitRepository) Delete(powerSupplyUnit model.PowerSupplyUnit) error {
	result := powerSupplyUnitRepo.Database.Delete(&powerSupplyUnit)
	return result.Error
}