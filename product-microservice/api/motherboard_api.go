package api

import (
	"product/dto"
	"product/mapper"
	"product/service"

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
	motherboards := motherboardApi.IMotherboardService.GetAll()
	c.JSON(http.StatusOK, gin.H{"motherboards": mapper.ToMotherboardDTOs(motherboards)})
}

func (motherboardApi *MotherboardAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	motherboard, err := motherboardApi.IMotherboardService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"motherboard": mapper.ToMotherboardDTO(motherboard)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (motherboardApi *MotherboardAPI) GetByName(c *gin.Context) {
	motherboard, err := motherboardApi.IMotherboardService.GetByName(c.Param("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"motherboard": mapper.ToMotherboardDTO(motherboard)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (motherboardApi *MotherboardAPI) Create(c *gin.Context) {
	var motherboardDTO dto.MotherboardDTO
	err := c.BindJSON(&motherboardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	motherboard := mapper.ToMotherboard(motherboardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := motherboardApi.IMotherboardService.Create(motherboard)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Motherboard stored successfully."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}

func (motherboardApi *MotherboardAPI) Update(c *gin.Context) {
	var motherboardDTO dto.MotherboardDTO
	err := c.BindJSON(&motherboardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := motherboardApi.IMotherboardService.Update(motherboardDTO)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Motherboard updated successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	} 
}

func (motherboardApi *MotherboardAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	error := motherboardApi.IMotherboardService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Motherboard deleted successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}