package repository

import (
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type keyboardRepository struct {
	Database *gorm.DB
}

type IKeyboardRepository interface {
	GetAll(page int, pageSize int) ([] model.Keyboard)
	GetById(id uuid.UUID) (model.Keyboard, error)
	SearchByName(page int, pageSize int, name string) ([]model.Keyboard, error)
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