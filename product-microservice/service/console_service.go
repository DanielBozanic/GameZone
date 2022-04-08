package service

import (
	"product/dto"
	"product/mapper"
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
	updatedConsole := mapper.ToConsole(consoleDTO)
	updatedConsole.Id = console.Id
	return consoleService.IConsoleRepository.Update(updatedConsole)
}

func (consoleService *consoleService) Delete(id uuid.UUID) error {
	console, err := consoleService.GetById(id)
	if err != nil {
		return err
	}
	return consoleService.IConsoleRepository.Delete(console)
}