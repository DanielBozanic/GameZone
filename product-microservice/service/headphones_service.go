package service

import (
	"errors"
	"fmt"
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/go-sql-driver/mysql"
)

type headphonesService struct {
	IHeadphonesRepository repository.IHeadphonesRepository
}

type IHeadphonesService interface {
	GetAll(page int, pageSize int) ([] model.Headphones)
	GetNumberOfRecords() int64
	GetById(id int) (model.Headphones, error)
	SearchByName(page int, pageSize int, name string) ([]model.Headphones, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.HeadphonesFilter) ([]model.Headphones, error)
	GetNumberOfRecordsFilter(filter filter.HeadphonesFilter) int64
	GetManufacturers() []string
	GetConnectionTypes() []string
	Create(headphones model.Headphones) string
	Update(headphonesDTO dto.HeadphonesDTO) string
	Delete(id int) error
}

func NewHeadphonesService(headphonesRepository repository.IHeadphonesRepository) IHeadphonesService {
	return &headphonesService{IHeadphonesRepository: headphonesRepository}
}

func (headphonesService *headphonesService) GetAll(page int, pageSize int) []model.Headphones {
	return headphonesService.IHeadphonesRepository.GetAll(page, pageSize)
}

func (headphonesService *headphonesService) GetNumberOfRecords() int64 {
	return headphonesService.IHeadphonesRepository.GetNumberOfRecords()
}

func (headphonesService *headphonesService) GetById(id int) (model.Headphones, error) {
	return headphonesService.IHeadphonesRepository.GetById(id)
}

func (headphonesService *headphonesService) SearchByName(page int, pageSize int, name string) ([]model.Headphones, error) {
	return headphonesService.IHeadphonesRepository.SearchByName(page, pageSize, name)
}

func (headphonesService *headphonesService) GetNumberOfRecordsSearch(name string) int64 {
	return headphonesService.IHeadphonesRepository.GetNumberOfRecordsSearch(name)
}

func (headphonesService *headphonesService) Filter(page int, pageSize int, filter filter.HeadphonesFilter) ([]model.Headphones, error) {
	return headphonesService.IHeadphonesRepository.Filter(page, pageSize, filter)
}

func (headphonesService *headphonesService) GetNumberOfRecordsFilter(filter filter.HeadphonesFilter) int64 {
	return headphonesService.IHeadphonesRepository.GetNumberOfRecordsFilter(filter)
}

func (headphonesService *headphonesService) GetManufacturers() []string {
	return headphonesService.IHeadphonesRepository.GetManufacturers()
}

func (headphonesService *headphonesService) GetConnectionTypes() []string {
	return headphonesService.IHeadphonesRepository.GetConnectionTypes()
}

func (headphonesService *headphonesService) Create(headphones model.Headphones) string {
	msg := ""
	headphones.Product.Type = model.HEADPHONES
	err := headphonesService.IHeadphonesRepository.Create(headphones)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (headphonesService *headphonesService) Update(headphonesDTO dto.HeadphonesDTO) string {
	msg := ""
	headphones, err := headphonesService.GetById(headphonesDTO.Product.Id)
	if err != nil {
		return err.Error()
	}
	updatedHeadphones := mapper.ToHeadphones(headphonesDTO)
	updatedHeadphones.Product.Id = headphones.Product.Id
	updatedHeadphones.ProductId = headphones.Product.Id
	updatedHeadphones.Product.Image.Id = headphones.Product.Image.Id
	updatedHeadphones.Product.ImageId = headphones.Product.Image.Id
	fmt.Println(updatedHeadphones.Wireless)
	err = headphonesService.IHeadphonesRepository.Update(updatedHeadphones)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (headphonesService *headphonesService) Delete(id int) error {
	headphones, err := headphonesService.GetById(id)
	if err != nil {
		return err
	}
	headphones.Product.Archived = true
	return headphonesService.IHeadphonesRepository.Update(headphones)
}