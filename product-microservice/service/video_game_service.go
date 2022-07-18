package service

import (
	"errors"
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/go-sql-driver/mysql"
)


type videoGameService struct {
	IVideoGameRepository repository.IVideoGameRepository
}

type IVideoGameService interface {
	GetAll(page int, pageSize int) ([]model.VideoGame)
	GetNumberOfRecords() int64
	GetById(id int) (model.VideoGame, error)
	SearchByName(page int, pageSize int, name string) ([]model.VideoGame, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.VideoGameFilter) ([]model.VideoGame, error)
	GetNumberOfRecordsFilter(filter filter.VideoGameFilter) int64
	GetPlatforms() []string
	GetGenres() []string
	Create(videoGame model.VideoGame) string
	Update(videoGameDTO dto.VideoGameDTO) string
	Delete(id int) error
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

func (videoGameService *videoGameService) GetById(id int) (model.VideoGame, error) {
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

func (videoGameService *videoGameService) Create(game model.VideoGame) string {
	msg := ""
	game.Product.Type = model.VIDEO_GAME
	err := videoGameService.IVideoGameRepository.Create(game)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (videoGameService *videoGameService) Update(videoGameDTO dto.VideoGameDTO) string {
	msg := ""
	videoGame, err := videoGameService.GetById(videoGameDTO.Product.Id)
	if err != nil {
		return err.Error()
	}

	updatedVideoGame := mapper.ToVideoGame(videoGameDTO)

	updatedVideoGame.Product.Id = videoGame.Product.Id
	updatedVideoGame.ProductId = videoGame.Product.Id
	updatedVideoGame.Product.Image.Id = videoGame.Product.Image.Id
	updatedVideoGame.Product.ImageId = videoGame.Product.Image.Id
	err = videoGameService.IVideoGameRepository.Update(updatedVideoGame)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (videoGameService *videoGameService) Delete(id int) error {
	videoGame, err := videoGameService.GetById(id)
	if err != nil {
		return err
	}
	videoGame.Product.Archived = true
	return videoGameService.IVideoGameRepository.Update(videoGame)
}