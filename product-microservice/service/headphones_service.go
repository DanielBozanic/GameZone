package service

import (
	"product/dto"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type headphonesService struct {
	IHeadphonesRepository repository.IHeadphonesRepository
}

type IHeadphonesService interface {
	GetAll() ([] model.Headphones)
	GetById(id uuid.UUID) (model.Headphones, error)
	GetByName(name string) (model.Headphones, error)
	Create(headphones model.Headphones) error
	Update(headphonesDTO dto.HeadphonesDTO) error
	Delete(id uuid.UUID) error
}

func NewHeadphonesService(headphonesRepository repository.IHeadphonesRepository) IHeadphonesService {
	return &headphonesService{IHeadphonesRepository: headphonesRepository}
}

func (headphonesService *headphonesService) GetAll() []model.Headphones {
	return headphonesService.IHeadphonesRepository.GetAll()
}

func (headphonesService *headphonesService) GetById(id uuid.UUID) (model.Headphones, error) {
	return headphonesService.IHeadphonesRepository.GetById(id)
}

func (headphonesService *headphonesService) GetByName(name string) (model.Headphones, error) {
	return headphonesService.IHeadphonesRepository.GetByName(name)
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