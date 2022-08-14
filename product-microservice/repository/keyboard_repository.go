package repository

import (
	"product/dto/filter"
	"product/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type keyboardRepository struct {
	Database *gorm.DB
}

type IKeyboardRepository interface {
	GetAll(page int, pageSize int) ([] model.Keyboard)
	GetNumberOfRecords() int64
	GetById(id int) (model.Keyboard, error)
	SearchByName(page int, pageSize int, name string) ([]model.Keyboard, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.KeyboardFilter) ([]model.Keyboard, error)
	GetNumberOfRecordsFilter(filter filter.KeyboardFilter) int64
	GetManufacturers() []string
	GetKeyboardConnectors() []string
	GetKeyTypes() []string
	Create(keyboard model.Keyboard) error
	Update(keyboard model.Keyboard) error
}

func NewKeyboardRepository(DB *gorm.DB) IKeyboardRepository {
	return &keyboardRepository{Database: DB}
}

func (keyboardRepo *keyboardRepository) GetAll(page int, pageSize int) []model.Keyboard {
	var keyboards []model.Keyboard
	offset := (page - 1) * pageSize
	keyboardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = keyboards.product_id").
		Where("products.archived = false").
		Order("products.price").
		Find(&keyboards)
	return keyboards
}

func (keyboardRepo *keyboardRepository) GetNumberOfRecords() int64 {
	var count int64
	keyboardRepo.Database.
		Joins("JOIN products ON products.id = keyboards.product_id").
		Where("products.archived = false").
		Model(&model.Keyboard{}).
		Count(&count)
	return count
}

func (keyboardRepo *keyboardRepository) GetById(id int) (model.Keyboard, error) {
	var keyboard model.Keyboard
	result := keyboardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = keyboards.product_id").
		Where("products.archived = false").
		First(&keyboard, id)
	return keyboard, result.Error
}

func (keyboardRepo *keyboardRepository) SearchByName(page int, pageSize int, name string) ([]model.Keyboard, error) {
	var keyboards []model.Keyboard
	offset := (page - 1) * pageSize
	result := keyboardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = keyboards.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Order("products.price").
		Find(&keyboards)
	return keyboards, result.Error
}

func (keyboardRepo *keyboardRepository) GetNumberOfRecordsSearch(name string) int64 {
	var keyboards []model.Keyboard
	var count int64
	keyboardRepo.Database.
		Joins("JOIN products ON products.id = keyboards.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Find(&keyboards).
		Count(&count)
	return count
}

func (keyboardRepo *keyboardRepository) Filter(page int, pageSize int, filter filter.KeyboardFilter) ([]model.Keyboard, error) {
	var keyboards []model.Keyboard
	offset := (page - 1) * pageSize
	result := keyboardRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = keyboards.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(keyboards.wireless IN ? OR ?) AND 
				(keyboards.keyboard_connector IN ? OR ?) AND
				(keyboards.key_type IN ? OR ?) AND products.archived = false`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.Wireless,
				len(filter.Wireless) == 0,
				filter.KeyboardConnectors,
				len(filter.KeyboardConnectors) == 0,
				filter.KeyTypes,
				len(filter.KeyTypes) == 0).
		Order("products.price").
		Find(&keyboards)
	return keyboards, result.Error
}

func (keyboardRepo *keyboardRepository) GetNumberOfRecordsFilter(filter filter.KeyboardFilter) int64 {
	var keyboards []model.Keyboard
	var count int64
	keyboardRepo.Database.
		Joins("JOIN products ON products.id = keyboards.product_id").
		Where(`(products.manufacturer IN ? OR ?) AND 
				(keyboards.wireless IN ? OR ?) AND 
				(keyboards.keyboard_connector IN ? OR ?) AND
				(keyboards.key_type IN ? OR ?) AND products.archived = false`,
				filter.Manufacturers,
				len(filter.Manufacturers) == 0,
				filter.Wireless,
				len(filter.Wireless) == 0,
				filter.KeyboardConnectors,
				len(filter.KeyboardConnectors) == 0,
				filter.KeyTypes,
				len(filter.KeyTypes) == 0).
		Find(&keyboards).
		Count(&count)
	return count
}

func (keyboardRepo *keyboardRepository) GetManufacturers() []string {
	var manufacturers []string
	keyboardRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = keyboards.product_id").
		Where("products.archived = false").
		Order("products.manufacturer * 1 ASC, products.manufacturer ASC").
		Model(&model.Keyboard{}).
		Distinct().
		Pluck("products.manufacturer", &manufacturers)
	return manufacturers
}

func (keyboardRepo *keyboardRepository) GetKeyboardConnectors() []string {
	var keyboardConnectors []string
	keyboardRepo.Database.
		Joins("JOIN products ON products.id = keyboards.product_id").
		Where("products.archived = false").
		Order("keyboards.keyboard_connector * 1 ASC, keyboards.keyboard_connector ASC").
		Model(&model.Keyboard{}).
		Distinct().
		Pluck("keyboards.keyboard_connector", &keyboardConnectors)
	return keyboardConnectors
}

func (keyboardRepo *keyboardRepository) GetKeyTypes() []string {
	var keyTypes []string
	keyboardRepo.Database.
		Joins("JOIN products ON products.id = keyboards.product_id").
		Where("products.archived = false").
		Order("keyboards.key_type * 1 ASC, keyboards.key_type ASC").
		Model(&model.Keyboard{}).
		Distinct().
		Pluck("keyboards.key_type", &keyTypes)
	return keyTypes
}

func (keyboardRepo *keyboardRepository) Create(keyboard model.Keyboard) error {
	result := keyboardRepo.Database.Create(&keyboard)
	return result.Error
}

func (keyboardRepo *keyboardRepository) Update(keyboard model.Keyboard) error {
	result := keyboardRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&keyboard)
	return result.Error
}