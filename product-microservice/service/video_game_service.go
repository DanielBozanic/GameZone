package service

import (
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)


type videoGameService struct {
	IVideoGameRepository repository.IVideoGameRepository
}

type IVideoGameService interface {
	GetAll() ([] model.VideoGame)
	GetById(id uuid.UUID) model.VideoGame
	GetByName(name string) model.VideoGame
	Save(videoGame model.VideoGame)
	Update(videoGame model.VideoGame)
	Delete(id uuid.UUID)
}

func NewVideoGameService(IVideoGameService repository.IVideoGameRepository) IVideoGameService {
	return &videoGameService{IVideoGameRepository: IVideoGameService}
}

func (videoGameService *videoGameService) GetAll() []model.VideoGame {
	return videoGameService.IVideoGameRepository.GetAll()
}

func (videoGameService *videoGameService) GetById(id uuid.UUID) model.VideoGame {
	return videoGameService.IVideoGameRepository.GetById(id)
}

func (videoGameService *videoGameService) GetByName(name string) model.VideoGame {
	return videoGameService.IVideoGameRepository.GetByName(name)
}

func (videoGameService *videoGameService) Save(game model.VideoGame) {
	game.Id = uuid.New()
	videoGameService.IVideoGameRepository.Save(game)
}

func (videoGameService *videoGameService) Update(game model.VideoGame) {
	videoGameService.IVideoGameRepository.Update(game)
}

func (videoGameService *videoGameService) Delete(id uuid.UUID) {
	videoGameService.IVideoGameRepository.Delete(id)
}