package repository

import (
	"product/dto"
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type videoGameRepository struct {
	Database *gorm.DB
}

type IVideoGameRepository interface {
	GetAll(page int, pageSize int) ([] model.VideoGame)
	GetById(id uuid.UUID) (model.VideoGame, error)
	SearchByName(page int, pageSize int, name string) ([]model.VideoGame, error)
	Filter(page int, pageSize int, filter dto.VideoGameFilter) ([]model.VideoGame, error)
	GetPlatforms() []string
	GetGenres() []string
	Create(videoGame model.VideoGame) error
	Update(videoGame model.VideoGame) error
	Delete(videoGame model.VideoGame) error
}

func NewVideoGameRepository(DB *gorm.DB) IVideoGameRepository {
	return &videoGameRepository{Database: DB}
}

func (videoGameRepo *videoGameRepository) GetAll(page int, pageSize int) []model.VideoGame {
	var games []model.VideoGame
	offset := (page - 1) * pageSize
	videoGameRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&games)
	return games
}

func (videoGameRepo *videoGameRepository) GetById(id uuid.UUID) (model.VideoGame, error) {
	var game model.VideoGame
	result := videoGameRepo.Database.Preload("Product").First(&game, id)
	return game, result.Error
}

func (videoGameRepo *videoGameRepository) SearchByName(page int, pageSize int, name string) ([]model.VideoGame, error) {
	var games []model.VideoGame
	offset := (page - 1) * pageSize
	result := videoGameRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = video_games.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&games)
	return games, result.Error
}

func (videoGameRepo *videoGameRepository) Filter(page int, pageSize int, filter dto.VideoGameFilter) ([]model.VideoGame, error) {
	var games []model.VideoGame
	offset := (page - 1) * pageSize
	result := videoGameRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Where("(platform IN ? OR ?) AND (genre IN ? OR ?)", 
			filter.Platforms, 
			len(filter.Platforms) == 0, 
			filter.Genres, 
			len(filter.Genres) == 0).
		Find(&games)
	return games, result.Error
}

func (videoGameRepo *videoGameRepository) GetPlatforms() []string {
	var platforms []string
	videoGameRepo.Database.
		Model(&model.VideoGame{}).
		Distinct().
		Pluck("platform", &platforms)
	return platforms
}

func (videoGameRepo *videoGameRepository) GetGenres() []string {
	var genres []string
	videoGameRepo.Database.
		Model(&model.VideoGame{}).
		Distinct().
		Pluck("genre", &genres)
	return genres
}

func (videoGameRepo *videoGameRepository) Create(game model.VideoGame) error {
	result := videoGameRepo.Database.Create(&game)
	return result.Error
}

func (videoGameRepo *videoGameRepository) Update(game model.VideoGame) error {
	result := videoGameRepo.Database.Save(&game)
	return result.Error
}

func (videoGameRepo *videoGameRepository) Delete(game model.VideoGame) error {
	result := videoGameRepo.Database.Delete(&game)
	return result.Error
}