package api

import (
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	graphicsCards := graphicsCardApi.IGraphicsCardService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToGraphicsCardDTOs(graphicsCards))
}

func (graphicsCardApi *GraphicsCardAPI) GetNumberOfRecords(c *gin.Context) {
	numberOfRecords := graphicsCardApi.IGraphicsCardService.GetNumberOfRecords()
	c.JSON(http.StatusOK, numberOfRecords)
}

func (graphicsCardApi *GraphicsCardAPI) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	graphicsCard, err := graphicsCardApi.IGraphicsCardService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToGraphicsCardDTO(graphicsCard))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (graphicsCardApi *GraphicsCardAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	graphicsCards, err := graphicsCardApi.IGraphicsCardService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToGraphicsCardDTOs(graphicsCards))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (graphicsCardApi *GraphicsCardAPI) GetNumberOfRecordsSearch(c *gin.Context) {
	numberOfRecords := graphicsCardApi.IGraphicsCardService.GetNumberOfRecordsSearch(c.Query("name"))
	c.JSON(http.StatusOK, numberOfRecords)
}

func (graphicsCardApi *GraphicsCardAPI) Filter(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	var filter filter.GraphicsCardFilter
	err = c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	graphicsCards, err := graphicsCardApi.IGraphicsCardService.Filter(page, pageSize, filter)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToGraphicsCardDTOs(graphicsCards))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (graphicsCardApi *GraphicsCardAPI) GetNumberOfRecordsFilter(c *gin.Context) {
	var filter filter.GraphicsCardFilter
	err := c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	numberOfRecords := graphicsCardApi.IGraphicsCardService.GetNumberOfRecordsFilter(filter)
	c.JSON(http.StatusOK, numberOfRecords)
}

func (graphicsCardApi *GraphicsCardAPI) GetManufacturers(c *gin.Context) {
	manufacturers := graphicsCardApi.IGraphicsCardService.GetManufacturers()
	c.JSON(http.StatusOK, manufacturers)
}

func (graphicsCardApi *GraphicsCardAPI) GetChipManufacturers(c *gin.Context) {
	chipManufacturers := graphicsCardApi.IGraphicsCardService.GetChipManufacturers()
	c.JSON(http.StatusOK, chipManufacturers)
}

func (graphicsCardApi *GraphicsCardAPI) GetMemorySizes(c *gin.Context) {
	memorySizes := graphicsCardApi.IGraphicsCardService.GetMemorySizes()
	c.JSON(http.StatusOK, memorySizes)
}

func (graphicsCardApi *GraphicsCardAPI) GetMemoryTypes(c *gin.Context) {
	memoryTypes := graphicsCardApi.IGraphicsCardService.GetMemoryTypes()
	c.JSON(http.StatusOK, memoryTypes)
}

func (graphicsCardApi *GraphicsCardAPI) GetModelNames(c *gin.Context) {
	modelNames := graphicsCardApi.IGraphicsCardService.GetModelNames()
	c.JSON(http.StatusOK, modelNames)
}

func (graphicsCardApi *GraphicsCardAPI) Create(c *gin.Context) {
	var graphicsCardDTO dto.GraphicsCardDTO
	err := c.BindJSON(&graphicsCardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	graphicsCard := mapper.ToGraphicsCard(graphicsCardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := graphicsCardApi.IGraphicsCardService.Create(graphicsCard)

	if msg == "" {
		c.JSON(http.StatusOK, "Graphics card added successfully.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (graphicsCardApi *GraphicsCardAPI) Update(c *gin.Context) {
	var graphicsCardDTO dto.GraphicsCardDTO
	err := c.BindJSON(&graphicsCardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := graphicsCardApi.IGraphicsCardService.Update(graphicsCardDTO)

	if msg == "" {
		c.JSON(http.StatusOK, "Graphics card updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}

func (graphicsCardApi *GraphicsCardAPI) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := graphicsCardApi.IGraphicsCardService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, "Graphics card deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}