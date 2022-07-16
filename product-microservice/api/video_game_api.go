package api

import (
	"product/dto"
	"product/dto/filter"
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
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	videoGames := videoGameApi.IVideoGameService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToVideoGameDTOs(videoGames))
}

func (videoGameApi *VideoGameAPI) GetNumberOfRecords(c *gin.Context) {
	numberOfRecords := videoGameApi.IVideoGameService.GetNumberOfRecords()
	c.JSON(http.StatusOK, numberOfRecords)
}

func (videoGameApi *VideoGameAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	videoGame, err := videoGameApi.IVideoGameService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToVideoGameDTO(videoGame))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (videoGameApi *VideoGameAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	videoGames, err := videoGameApi.IVideoGameService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToVideoGameDTOs(videoGames))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (videoGameApi *VideoGameAPI) GetNumberOfRecordsSearch(c *gin.Context) {
	numberOfRecords := videoGameApi.IVideoGameService.GetNumberOfRecordsSearch(c.Query("name"))
	c.JSON(http.StatusOK, numberOfRecords)
}

func (videoGameApi *VideoGameAPI) Filter(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	var filter filter.VideoGameFilter
	err = c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	videoGames, err := videoGameApi.IVideoGameService.Filter(page, pageSize, filter)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToVideoGameDTOs(videoGames))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (videoGameApi *VideoGameAPI) GetNumberOfRecordsFilter(c *gin.Context) {
	var filter filter.VideoGameFilter
	err := c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	numberOfRecords := videoGameApi.IVideoGameService.GetNumberOfRecordsFilter(filter)
	c.JSON(http.StatusOK, numberOfRecords)
}

func (videoGameApi *VideoGameAPI) GetPlatforms(c *gin.Context) {
	platforms := videoGameApi.IVideoGameService.GetPlatforms()
	c.JSON(http.StatusOK, platforms)
}

func (videoGameApi *VideoGameAPI) GetGenres(c *gin.Context) {
	genres := videoGameApi.IVideoGameService.GetGenres()
	c.JSON(http.StatusOK, genres)
}

func (videoGameApi *VideoGameAPI) Create(c *gin.Context) {
	var videoGameDTO dto.VideoGameDTO
	err := c.BindJSON(&videoGameDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	videoGame := mapper.ToVideoGame(videoGameDTO)
	msg := videoGameApi.IVideoGameService.Create(videoGame)

	if msg == "" {
		c.JSON(http.StatusOK, "Video game added successfully.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (videoGameApi *VideoGameAPI) Update(c *gin.Context) {
	var videoGameDTO dto.VideoGameDTO
	err := c.BindJSON(&videoGameDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := videoGameApi.IVideoGameService.Update(videoGameDTO)

	if msg == "" {
		c.JSON(http.StatusOK, "Video game updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}

func (videoGameApi *VideoGameAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := videoGameApi.IVideoGameService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, "Video game deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}