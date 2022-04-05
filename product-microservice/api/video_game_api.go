package api

import (
	"product/dto"
	"product/mapper"
	"product/model"
	"product/service"
	"time"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type VideoGameAPI struct {
	IVideoGameService service.IVideoGameService
}

func NewVideoGameAPI(videoGameService service.IVideoGameService) VideoGameAPI {
	return VideoGameAPI{IVideoGameService: videoGameService}
}

func (videoGameApi *VideoGameAPI) GetAll(c *gin.Context) {
	videoGames := videoGameApi.IVideoGameService.GetAll()
	c.JSON(http.StatusOK, gin.H{"video_games": mapper.ToVideoGameDTOs(videoGames)})
}

func (videoGameApi *VideoGameAPI) GetByID(c *gin.Context) {
	videoGame := videoGameApi.IVideoGameService.GetById(uuid.Must(uuid.Parse(c.Param("id"))))
	
	c.JSON(http.StatusOK, gin.H{"video_game": mapper.ToVideoGameDTO(videoGame)})
}

func (videoGameApi *VideoGameAPI) GetByName(c *gin.Context) {
	videoGame := videoGameApi.IVideoGameService.GetByName(c.Param("name"))
	
	c.JSON(http.StatusOK, gin.H{"video_game": mapper.ToVideoGameDTO(videoGame)})
}

func (videoGameApi *VideoGameAPI) Create(c *gin.Context) {
	var videoGameDTO dto.VideoGameDTO
	err := c.BindJSON(&videoGameDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	videoGameApi.IVideoGameService.Save(mapper.ToVideoGame(videoGameDTO))

	c.JSON(http.StatusOK, gin.H{"msg": "Video game stored successfully."})
}

func (videoGameApi *VideoGameAPI) Update(c *gin.Context) {
	var videoGameDTO dto.VideoGameDTO
	err := c.BindJSON(&videoGameDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	videoGame := videoGameApi.IVideoGameService.GetById(videoGameDTO.Id)
	if videoGame == (model.VideoGame{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	releaseDate, error := time.Parse("2006-01-02", videoGameDTO.ReleaseDate)
	if error != nil {
		panic(error)
	}
	
	videoGame.Name = videoGameDTO.Name
	videoGame.Price = videoGameDTO.Price
	videoGame.Amount = videoGameDTO.Amount
	videoGame.Genre = videoGameDTO.Genre
	videoGame.Rating = videoGameDTO.Rating
	videoGame.ReleaseDate = releaseDate
	videoGame.Publisher = videoGameDTO.Publisher
	videoGameApi.IVideoGameService.Update(videoGame)

	c.JSON(http.StatusOK, gin.H{"msg": "Video game updated successfully."})
}

func (videoGameApi *VideoGameAPI) Delete(c *gin.Context) {
	videoGameApi.IVideoGameService.Delete(uuid.Must(uuid.Parse(c.Param("id"))))
	c.JSON(http.StatusOK, gin.H{"msg": "Video game deleted successfully."})
}

