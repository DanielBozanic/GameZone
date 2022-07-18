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

type ramService struct {
	IRamRepository repository.IRamRepository
}

type IRamService interface {
	GetAll(page int, pageSize int) ([] model.Ram)
	GetNumberOfRecords() int64
	GetById(id int) (model.Ram, error)
	SearchByName(page int, pageSize int, name string) ([]model.Ram, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.RAMFilter) ([]model.Ram, error)
	GetNumberOfRecordsFilter(filter filter.RAMFilter) int64
	GetManufacturers() []string
	GetCapacities() []string
	GetMemoryTypes() []string
	GetSpeeds() []string
	Create(ram model.Ram) string
	Update(ramDTO dto.RamDTO) string
	Delete(id int) error
}

func NewRamServiceService(ramRepository repository.IRamRepository) IRamService {
	return &ramService{IRamRepository: ramRepository}
}

func (ramService *ramService) GetAll(page int, pageSize int) []model.Ram {
	return ramService.IRamRepository.GetAll(page, pageSize)
}

func (ramService *ramService) GetNumberOfRecords() int64 {
	return ramService.IRamRepository.GetNumberOfRecords()
}

func (ramService *ramService) GetById(id int) (model.Ram, error) {
	return ramService.IRamRepository.GetById(id)
}

func (ramService *ramService) SearchByName(page int, pageSize int, name string) ([]model.Ram, error) {
	return ramService.IRamRepository.SearchByName(page, pageSize, name)
}

func (ramService *ramService) GetNumberOfRecordsSearch(name string) int64 {
	return ramService.IRamRepository.GetNumberOfRecordsSearch(name)
}
 
func (ramService *ramService) Filter(page int, pageSize int, filter filter.RAMFilter) ([]model.Ram, error) {
	return ramService.IRamRepository.Filter(page, pageSize, filter)
}

func (ramService *ramService) GetNumberOfRecordsFilter(filter filter.RAMFilter) int64 {
	return ramService.IRamRepository.GetNumberOfRecordsFilter(filter)
}

func (ramService *ramService) GetManufacturers() []string {
	return ramService.IRamRepository.GetManufacturers()
}

func (ramService *ramService) GetCapacities() []string {
	return ramService.IRamRepository.GetCapacities()
}

func (ramService *ramService) GetMemoryTypes() []string {
	return ramService.IRamRepository.GetMemoryTypes()
}


func (ramService *ramService) GetSpeeds() []string {
	return ramService.IRamRepository.GetSpeeds()
}

func (ramService *ramService) Create(ram model.Ram) string {
	msg := ""
	ram.Product.Type = model.RAM
	err := ramService.IRamRepository.Create(ram)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (ramService *ramService) Update(ramDTO dto.RamDTO) string {
	msg := ""
	ram, err := ramService.GetById(ramDTO.Product.Id)
	if err != nil {
		return err.Error()
	}
	updatedRam := mapper.ToRam(ramDTO)
	updatedRam.Product.Id = ram.Product.Id
	updatedRam.ProductId = ram.Product.Id
	updatedRam.Product.Image.Id = ram.Product.Image.Id
	updatedRam.Product.ImageId = ram.Product.Image.Id
	err =  ramService.IRamRepository.Update(updatedRam)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (ramService *ramService) Delete(id int) error {
	ram, err := ramService.GetById(id)
	if err != nil {
		return err
	}
	*ram.Product.Archived = true
	return ramService.IRamRepository.Update(ram)
}