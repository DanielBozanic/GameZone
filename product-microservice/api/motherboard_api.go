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


type MotherboardAPI struct {
	IMotherboardService service.IMotherboardService
}

func NewMotherboardAPI(motherboardService service.IMotherboardService) MotherboardAPI {
	return MotherboardAPI{IMotherboardService: motherboardService}
}

func (motherboardApi *MotherboardAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	motherboards := motherboardApi.IMotherboardService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToMotherboardDTOs(motherboards))
}

func (motherboardApi *MotherboardAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	motherboard, err := motherboardApi.IMotherboardService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToMotherboardDTO(motherboard))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (motherboardApi *MotherboardAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	motherboards, err := motherboardApi.IMotherboardService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToMotherboardDTOs(motherboards))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (motherboardApi *MotherboardAPI) Create(c *gin.Context) {
	var motherboardDTO dto.MotherboardDTO
	err := c.BindJSON(&motherboardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	motherboard := mapper.ToMotherboard(motherboardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := motherboardApi.IMotherboardService.Create(motherboard)

	if error == nil {
		c.JSON(http.StatusOK, "Motherboard stored successfully.")
	} else {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}

func (motherboardApi *MotherboardAPI) Update(c *gin.Context) {
	var motherboardDTO dto.MotherboardDTO
	err := c.BindJSON(&motherboardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := motherboardApi.IMotherboardService.Update(motherboardDTO)

	if error == nil {
		c.JSON(http.StatusOK, "Motherboard updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	} 
}

func (motherboardApi *MotherboardAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := motherboardApi.IMotherboardService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, "Motherboard deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}