package api

import (
	"product/dto"
	"product/mapper"
	"product/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type PowerSupplyUnitAPI struct {
	IPowerSupplyUnitService service.IPowerSupplyUnitService
}

func NewPowerSupplyUnitAPI(powerSupplyUnitService service.IPowerSupplyUnitService) PowerSupplyUnitAPI {
	return PowerSupplyUnitAPI{IPowerSupplyUnitService: powerSupplyUnitService}
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) GetAll(c *gin.Context) {
	powerSupplyUnits := powerSupplyUnitApi.IPowerSupplyUnitService.GetAll()
	c.JSON(http.StatusOK, gin.H{"psus": mapper.ToPowerSupplyUnitDTOs(powerSupplyUnits)})
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	powerSupplyUnit, err := powerSupplyUnitApi.IPowerSupplyUnitService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"psu": mapper.ToPowerSupplyUnitDTO(powerSupplyUnit)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) GetByName(c *gin.Context) {
	powerSupplyUnit, err := powerSupplyUnitApi.IPowerSupplyUnitService.GetByName(c.Param("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"psu": mapper.ToPowerSupplyUnitDTO(powerSupplyUnit)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) Create(c *gin.Context) {
	var powerSupplyUnitDTO dto.PowerSupplyUnitDTO
	err := c.BindJSON(&powerSupplyUnitDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	powerSupplyUnit := mapper.ToPowerSupplyUnit(powerSupplyUnitDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := powerSupplyUnitApi.IPowerSupplyUnitService.Create(powerSupplyUnit)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "PSU stored successfully."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) Update(c *gin.Context) {
	var powerSupplyUnitDTO dto.PowerSupplyUnitDTO
	err := c.BindJSON(&powerSupplyUnitDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := powerSupplyUnitApi.IPowerSupplyUnitService.Update(powerSupplyUnitDTO)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "PSU updated successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	} 
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	error := powerSupplyUnitApi.IPowerSupplyUnitService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "PSU deleted successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}