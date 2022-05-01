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
	GetAll(page int, pageSize int) ([] model.PowerSupplyUnit)
	GetById(id uuid.UUID) (model.PowerSupplyUnit, error)
	SearchByName(page int, pageSize int, name string) ([]model.PowerSupplyUnit, error)
	Create(powerSupplyUnit model.PowerSupplyUnit) error
	Update(powerSupplyUnit model.PowerSupplyUnit) error
	Delete(powerSupplyUnit model.PowerSupplyUnit) error
}

func NewPowerSupplyUnitRepository(DB *gorm.DB) IPowerSupplyUnitRepository {
	return &powerSupplyUnitRepository{Database: DB}
}

func (powerSupplyUnitRepo *powerSupplyUnitRepository) GetAll(page int, pageSize int) []model.PowerSupplyUnit {
	var powerSupplyUnits []model.PowerSupplyUnit
	offset := (page - 1) * pageSize
	powerSupplyUnitRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&powerSupplyUnits)
	return powerSupplyUnits
}

func (powerSupplyUnitRepo *powerSupplyUnitRepository) GetById(id uuid.UUID) (model.PowerSupplyUnit, error) {
	var powerSupplyUnit model.PowerSupplyUnit
	result := powerSupplyUnitRepo.Database.Preload("Product").First(&powerSupplyUnit, id)
	return powerSupplyUnit, result.Error
}

func (powerSupplyUnitRepo *powerSupplyUnitRepository) SearchByName(page int, pageSize int, name string) ([]model.PowerSupplyUnit, error) {
	var powerSupplyUnits []model.PowerSupplyUnit
	offset := (page - 1) * pageSize
	result := powerSupplyUnitRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = power_supply_units.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&powerSupplyUnits)
	return powerSupplyUnits, result.Error
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