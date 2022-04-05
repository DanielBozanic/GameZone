package repository

import (
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type videoGameRepository struct {
	Database *gorm.DB
}

type IVideoGameRepository interface {
	GetAll() ([] model.VideoGame)
	GetById(id uuid.UUID) model.VideoGame
	GetByName(name string) model.VideoGame
	Save(videoGame model.VideoGame)
	Update(videoGame model.VideoGame)
	Delete(id uuid.UUID)
}

func NewVideoGameRepository(DB *gorm.DB) IVideoGameRepository {
	return &videoGameRepository{Database: DB}
}

func (videoGameRepo *videoGameRepository) GetAll() []model.VideoGame {
	var games []model.VideoGame
	videoGameRepo.Database.Find(&games)
	return games
}

func (videoGameRepo *videoGameRepository) GetById(id uuid.UUID) model.VideoGame {
	var game model.VideoGame
	videoGameRepo.Database.First(&game, id)
	return game
}

func (videoGameRepo *videoGameRepository) GetByName(name string) model.VideoGame {
	var game model.VideoGame
	videoGameRepo.Database.First(&game, "name = ?", name).First(&game)
	return game
}

func (videoGameRepo *videoGameRepository) Save(game model.VideoGame) {
	videoGameRepo.Database.Save(&game)
}

func (videoGameRepo *videoGameRepository) Update(game model.VideoGame) {
	videoGameRepo.Database.Save(&game)
}

func (videoGameRepo *videoGameRepository) Delete(id uuid.UUID) {
	var product model.VideoGame
	videoGameRepo.Database.First(&product, id)
	videoGameRepo.Database.Delete(&product)
}