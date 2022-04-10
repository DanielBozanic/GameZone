package api

import (
	"product/dto"
	"product/mapper"
	"product/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type MonitorAPI struct {
	IMonitorService service.IMonitorService
}

func NewMonitorAPI(monitorService service.IMonitorService) MonitorAPI {
	return MonitorAPI{IMonitorService: monitorService}
}

func (monitorApi *MonitorAPI) GetAll(c *gin.Context) {
	monitors := monitorApi.IMonitorService.GetAll()
	c.JSON(http.StatusOK, gin.H{"monitors": mapper.ToMonitorDTOs(monitors)})
}

func (monitorApi *MonitorAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	monitor, err := monitorApi.IMonitorService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"monitor": mapper.ToMonitorDTO(monitor)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (monitorApi *MonitorAPI) GetByName(c *gin.Context) {
	monitor, err := monitorApi.IMonitorService.GetByName(c.Param("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"monitor": mapper.ToMonitorDTO(monitor)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (monitorApi *MonitorAPI) Create(c *gin.Context) {
	var monitorDTO dto.MonitorDTO
	err := c.BindJSON(&monitorDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	monitor := mapper.ToMonitor(monitorDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := monitorApi.IMonitorService.Create(monitor)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Monitor stored successfully."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}

func (monitorApi *MonitorAPI) Update(c *gin.Context) {
	var monitorDTO dto.MonitorDTO
	err := c.BindJSON(&monitorDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := monitorApi.IMonitorService.Update(monitorDTO)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Monitor updated successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	} 
}

func (monitorApi *MonitorAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	error := monitorApi.IMonitorService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Monitor deleted successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}