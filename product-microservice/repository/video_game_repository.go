package repository

import (
	"fmt"
	"product/dto/filter"
	"product/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type videoGameRepository struct {
	Database *gorm.DB
}

type IVideoGameRepository interface {
	GetAll(page int, pageSize int) [] model.VideoGame
	GetNumberOfRecords() int64
	GetById(id int) (model.VideoGame, error)
	SearchByName(page int, pageSize int, name string) ([]model.VideoGame, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.VideoGameFilter) ([]model.VideoGame, error)
	GetNumberOfRecordsFilter(filter filter.VideoGameFilter) int64
	GetPlatforms() []string
	GetGenres() []string
	Create(videoGame model.VideoGame) error
	Update(videoGame model.VideoGame) error
}

func NewVideoGameRepository(DB *gorm.DB) IVideoGameRepository {
	return &videoGameRepository{Database: DB}
}

func (videoGameRepo *videoGameRepository) GetAll(page int, pageSize int) ([]model.VideoGame) {
	var games []model.VideoGame
	offset := (page - 1) * pageSize
	videoGameRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = video_games.product_id").
		Where("products.archived = false").
		Find(&games)
	return games
}

func (videoGameRepo *videoGameRepository) GetNumberOfRecords() int64 {
	var count int64
	videoGameRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = video_games.product_id").
		Where("products.archived = false").
		Model(&model.VideoGame{}).
		Count(&count)
	return count
}

func (videoGameRepo *videoGameRepository) GetById(id int) (model.VideoGame, error) {
	var game model.VideoGame
	result := videoGameRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = video_games.product_id").
		Where("products.archived = false").
		First(&game, id)
	return game, result.Error
}

func (videoGameRepo *videoGameRepository) SearchByName(page int, pageSize int, name string) ([]model.VideoGame, error) {
	var games []model.VideoGame
	offset := (page - 1) * pageSize
	result := videoGameRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = video_games.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Find(&games)
	return games, result.Error
}

func (videoGameRepo *videoGameRepository) GetNumberOfRecordsSearch(name string) int64 {
	var games []model.VideoGame
	var count int64
	videoGameRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = video_games.product_id").
		Where("products.name LIKE ? AND products.archived = false", "%" + name + "%").
		Find(&games).
		Count(&count)
	return count
}

func (videoGameRepo *videoGameRepository) Filter(page int, pageSize int, filter filter.VideoGameFilter) ([]model.VideoGame, error) {
	var games []model.VideoGame
	offset := (page - 1) * pageSize
	result := videoGameRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = video_games.product_id").
		Where("(platform IN ? OR ?) AND (genre IN ? OR ?) AND products.archived = false", 
			filter.Platforms, 
			len(filter.Platforms) == 0, 
			filter.Genres, 
			len(filter.Genres) == 0).
		Find(&games)
	return games, result.Error
}

func (videoGameRepo *videoGameRepository) GetNumberOfRecordsFilter(filter filter.VideoGameFilter) int64 {
	var games []model.VideoGame
	var count int64
	videoGameRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = video_games.product_id").
		Where("(platform IN ? OR ?) AND (genre IN ? OR ?) AND products.archived = false", 
			filter.Platforms, 
			len(filter.Platforms) == 0, 
			filter.Genres, 
			len(filter.Genres) == 0).
		Find(&games).
		Count(&count)
	fmt.Printf("%d", uint64(count))
	return count
}

func (videoGameRepo *videoGameRepository) GetPlatforms() []string {
	var platforms []string
	videoGameRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = video_games.product_id").
		Where("products.archived = false").
		Model(&model.VideoGame{}).
		Distinct().
		Pluck("platform", &platforms)
	return platforms
}

func (videoGameRepo *videoGameRepository) GetGenres() []string {
	var genres []string
	videoGameRepo.Database.
		Preload(clause.Associations).Preload("Product." + clause.Associations).
		Joins("JOIN products ON products.id = video_games.product_id").
		Where("products.archived = false").
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
	result := videoGameRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&game)
	return result.Error
}