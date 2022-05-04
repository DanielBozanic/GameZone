package repository

import (
	"product/dto/filter"
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
	Filter(page int, pageSize int, filter filter.PowerSupplyUnitFilter) ([]model.PowerSupplyUnit, error)
	GetManufacturers() []string
	GetPowers() []string
	GetTypes() []string
	GetFormFactors() []string
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

func (powerSupplyUnitRepo *powerSupplyUnitRepository) Filter(page int, pageSize int, filter filter.PowerSupplyUnitFilter) ([]model.PowerSupplyUnit, error) {
	var psus []model.PowerSupplyUnit
	offset := (page - 1) * pageSize
	result := powerSupplyUnitRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = power_supply_units.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(power IN ? OR ?) AND 
				(type IN ? OR ?) AND 
				(form_factor IN ? OR ?)`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.Powers,
				len(filter.Powers) == 0,
				filter.Types,
				len(filter.Types) == 0,
				filter.FormFactors,
				len(filter.FormFactors) == 0).
		Find(&psus)
	return psus, result.Error
}

func (powerSupplyUnitRepo *powerSupplyUnitRepository) GetManufacturers() []string {
	var manufacturers []string
	powerSupplyUnitRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = power_supply_units.product_id").
		Model(&model.PowerSupplyUnit{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (powerSupplyUnitRepo *powerSupplyUnitRepository) GetPowers() []string {
	var powers []string
	powerSupplyUnitRepo.Database.
		Model(&model.PowerSupplyUnit{}).
		Distinct().
		Pluck("power", &powers)
	return powers
}

func (powerSupplyUnitRepo *powerSupplyUnitRepository) GetTypes() []string {
	var types []string
	powerSupplyUnitRepo.Database.
		Model(&model.PowerSupplyUnit{}).
		Distinct().
		Pluck("type", &types)
	return types
}

func (powerSupplyUnitRepo *powerSupplyUnitRepository) GetFormFactors() []string {
	var formFactors []string
	powerSupplyUnitRepo.Database.
		Model(&model.PowerSupplyUnit{}).
		Distinct().
		Pluck("form_factor", &formFactors)
	return formFactors
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