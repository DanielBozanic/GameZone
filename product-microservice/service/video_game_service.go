package service

import (
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)


type videoGameService struct {
	IVideoGameRepository repository.IVideoGameRepository
}

type IVideoGameService interface {
	GetAll(page int, pageSize int) ([]model.VideoGame)
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.VideoGame, error)
	SearchByName(page int, pageSize int, name string) ([]model.VideoGame, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.VideoGameFilter) ([]model.VideoGame, error)
	GetNumberOfRecordsFilter(filter filter.VideoGameFilter) int64
	GetPlatforms() []string
	GetGenres() []string
	Create(videoGame model.VideoGame) error
	Update(videoGameDTO dto.VideoGameDTO) error
	Delete(id uuid.UUID) error
}

func NewVideoGameService(videoGameRepository repository.IVideoGameRepository) IVideoGameService {
	return &videoGameService{IVideoGameRepository: videoGameRepository}
}

func (videoGameService *videoGameService) GetAll(page int, pageSize int) ([]model.VideoGame) {
	return videoGameService.IVideoGameRepository.GetAll(page, pageSize)
}

func (videoGameService *videoGameService) GetNumberOfRecords() int64 {
	return videoGameService.IVideoGameRepository.GetNumberOfRecords()
}

func (videoGameService *videoGameService) GetById(id uuid.UUID) (model.VideoGame, error) {
	return videoGameService.IVideoGameRepository.GetById(id)
}

func (videoGameService *videoGameService) SearchByName(page int, pageSize int, name string) ([]model.VideoGame, error) {
	return videoGameService.IVideoGameRepository.SearchByName(page, pageSize, name)
}

func (videoGameService *videoGameService) GetNumberOfRecordsSearch(name string) int64 {
	return videoGameService.IVideoGameRepository.GetNumberOfRecordsSearch(name)
}

func (videoGameService *videoGameService) Filter(page int, pageSize int, filter filter.VideoGameFilter) ([]model.VideoGame, error) {
	return videoGameService.IVideoGameRepository.Filter(page, pageSize, filter)
}

func (videoGameService *videoGameService) GetNumberOfRecordsFilter(filter filter.VideoGameFilter) int64 {
	return videoGameService.IVideoGameRepository.GetNumberOfRecordsFilter(filter);
}

func (videoGameService *videoGameService) GetPlatforms() []string {
	return videoGameService.IVideoGameRepository.GetPlatforms()
}

func (videoGameService *videoGameService) GetGenres() []string {
	return videoGameService.IVideoGameRepository.GetGenres()
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