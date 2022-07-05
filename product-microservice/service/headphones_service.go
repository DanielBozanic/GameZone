package service

import (
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type headphonesService struct {
	IHeadphonesRepository repository.IHeadphonesRepository
}

type IHeadphonesService interface {
	GetAll(page int, pageSize int) ([] model.Headphones)
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.Headphones, error)
	SearchByName(page int, pageSize int, name string) ([]model.Headphones, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.HeadphonesFilter) ([]model.Headphones, error)
	GetNumberOfRecordsFilter(filter filter.HeadphonesFilter) int64
	GetManufacturers() []string
	GetConnectionTypes() []string
	Create(headphones model.Headphones) error
	Update(headphonesDTO dto.HeadphonesDTO) error
	Delete(id uuid.UUID) error
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

func (headphonesService *headphonesService) GetById(id uuid.UUID) (model.Headphones, error) {
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

func (headphonesService *headphonesService) Create(headphones model.Headphones) error {
	headphones.Product.Id = uuid.New()
	headphones.ProductId = headphones.Product.Id
	headphones.Product.Type = model.HEADPHONES
	return headphonesService.IHeadphonesRepository.Create(headphones)
}

func (headphonesService *headphonesService) Update(headphonesDTO dto.HeadphonesDTO) error {
	headphones, err := headphonesService.GetById(headphonesDTO.Product.Id)
	if err != nil {
		return err
	}
	updatedHeadphones := mapper.ToHeadphones(headphonesDTO)
	updatedHeadphones.Product.Id = headphones.Product.Id
	updatedHeadphones.ProductId = updatedHeadphones.Product.Id
	return headphonesService.IHeadphonesRepository.Update(updatedHeadphones)
}

func (headphonesService *headphonesService) Delete(id uuid.UUID) error {
	headphones, err := headphonesService.GetById(id)
	if err != nil {
		return err
	}
	return headphonesService.IHeadphonesRepository.Delete(headphones)
}