package service

import (
	"product/dto"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type consoleService struct {
	IConsoleRepository repository.IConsoleRepository
}

type IConsoleService interface {
	GetAll() ([] model.Console)
	GetById(id uuid.UUID) (model.Console, error)
	GetByName(name string) (model.Console, error)
	Create(videoGame model.Console) error
	Update(videoGameDTO dto.ConsoleDTO) error
	Delete(id uuid.UUID) error
}

func NewConsoleService(consoleRepository repository.IConsoleRepository) IConsoleService {
	return &consoleService{IConsoleRepository: consoleRepository}
}

func (consoleService *consoleService) GetAll() []model.Console {
	return consoleService.IConsoleRepository.GetAll()
}

func (consoleService *consoleService) GetById(id uuid.UUID) (model.Console, error) {
	return consoleService.IConsoleRepository.GetById(id)
}

func (consoleService *consoleService) GetByName(name string) (model.Console, error) {
	return consoleService.IConsoleRepository.GetByName(name)
}

func (consoleService *consoleService) Create(console model.Console) error {
	console.Id = uuid.New()
	return consoleService.IConsoleRepository.Create(console)
}

func (consoleService *consoleService) Update(consoleDTO dto.ConsoleDTO) error {
	console, err := consoleService.GetById(consoleDTO.Id)
	if err != nil {
		return err
	}
	console.Name = consoleDTO.Name
	console.Price = consoleDTO.Price
	console.Amount = consoleDTO.Amount
	console.Manufacturer = consoleDTO.Manufacturer
	console.Platform = consoleDTO.Platform
	console.Amount = consoleDTO.Amount
	return consoleService.IConsoleRepository.Update(console)
}

func (consoleService *consoleService) Delete(id uuid.UUID) error {
	console, err := consoleService.GetById(id)
	if err != nil {
		return err
	}
	return consoleService.IConsoleRepository.Delete(console)
}