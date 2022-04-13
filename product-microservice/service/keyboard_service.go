package service

import (
	"product/dto"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type keyboardService struct {
	IKeyboardRepository repository.IKeyboardRepository
}

type IKeyboardService interface {
	GetAll() ([] model.Keyboard)
	GetById(id uuid.UUID) (model.Keyboard, error)
	GetByName(name string) (model.Keyboard, error)
	Create(keyboard model.Keyboard) error
	Update(keyboardDTO dto.KeyboardDTO) error
	Delete(id uuid.UUID) error
}

func NewKeyboardService(keyboardRepository repository.IKeyboardRepository) IKeyboardService {
	return &keyboardService{IKeyboardRepository: keyboardRepository}
}

func (keyboardService *keyboardService) GetAll() []model.Keyboard {
	return keyboardService.IKeyboardRepository.GetAll()
}

func (keyboardService *keyboardService) GetById(id uuid.UUID) (model.Keyboard, error) {
	return keyboardService.IKeyboardRepository.GetById(id)
}

func (keyboardService *keyboardService) GetByName(name string) (model.Keyboard, error) {
	return keyboardService.IKeyboardRepository.GetByName(name)
}

func (keyboardService *keyboardService) Create(keyboard model.Keyboard) error {
	keyboard.Id = uuid.New()
	return keyboardService.IKeyboardRepository.Create(keyboard)
}

func (keyboardService *keyboardService) Update(keyboardDTO dto.KeyboardDTO) error {
	keyboard, err := keyboardService.GetById(keyboardDTO.Id)
	if err != nil {
		return err
	}
	updatedKeyboard := mapper.ToKeyboard(keyboardDTO)
	updatedKeyboard.Id = keyboard.Id
	return keyboardService.IKeyboardRepository.Update(updatedKeyboard)
}

func (keyboardService *keyboardService) Delete(id uuid.UUID) error {
	keyboard, err := keyboardService.GetById(id)
	if err != nil {
		return err
	}
	return keyboardService.IKeyboardRepository.Delete(keyboard)
}