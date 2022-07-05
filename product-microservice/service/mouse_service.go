package service

import (
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type mouseService struct {
	IMouseRepository repository.IMouseRepository
}

type IMouseService interface {
	GetAll(page int, pageSize int) ([] model.Mouse)
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.Mouse, error)
	SearchByName(page int, pageSize int, name string) ([]model.Mouse, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.MouseFilter) ([]model.Mouse, error)
	GetNumberOfRecordsFilter(filter filter.MouseFilter) int64
	GetManufacturers() []string
	GetDPIs() []string
	GetConnections() []string
	Create(mouse model.Mouse) error
	Update(mouseDTO dto.MouseDTO) error
	Delete(id uuid.UUID) error
}

func NewMouseService(mouseRepository repository.IMouseRepository) IMouseService {
	return &mouseService{IMouseRepository: mouseRepository}
}

func (mouseService *mouseService) GetAll(page int, pageSize int) []model.Mouse {
	return mouseService.IMouseRepository.GetAll(page, pageSize)
}

func (mouseService *mouseService) GetNumberOfRecords() int64 {
	return mouseService.IMouseRepository.GetNumberOfRecords()
}

func (mouseService *mouseService) GetById(id uuid.UUID) (model.Mouse, error) {
	return mouseService.IMouseRepository.GetById(id)
}

func (mouseService *mouseService) SearchByName(page int, pageSize int, name string) ([]model.Mouse, error) {
	return mouseService.IMouseRepository.SearchByName(page, pageSize, name)
}

func (mouseService *mouseService) GetNumberOfRecordsSearch(name string) int64 {
	return mouseService.IMouseRepository.GetNumberOfRecordsSearch(name)
}

func (mouseService *mouseService) Filter(page int, pageSize int, filter filter.MouseFilter) ([]model.Mouse, error) {
	return mouseService.IMouseRepository.Filter(page, pageSize, filter)
}

func (mouseService *mouseService) GetNumberOfRecordsFilter(filter filter.MouseFilter) int64 {
	return mouseService.IMouseRepository.GetNumberOfRecordsFilter(filter)
}

func (mouseService *mouseService) GetManufacturers() []string {
	return mouseService.IMouseRepository.GetManufacturers()
}

func (mouseService *mouseService) GetDPIs() []string {
	return mouseService.IMouseRepository.GetDPIs()
}

func (mouseService *mouseService) GetConnections() []string {
	return mouseService.IMouseRepository.GetConnections()
}

func (mouseService *mouseService) Create(mouse model.Mouse) error {
	mouse.Product.Id = uuid.New()
	mouse.ProductId = mouse.Product.Id
	mouse.Product.Type = model.MOUSE
	return mouseService.IMouseRepository.Create(mouse)
}

func (mouseService *mouseService) Update(mouseDTO dto.MouseDTO) error {
	mouse, err := mouseService.GetById(mouseDTO.Product.Id)
	if err != nil {
		return err
	}
	updatedMouse := mapper.ToMouse(mouseDTO)
	updatedMouse.Product.Id = mouse.Product.Id
	updatedMouse.ProductId = mouse.Product.Id
	return mouseService.IMouseRepository.Update(updatedMouse)
}

func (mouseService *mouseService) Delete(id uuid.UUID) error {
	mouse, err := mouseService.GetById(id)
	if err != nil {
		return err
	}
	return mouseService.IMouseRepository.Delete(mouse)
}