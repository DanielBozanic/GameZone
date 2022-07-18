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

type hardDiskDriveService struct {
	IHardDiskDriveRepository repository.IHardDiskDriveRepository
}

type IHardDiskDriveService interface {
	GetAll(page int, pageSize int) ([] model.HardDiskDrive)
	GetNumberOfRecords() int64
	GetById(id int) (model.HardDiskDrive, error)
	SearchByName(page int, pageSize int, name string) ([]model.HardDiskDrive, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.HardDiskDriveFilter) ([]model.HardDiskDrive, error)
	GetNumberOfRecordsFilter(filter filter.HardDiskDriveFilter) int64
	GetCapacities() []string
	GetForms() []string
	GetManufacturers() []string
	GetDiskSpeeds() []string
	Create(hardDiskDrive model.HardDiskDrive) string
	Update(hardDiskDriveDTO dto.HardDiskDriveDTO) string
	Delete(id int) error
}

func NewHardDiskDriveService(hardDiskDriveRepository repository.IHardDiskDriveRepository) IHardDiskDriveService {
	return &hardDiskDriveService{IHardDiskDriveRepository: hardDiskDriveRepository}
}

func (hardDiskDriveService *hardDiskDriveService) GetAll(page int, pageSize int) []model.HardDiskDrive {
	return hardDiskDriveService.IHardDiskDriveRepository.GetAll(page, pageSize)
}

func (hardDiskDriveService *hardDiskDriveService) GetNumberOfRecords() int64 {
	return hardDiskDriveService.IHardDiskDriveRepository.GetNumberOfRecords()
}

func (hardDiskDriveService *hardDiskDriveService) GetById(id int) (model.HardDiskDrive, error) {
	return hardDiskDriveService.IHardDiskDriveRepository.GetById(id)
}

func (hardDiskDriveService *hardDiskDriveService) SearchByName(page int, pageSize int, name string) ([]model.HardDiskDrive, error) {
	return hardDiskDriveService.IHardDiskDriveRepository.SearchByName(page, pageSize, name)
}

func (hardDiskDriveService *hardDiskDriveService) GetNumberOfRecordsSearch(name string) int64 {
	return hardDiskDriveService.IHardDiskDriveRepository.GetNumberOfRecordsSearch(name)
}

func (hardDiskDriveService *hardDiskDriveService) Filter(page int, pageSize int, filter filter.HardDiskDriveFilter) ([]model.HardDiskDrive, error) {
	return hardDiskDriveService.IHardDiskDriveRepository.Filter(page, pageSize, filter)
}

func (hardDiskDriveService *hardDiskDriveService) GetNumberOfRecordsFilter(filter filter.HardDiskDriveFilter) int64 {
	return hardDiskDriveService.IHardDiskDriveRepository.GetNumberOfRecordsFilter(filter)
}

func (hardDiskDriveService *hardDiskDriveService) GetCapacities() []string {
	return hardDiskDriveService.IHardDiskDriveRepository.GetCapacities()
}

func (hardDiskDriveService *hardDiskDriveService) GetForms() []string {
	return hardDiskDriveService.IHardDiskDriveRepository.GetForms()
}

func (hardDiskDriveService *hardDiskDriveService) GetManufacturers() []string {
	return hardDiskDriveService.IHardDiskDriveRepository.GetManufacturers()
}

func (hardDiskDriveService *hardDiskDriveService) GetDiskSpeeds() []string {
	return hardDiskDriveService.IHardDiskDriveRepository.GetDiskSpeeds()
}

func (hardDiskDriveService *hardDiskDriveService) Create(hardDiskDrive model.HardDiskDrive) string {
	msg := ""
	hardDiskDrive.Product.Type = model.HARD_DISK_DRIVE
	err := hardDiskDriveService.IHardDiskDriveRepository.Create(hardDiskDrive)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (hardDiskDriveService *hardDiskDriveService) Update(hardDiskDriveDTO dto.HardDiskDriveDTO) string {
	msg := ""
	hardDiskDrive, err := hardDiskDriveService.GetById(hardDiskDriveDTO.Product.Id)
	if err != nil {
		return err.Error()
	}
	updatedHardDiskDrive := mapper.ToHardDiskDrive(hardDiskDriveDTO)
	updatedHardDiskDrive.Product.Id = hardDiskDrive.Product.Id
	updatedHardDiskDrive.ProductId = hardDiskDrive.Product.Id
	updatedHardDiskDrive.Product.Image.Id = hardDiskDrive.Product.Image.Id
	updatedHardDiskDrive.Product.ImageId = hardDiskDrive.Product.Image.Id
	err = hardDiskDriveService.IHardDiskDriveRepository.Update(updatedHardDiskDrive)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (hardDiskDriveService *hardDiskDriveService) Delete(id int) error {
	hardDiskDrive, err := hardDiskDriveService.GetById(id)
	if err != nil {
		return err
	}
	hardDiskDrive.Product.Archived = true
	return hardDiskDriveService.IHardDiskDriveRepository.Update(hardDiskDrive)
}