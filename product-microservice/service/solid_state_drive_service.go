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

type solidStateDriveService struct {
	ISolidStateDriveRepository repository.ISolidStateDriveRepository
}

type ISolidStateDriveService interface {
	GetAll(page int, pageSize int) ([] model.SolidStateDrive)
	GetNumberOfRecords() int64
	GetById(id int) (model.SolidStateDrive, error)
	SearchByName(page int, pageSize int, name string) ([]model.SolidStateDrive, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.SolidStateDriveFilter) ([]model.SolidStateDrive, error)
	GetNumberOfRecordsFilter(filter filter.SolidStateDriveFilter) int64
	GetCapacities() []string
	GetForms() []string
	GetManufacturers() []string
	GetMaxSequentialReads() []string
	GetMaxSequentialWrites() []string
	Create(solidStateDrive model.SolidStateDrive) string
	Update(solidStateDriveDTO dto.SolidStateDriveDTO) string
	Delete(id int) error
}

func NewSolidStateDriveService(solidStateDriveRepository repository.ISolidStateDriveRepository) ISolidStateDriveService {
	return &solidStateDriveService{ISolidStateDriveRepository: solidStateDriveRepository}
}

func (solidStateDriveService *solidStateDriveService) GetAll(page int, pageSize int) []model.SolidStateDrive {
	return solidStateDriveService.ISolidStateDriveRepository.GetAll(page, pageSize)
}

func (solidStateDriveService *solidStateDriveService) GetNumberOfRecords() int64 {
	return solidStateDriveService.ISolidStateDriveRepository.GetNumberOfRecords()
}

func (solidStateDriveService *solidStateDriveService) GetById(id int) (model.SolidStateDrive, error) {
	return solidStateDriveService.ISolidStateDriveRepository.GetById(id)
}

func (solidStateDriveService *solidStateDriveService) SearchByName(page int, pageSize int, name string) ([]model.SolidStateDrive, error) {
	return solidStateDriveService.ISolidStateDriveRepository.SearchByName(page, pageSize, name)
}

func (solidStateDriveService *solidStateDriveService) GetNumberOfRecordsSearch(name string) int64 {
	return solidStateDriveService.ISolidStateDriveRepository.GetNumberOfRecordsSearch(name)
}

func (solidStateDriveService *solidStateDriveService) Filter(page int, pageSize int, filter filter.SolidStateDriveFilter) ([]model.SolidStateDrive, error) {
	return solidStateDriveService.ISolidStateDriveRepository.Filter(page, pageSize, filter)
}

func (solidStateDriveService *solidStateDriveService) GetNumberOfRecordsFilter(filter filter.SolidStateDriveFilter) int64 {
	return solidStateDriveService.ISolidStateDriveRepository.GetNumberOfRecordsFilter(filter)
}

func (solidStateDriveService *solidStateDriveService) GetCapacities() []string {
	return solidStateDriveService.ISolidStateDriveRepository.GetCapacities()
}

func (solidStateDriveService *solidStateDriveService) GetForms() []string {
	return solidStateDriveService.ISolidStateDriveRepository.GetForms()
}

func (solidStateDriveService *solidStateDriveService) GetManufacturers() []string {
	return solidStateDriveService.ISolidStateDriveRepository.GetManufacturers()
}

func (solidStateDriveService *solidStateDriveService) GetMaxSequentialReads() []string {
	return solidStateDriveService.ISolidStateDriveRepository.GetMaxSequentialReads()
}

func (solidStateDriveService *solidStateDriveService) GetMaxSequentialWrites() []string {
	return solidStateDriveService.ISolidStateDriveRepository.GetMaxSequentialWrites()
}

func (solidStateDriveService *solidStateDriveService) Create(solidStateDrive model.SolidStateDrive) string {
	msg := ""
	solidStateDrive.Product.Type = model.SOLID_STATE_DRIVE
	err := solidStateDriveService.ISolidStateDriveRepository.Create(solidStateDrive)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (solidStateDriveService *solidStateDriveService) Update(solidStateDriveDTO dto.SolidStateDriveDTO) string {
	msg := ""
	solidStateDrive, err := solidStateDriveService.GetById(solidStateDriveDTO.Product.Id)
	if err != nil {
		return err.Error()
	}
	updatedSolidStateDrive := mapper.ToSolidStateDrive(solidStateDriveDTO)
	updatedSolidStateDrive.Product.Id = solidStateDrive.Product.Id
	updatedSolidStateDrive.ProductId = solidStateDrive.Product.Id
	updatedSolidStateDrive.Product.Image.Id = solidStateDrive.Product.Image.Id
	updatedSolidStateDrive.Product.ImageId = solidStateDrive.Product.Image.Id
	err = solidStateDriveService.ISolidStateDriveRepository.Update(updatedSolidStateDrive)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (solidStateDriveService *solidStateDriveService) Delete(id int) error {
	solidStateDrive, err := solidStateDriveService.GetById(id)
	if err != nil {
		return err
	}
	*solidStateDrive.Product.Archived = true
	return solidStateDriveService.ISolidStateDriveRepository.Update(solidStateDrive)
}