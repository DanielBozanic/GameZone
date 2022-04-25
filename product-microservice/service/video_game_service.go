package service

import (
	"product/dto"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)


type videoGameService struct {
	IVideoGameRepository repository.IVideoGameRepository
}

type IVideoGameService interface {
	GetAll() ([] model.VideoGame)
	GetById(id uuid.UUID) (model.VideoGame, error)
	GetByName(name string) ([]model.VideoGame, error)
	Create(videoGame model.VideoGame) error
	Update(videoGameDTO dto.VideoGameDTO) error
	Delete(id uuid.UUID) error
}

func NewVideoGameService(videoGameRepository repository.IVideoGameRepository) IVideoGameService {
	return &videoGameService{IVideoGameRepository: videoGameRepository}
}

func (videoGameService *videoGameService) GetAll() []model.VideoGame {
	return videoGameService.IVideoGameRepository.GetAll()
}

func (videoGameService *videoGameService) GetById(id uuid.UUID) (model.VideoGame, error) {
	return videoGameService.IVideoGameRepository.GetById(id)
}

func (videoGameService *videoGameService) GetByName(name string) ([]model.VideoGame, error) {
	return videoGameService.IVideoGameRepository.GetByName(name)
}

func (videoGameService *videoGameService) Create(game model.VideoGame) error {
	game.Product.Id = uuid.New()
	game.ProductId = game.Product.Id
	game.Product.Type = model.VIDEO_GAME
	return videoGameService.IVideoGameRepository.Create(game)
}

func (videoGameService *videoGameService) Update(videoGameDTO dto.VideoGameDTO) error {
	videoGame, err := videoGameService.GetById(videoGameDTO.Product.Id)
	if err != nil {
		return err
	}

	updatedVideoGame, error := mapper.ToVideoGame(videoGameDTO)
	if error != nil {
		return error
	}
	
	updatedVideoGame.Product.Id = videoGame.Product.Id
	updatedVideoGame.ProductId = videoGame.Product.Id
	return videoGameService.IVideoGameRepository.Update(updatedVideoGame)
}

func (videoGameService *videoGameService) Delete(id uuid.UUID) error {
	videoGame, err := videoGameService.GetById(id)
	if err != nil {
		return err
	}
	return videoGameService.IVideoGameRepository.Delete(videoGame)
}