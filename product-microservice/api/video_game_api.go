package api

import (
	"product/dto"
	"product/mapper"
	"product/service"
	"strconv"

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
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }
	videoGames := videoGameApi.IVideoGameService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, gin.H{"video_games": mapper.ToVideoGameDTOs(videoGames)})
}

func (videoGameApi *VideoGameAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	videoGame, err := videoGameApi.IVideoGameService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"video_game": mapper.ToVideoGameDTO(videoGame)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (videoGameApi *VideoGameAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }
	videoGames, err := videoGameApi.IVideoGameService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"video_games": mapper.ToVideoGameDTOs(videoGames)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (videoGameApi *VideoGameAPI) Create(c *gin.Context) {
	var videoGameDTO dto.VideoGameDTO
	err := c.BindJSON(&videoGameDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	videoGame, err := mapper.ToVideoGame(videoGameDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := videoGameApi.IVideoGameService.Create(videoGame)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Video game stored successfully."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}

func (videoGameApi *VideoGameAPI) Update(c *gin.Context) {
	var videoGameDTO dto.VideoGameDTO
	err := c.BindJSON(&videoGameDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := videoGameApi.IVideoGameService.Update(videoGameDTO)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Video game updated successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	} 
}

func (videoGameApi *VideoGameAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	error := videoGameApi.IVideoGameService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Video game deleted successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}