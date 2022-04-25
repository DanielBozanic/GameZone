package service

import (
	"product/dto"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type mouseService struct {
	IMouseRepository repository.IMouseRepository
}

type IMouseService interface {
	GetAll() ([] model.Mouse)
	GetById(id uuid.UUID) (model.Mouse, error)
	GetByName(name string) (model.Mouse, error)
	Create(mouse model.Mouse) error
	Update(mouseDTO dto.MouseDTO) error
	Delete(id uuid.UUID) error
}

func NewMouseService(mouseRepository repository.IMouseRepository) IMouseService {
	return &mouseService{IMouseRepository: mouseRepository}
}

func (mouseService *mouseService) GetAll() []model.Mouse {
	return mouseService.IMouseRepository.GetAll()
}

func (mouseService *mouseService) GetById(id uuid.UUID) (model.Mouse, error) {
	return mouseService.IMouseRepository.GetById(id)
}

func (mouseService *mouseService) GetByName(name string) (model.Mouse, error) {
	return mouseService.IMouseRepository.GetByName(name)
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