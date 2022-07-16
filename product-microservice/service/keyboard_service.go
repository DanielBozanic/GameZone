package service

import (
	"errors"
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type keyboardService struct {
	IKeyboardRepository repository.IKeyboardRepository
}

type IKeyboardService interface {
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
	Create(keyboard model.Keyboard) string
	Update(keyboardDTO dto.KeyboardDTO) string
	Delete(id uuid.UUID) error
}

func NewKeyboardService(keyboardRepository repository.IKeyboardRepository) IKeyboardService {
	return &keyboardService{IKeyboardRepository: keyboardRepository}
}

func (keyboardService *keyboardService) GetAll(page int, pageSize int) []model.Keyboard {
	return keyboardService.IKeyboardRepository.GetAll(page, pageSize)
}

func (keyboardService *keyboardService) GetNumberOfRecords() int64 {
	return keyboardService.IKeyboardRepository.GetNumberOfRecords()
}

func (keyboardService *keyboardService) GetById(id uuid.UUID) (model.Keyboard, error) {
	return keyboardService.IKeyboardRepository.GetById(id)
}

func (keyboardService *keyboardService) SearchByName(page int, pageSize int, name string) ([]model.Keyboard, error) {
	return keyboardService.IKeyboardRepository.SearchByName(page, pageSize, name)
}

func (keyboardService *keyboardService) GetNumberOfRecordsSearch(name string) int64 {
	return keyboardService.IKeyboardRepository.GetNumberOfRecordsSearch(name)
}

func (keyboardService *keyboardService) Filter(page int, pageSize int, filter filter.KeyboardFilter) ([]model.Keyboard, error) {
	return keyboardService.IKeyboardRepository.Filter(page, pageSize, filter)
}

func (keyboardService *keyboardService) GetNumberOfRecordsFilter(filter filter.KeyboardFilter) int64 {
	return keyboardService.IKeyboardRepository.GetNumberOfRecordsFilter(filter)
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

func (keyboardService *keyboardService) Create(keyboard model.Keyboard) string {
	msg := ""
	keyboard.Product.Id = uuid.New()
	keyboard.ProductId = keyboard.Product.Id
	keyboard.Product.Type = model.KEYBOARD
	err := keyboardService.IKeyboardRepository.Create(keyboard)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (keyboardService *keyboardService) Update(keyboardDTO dto.KeyboardDTO) string {
	msg := ""
	keyboard, err := keyboardService.GetById(keyboardDTO.Product.Id)
	if err != nil {
		return err.Error()
	}
	updatedKeyboard := mapper.ToKeyboard(keyboardDTO)
	updatedKeyboard.Product.Id = keyboard.Product.Id
	updatedKeyboard.ProductId = keyboard.Product.Id
	err = keyboardService.IKeyboardRepository.Update(updatedKeyboard)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (keyboardService *keyboardService) Delete(id uuid.UUID) error {
	keyboard, err := keyboardService.GetById(id)
	if err != nil {
		return err
	}
	return keyboardService.IKeyboardRepository.Delete(keyboard)
}