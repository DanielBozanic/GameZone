package service

import (
	"errors"
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/go-sql-driver/mysql"
)

type consoleService struct {
	IConsoleRepository repository.IConsoleRepository
}

type IConsoleService interface {
	GetAll(page int, pageSize int) ([] model.Console)
	GetNumberOfRecords() int64
	GetById(id int) (model.Console, error)
	SearchByName(page int, pageSize int, name string) ([]model.Console, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.ConsoleFilter) ([]model.Console, error)
	GetNumberOfRecordsFilter(filter filter.ConsoleFilter) int64
	GetPlatforms() []string
	Create(videoGame model.Console) string
	Update(videoGameDTO dto.ConsoleDTO) string
	Delete(id int) error
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

func (consoleService *consoleService) GetById(id int) (model.Console, error) {
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

func (consoleService *consoleService) Create(console model.Console) string {
	msg := ""
	console.Product.Type = model.CONSOLE
	err := consoleService.IConsoleRepository.Create(console)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (consoleService *consoleService) Update(consoleDTO dto.ConsoleDTO) string {
	msg := ""
	console, err := consoleService.GetById(consoleDTO.Product.Id)
	if err != nil {
		return err.Error()
	}
	updatedConsole := mapper.ToConsole(consoleDTO)
	updatedConsole.Product.Id = console.Product.Id
	updatedConsole.ProductId = console.Product.Id
	updatedConsole.Product.Image.Id = console.Product.Image.Id
	updatedConsole.Product.ImageId = console.Product.Image.Id
	err = consoleService.IConsoleRepository.Update(updatedConsole)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (consoleService *consoleService) Delete(id int) error {
	console, err := consoleService.GetById(id)
	if err != nil {
		return err
	}
	console.Product.Archived = true
	return consoleService.IConsoleRepository.Update(console)
}