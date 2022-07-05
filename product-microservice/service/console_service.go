package service

import (
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type consoleService struct {
	IConsoleRepository repository.IConsoleRepository
}

type IConsoleService interface {
	GetAll(page int, pageSize int) ([] model.Console)
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.Console, error)
	SearchByName(page int, pageSize int, name string) ([]model.Console, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.ConsoleFilter) ([]model.Console, error)
	GetNumberOfRecordsFilter(filter filter.ConsoleFilter) int64
	GetPlatforms() []string
	Create(videoGame model.Console) error
	Update(videoGameDTO dto.ConsoleDTO) error
	Delete(id uuid.UUID) error
}

func NewConsoleService(consoleRepository repository.IConsoleRepository) IConsoleService {
	return &consoleService{IConsoleRepository: consoleRepository}
}

func (consoleService *consoleService) GetAll(page int, pageSize int) []model.Console {
	return consoleService.IConsoleRepository.GetAll(page, pageSize)
}

func (consoleService *consoleService) GetNumberOfRecords() int64 {
	return consoleService.IConsoleRepository.GetNumberOfRecords()
}

func (consoleService *consoleService) GetById(id uuid.UUID) (model.Console, error) {
	return consoleService.IConsoleRepository.GetById(id)
}

func (consoleService *consoleService) SearchByName(page int, pageSize int, name string) ([]model.Console, error) {
	return consoleService.IConsoleRepository.SearchByName(page, pageSize, name)
}

func (consoleService *consoleService) GetNumberOfRecordsSearch(name string) int64 {
	return consoleService.IConsoleRepository.GetNumberOfRecordsSearch(name)
}

func (consoleService *consoleService) Filter(page int, pageSize int, filter filter.ConsoleFilter) ([]model.Console, error) {
	return consoleService.IConsoleRepository.Filter(page, pageSize, filter)
}

func (consoleService *consoleService) GetNumberOfRecordsFilter(filter filter.ConsoleFilter) int64 {
	return consoleService.IConsoleRepository.GetNumberOfRecordsFilter(filter)
}

func (consoleService *consoleService) GetPlatforms() []string {
	return consoleService.IConsoleRepository.GetPlatforms()
}

func (consoleService *consoleService) Create(console model.Console) error {
	console.Product.Id = uuid.New()
	console.ProductId = console.Product.Id
	console.Product.Type = model.CONSOLE
	return consoleService.IConsoleRepository.Create(console)
}

func (consoleService *consoleService) Update(consoleDTO dto.ConsoleDTO) error {
	console, err := consoleService.GetById(consoleDTO.Product.Id)
	if err != nil {
		return err
	}
	updatedConsole := mapper.ToConsole(consoleDTO)
	updatedConsole.Product.Id = console.Product.Id
	updatedConsole.ProductId = console.Product.Id
	return consoleService.IConsoleRepository.Update(updatedConsole)
}

func (consoleService *consoleService) Delete(id uuid.UUID) error {
	console, err := consoleService.GetById(id)
	if err != nil {
		return err
	}
	return consoleService.IConsoleRepository.Delete(console)
}