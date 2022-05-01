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


type GraphicsCardAPI struct {
	IGraphicsCardService service.IGraphicsCardService
}

func NewGraphicsCardAPI(graphicsCardService service.IGraphicsCardService) GraphicsCardAPI {
	return GraphicsCardAPI{IGraphicsCardService: graphicsCardService}
}

func (graphicsCardApi *GraphicsCardAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }
	
	graphicsCards := graphicsCardApi.IGraphicsCardService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, gin.H{"graphics_cards": mapper.ToGraphicsCardDTOs(graphicsCards)})
}

func (graphicsCardApi *GraphicsCardAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	graphicsCard, err := graphicsCardApi.IGraphicsCardService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"graphics_card": mapper.ToGraphicsCardDTO(graphicsCard)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (graphicsCardApi *GraphicsCardAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }

	graphicsCards, err := graphicsCardApi.IGraphicsCardService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"graphics_cards": mapper.ToGraphicsCardDTOs(graphicsCards)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (graphicsCardApi *GraphicsCardAPI) Create(c *gin.Context) {
	var graphicsCardDTO dto.GraphicsCardDTO
	err := c.BindJSON(&graphicsCardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	graphicsCard := mapper.ToGraphicsCard(graphicsCardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := graphicsCardApi.IGraphicsCardService.Create(graphicsCard)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Graphics card stored successfully."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}

func (graphicsCardApi *GraphicsCardAPI) Update(c *gin.Context) {
	var graphicsCardDTO dto.GraphicsCardDTO
	err := c.BindJSON(&graphicsCardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := graphicsCardApi.IGraphicsCardService.Update(graphicsCardDTO)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Graphics card updated successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	} 
}

func (graphicsCardApi *GraphicsCardAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	error := graphicsCardApi.IGraphicsCardService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Graphics card deleted successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}