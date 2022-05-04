package service

import (
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type keyboardService struct {
	IKeyboardRepository repository.IKeyboardRepository
}

type IKeyboardService interface {
	GetAll(page int, pageSize int) ([] model.Keyboard)
	GetById(id uuid.UUID) (model.Keyboard, error)
	SearchByName(page int, pageSize int, name string) ([]model.Keyboard, error)
	Filter(page int, pageSize int, filter filter.KeyboardFilter) ([]model.Keyboard, error)
	GetManufacturers() []string
	GetKeyboardConnectors() []string
	GetKeyTypes() []string
	Create(keyboard model.Keyboard) error
	Update(keyboardDTO dto.KeyboardDTO) error
	Delete(id uuid.UUID) error
}

func NewKeyboardService(keyboardRepository repository.IKeyboardRepository) IKeyboardService {
	return &keyboardService{IKeyboardRepository: keyboardRepository}
}

func (keyboardService *keyboardService) GetAll(page int, pageSize int) []model.Keyboard {
	return keyboardService.IKeyboardRepository.GetAll(page, pageSize)
}

func (keyboardService *keyboardService) GetById(id uuid.UUID) (model.Keyboard, error) {
	return keyboardService.IKeyboardRepository.GetById(id)
}

func (keyboardService *keyboardService) SearchByName(page int, pageSize int, name string) ([]model.Keyboard, error) {
	return keyboardService.IKeyboardRepository.SearchByName(page, pageSize, name)
}

func (keyboardService *keyboardService) Filter(page int, pageSize int, filter filter.KeyboardFilter) ([]model.Keyboard, error) {
	return keyboardService.IKeyboardRepository.Filter(page, pageSize, filter)
}

func (keyboardService *keyboardService) GetManufacturers() []string {
	return keyboardService.IKeyboardRepository.GetManufacturers()
}

func (keyboardService *keyboardService) GetKeyboardConnectors() []string {
	return keyboardService.IKeyboardRepository.GetKeyboardConnectors()
}

func (keyboardService *keyboardService) GetKeyTypes() []string {
	return keyboardService.IKeyboardRepository.GetKeyTypes()
}

func (keyboardService *keyboardService) Create(keyboard model.Keyboard) error {
	keyboard.Product.Id = uuid.New()
	keyboard.ProductId = keyboard.Product.Id
	keyboard.Product.Type = model.KEYBOARD
	return keyboardService.IKeyboardRepository.Create(keyboard)
}

func (keyboardService *keyboardService) Update(keyboardDTO dto.KeyboardDTO) error {
	keyboard, err := keyboardService.GetById(keyboardDTO.Product.Id)
	if err != nil {
		return err
	}
	updatedKeyboard := mapper.ToKeyboard(keyboardDTO)
	updatedKeyboard.Product.Id = keyboard.Product.Id
	updatedKeyboard.ProductId = keyboard.Product.Id
	return keyboardService.IKeyboardRepository.Update(updatedKeyboard)
}

func (keyboardService *keyboardService) Delete(id uuid.UUID) error {
	keyboard, err := keyboardService.GetById(id)
	if err != nil {
		return err
	}
	return keyboardService.IKeyboardRepository.Delete(keyboard)
}