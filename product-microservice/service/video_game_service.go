package service

import (
	"product/dto"
	"product/model"
	"product/repository"
	"time"

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

func NewVideoGameService(IVideoGameService repository.IVideoGameRepository) IVideoGameService {
	return &videoGameService{IVideoGameRepository: IVideoGameService}
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
	game.Id = uuid.New()
	return videoGameService.IVideoGameRepository.Create(game)
}

func (videoGameService *videoGameService) Update(videoGameDTO dto.VideoGameDTO) error {
	releaseDate, error := time.Parse("2006-01-02", videoGameDTO.ReleaseDate)
	if error != nil {
		return error
	}

	videoGame, err := videoGameService.GetById(videoGameDTO.Id)
	if err != nil {
		return err
	}

	videoGame.Name = videoGameDTO.Name
	videoGame.Price = videoGameDTO.Price
	videoGame.Amount = videoGameDTO.Amount
	videoGame.Genre = videoGameDTO.Genre
	videoGame.Rating = videoGameDTO.Rating
	videoGame.ReleaseDate = releaseDate
	videoGame.Publisher = videoGameDTO.Publisher
	return videoGameService.IVideoGameRepository.Update(videoGame)
}

func (videoGameService *videoGameService) Delete(id uuid.UUID) error {
	videoGame, err := videoGameService.GetById(id)
	if err != nil {
		return err
	}
	return videoGameService.IVideoGameRepository.Delete(videoGame)
}