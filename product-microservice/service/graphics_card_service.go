package service

import (
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type graphicsCardService struct {
	IGraphicsCardRepository repository.IGraphicsCardRepository
}

type IGraphicsCardService interface {
	GetAll(page int, pageSize int) ([] model.GraphicsCard)
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.GraphicsCard, error)
	SearchByName(page int, pageSize int, name string) ([]model.GraphicsCard, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.GraphicsCardFilter) ([]model.GraphicsCard, error)
	GetNumberOfRecordsFilter(filter filter.GraphicsCardFilter) int64
	GetManufacturers() []string
	GetChipManufacturers() []string
	GetMemorySizes() []string
	GetMemoryTypes() []string
	GetModelNames() []string
	Create(graphicsCard model.GraphicsCard) error
	Update(graphicsCardDTO dto.GraphicsCardDTO) error
	Delete(id uuid.UUID) error
}

func NewGraphicsCardService(graphicsCardRepository repository.IGraphicsCardRepository) IGraphicsCardService {
	return &graphicsCardService{IGraphicsCardRepository: graphicsCardRepository}
}

func (graphicsCardService *graphicsCardService) GetAll(page int, pageSize int) []model.GraphicsCard {
	return graphicsCardService.IGraphicsCardRepository.GetAll(page, pageSize)
}

func (graphicsCardService *graphicsCardService) GetNumberOfRecords() int64 {
	return graphicsCardService.IGraphicsCardRepository.GetNumberOfRecords()
}

func (graphicsCardService *graphicsCardService) GetById(id uuid.UUID) (model.GraphicsCard, error) {
	return graphicsCardService.IGraphicsCardRepository.GetById(id)
}

func (graphicsCardService *graphicsCardService) SearchByName(page int, pageSize int, name string) ([]model.GraphicsCard, error) {
	return graphicsCardService.IGraphicsCardRepository.SearchByName(page, pageSize, name)
}

func (graphicsCardService *graphicsCardService) GetNumberOfRecordsSearch(name string) int64 {
	return graphicsCardService.IGraphicsCardRepository.GetNumberOfRecordsSearch(name)
}

func (graphicsCardService *graphicsCardService) Filter(page int, pageSize int, filter filter.GraphicsCardFilter) ([]model.GraphicsCard, error) {
	return graphicsCardService.IGraphicsCardRepository.Filter(page, pageSize, filter)
}

func (graphicsCardService *graphicsCardService) GetNumberOfRecordsFilter(filter filter.GraphicsCardFilter) int64 {
	return graphicsCardService.IGraphicsCardRepository.GetNumberOfRecordsFilter(filter)
}

func (graphicsCardService *graphicsCardService) GetManufacturers() []string {
	return graphicsCardService.IGraphicsCardRepository.GetManufacturers()
}

func (graphicsCardService *graphicsCardService) GetChipManufacturers() []string {
	return graphicsCardService.IGraphicsCardRepository.GetChipManufacturers()
}

func (graphicsCardService *graphicsCardService) GetMemorySizes() []string {
	return graphicsCardService.IGraphicsCardRepository.GetMemorySizes()
}

func (graphicsCardService *graphicsCardService) GetMemoryTypes() []string {
	return graphicsCardService.IGraphicsCardRepository.GetMemoryTypes()
}

func (graphicsCardService *graphicsCardService) GetModelNames() []string {
	return graphicsCardService.IGraphicsCardRepository.GetModelNames()
}

func (graphicsCardService *graphicsCardService) Create(graphicsCard model.GraphicsCard) error {
	graphicsCard.Product.Id = uuid.New()
	graphicsCard.ProductId = graphicsCard.Product.Id
	graphicsCard.Product.Type = model.GRAPHICS_CARD
	return graphicsCardService.IGraphicsCardRepository.Create(graphicsCard)
}

func (graphicsCardService *graphicsCardService) Update(graphicsCardDTO dto.GraphicsCardDTO) error {
	graphicsCard, err := graphicsCardService.GetById(graphicsCardDTO.Product.Id)
	if err != nil {
		return err
	}
	updatedGraphicsCard := mapper.ToGraphicsCard(graphicsCardDTO)
	updatedGraphicsCard.Product.Id = graphicsCard.Product.Id
	updatedGraphicsCard.ProductId = graphicsCard.Product.Id
	return graphicsCardService.IGraphicsCardRepository.Update(updatedGraphicsCard)
}

func (graphicsCardService *graphicsCardService) Delete(id uuid.UUID) error {
	graphicsCard, err := graphicsCardService.GetById(id)
	if err != nil {
		return err
	}
	return graphicsCardService.IGraphicsCardRepository.Delete(graphicsCard)
}