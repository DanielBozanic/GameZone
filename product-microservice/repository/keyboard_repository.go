package repository

import (
	"product/dto/filter"
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type keyboardRepository struct {
	Database *gorm.DB
}

type IKeyboardRepository interface {
	GetAll(page int, pageSize int) ([] model.Keyboard)
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.Keyboard, error)
	SearchByName(page int, pageSize int, name string) ([]model.Keyboard, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.KeyboardFilter) ([]model.Keyboard, error)
	GetNumberOfRecordsFilter(filter filter.KeyboardFilter) int64
	GetManufacturers() []string
	GetKeyboardConnectors() []string
	GetKeyTypes() []string
	Create(keyboard model.Keyboard) error
	Update(keyboard model.Keyboard) error
	Delete(keyboard model.Keyboard) error
}

func NewKeyboardRepository(DB *gorm.DB) IKeyboardRepository {
	return &keyboardRepository{Database: DB}
}

func (keyboardRepo *keyboardRepository) GetAll(page int, pageSize int) []model.Keyboard {
	var keyboards []model.Keyboard
	offset := (page - 1) * pageSize
	keyboardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&keyboards)
	return keyboards
}

func (keyboardRepo *keyboardRepository) GetNumberOfRecords() int64 {
	var count int64
	keyboardRepo.Database.Model(&model.Keyboard{}).Count(&count)
	return count
}

func (keyboardRepo *keyboardRepository) GetById(id uuid.UUID) (model.Keyboard, error) {
	var keyboard model.Keyboard
	result := keyboardRepo.Database.Preload("Product").First(&keyboard, id)
	return keyboard, result.Error
}

func (keyboardRepo *keyboardRepository) SearchByName(page int, pageSize int, name string) ([]model.Keyboard, error) {
	var keyboards []model.Keyboard
	offset := (page - 1) * pageSize
	result := keyboardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = keyboards.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&keyboards)
	return keyboards, result.Error
}

func (keyboardRepo *keyboardRepository) GetNumberOfRecordsSearch(name string) int64 {
	var count int64
	keyboardRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = keyboards.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Count(&count)
	return count
}

func (keyboardRepo *keyboardRepository) Filter(page int, pageSize int, filter filter.KeyboardFilter) ([]model.Keyboard, error) {
	var keyboards []model.Keyboard
	offset := (page - 1) * pageSize
	result := keyboardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = keyboards.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(wireless IN ? OR ?) AND 
				(keyboard_connector IN ? OR ?) AND
				(key_type IN ? OR ?)`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.Wireless,
				len(filter.Wireless) == 0,
				filter.KeyboardConnectors,
				len(filter.KeyboardConnectors) == 0,
				filter.KeyTypes,
				len(filter.KeyTypes) == 0).
		Find(&keyboards)
	return keyboards, result.Error
}

func (keyboardRepo *keyboardRepository) GetNumberOfRecordsFilter(filter filter.KeyboardFilter) int64 {
	var count int64
	keyboardRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = keyboards.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(wireless IN ? OR ?) AND 
				(keyboard_connector IN ? OR ?) AND
				(key_type IN ? OR ?)`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.Wireless,
				len(filter.Wireless) == 0,
				filter.KeyboardConnectors,
				len(filter.KeyboardConnectors) == 0,
				filter.KeyTypes,
				len(filter.KeyTypes) == 0).
		Count(&count)
	return count
}

func (keyboardRepo *keyboardRepository) GetManufacturers() []string {
	var manufacturers []string
	keyboardRepo.Database.
		Preload("Product").
		Joins("JOIN products ON products.id = keyboards.product_id").
		Model(&model.Keyboard{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (keyboardRepo *keyboardRepository) GetKeyboardConnectors() []string {
	var keyboardConnectors []string
	keyboardRepo.Database.
		Model(&model.Keyboard{}).
		Distinct().
		Pluck("keyboard_connector", &keyboardConnectors)
	return keyboardConnectors
}

func (keyboardRepo *keyboardRepository) GetKeyTypes() []string {
	var keyTypes []string
	keyboardRepo.Database.
		Model(&model.Keyboard{}).
		Distinct().
		Pluck("key_type", &keyTypes)
	return keyTypes
}

func (keyboardRepo *keyboardRepository) Create(keyboard model.Keyboard) error {
	result := keyboardRepo.Database.Create(&keyboard)
	return result.Error
}

func (keyboardRepo *keyboardRepository) Update(keyboard model.Keyboard) error {
	result := keyboardRepo.Database.Save(&keyboard)
	return result.Error
}

func (keyboardRepo *keyboardRepository) Delete(keyboard model.Keyboard) error {
	result := keyboardRepo.Database.Delete(&keyboard)
	return result.Error
}