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


type MonitorAPI struct {
	IMonitorService service.IMonitorService
}

func NewMonitorAPI(monitorService service.IMonitorService) MonitorAPI {
	return MonitorAPI{IMonitorService: monitorService}
}

func (monitorApi *MonitorAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	monitors := monitorApi.IMonitorService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToMonitorDTOs(monitors))
}

func (monitorApi *MonitorAPI) GetNumberOfRecords(c *gin.Context) {
	numberOfRecords := monitorApi.IMonitorService.GetNumberOfRecords()
	c.JSON(http.StatusOK, numberOfRecords)
}

func (monitorApi *MonitorAPI) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	monitor, err := monitorApi.IMonitorService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToMonitorDTO(monitor))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (monitorApi *MonitorAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	monitors, err := monitorApi.IMonitorService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToMonitorDTOs(monitors))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (monitorApi *MonitorAPI) GetNumberOfRecordsSearch(c *gin.Context) {
	numberOfRecords := monitorApi.IMonitorService.GetNumberOfRecordsSearch(c.Query("name"))
	c.JSON(http.StatusOK, numberOfRecords)
}

func (monitorApi *MonitorAPI) Filter(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	var filter filter.MonitorFilter
	err = c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	monitors, err := monitorApi.IMonitorService.Filter(page, pageSize, filter)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToMonitorDTOs(monitors))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (monitorApi *MonitorAPI) GetNumberOfRecordsFilter(c *gin.Context) {
	var filter filter.MonitorFilter
	err := c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	numberOfRecords := monitorApi.IMonitorService.GetNumberOfRecordsFilter(filter)
	c.JSON(http.StatusOK, numberOfRecords)
}

func (monitorApi *MonitorAPI) GetManufacturers(c *gin.Context) {
	manufacturers := monitorApi.IMonitorService.GetManufacturers()
	c.JSON(http.StatusOK, manufacturers)
}

func (monitorApi *MonitorAPI) GetAspectRatios(c *gin.Context) {
	aspectRatios := monitorApi.IMonitorService.GetAspectRatios()
	c.JSON(http.StatusOK, aspectRatios)
}

func (monitorApi *MonitorAPI) GetResolutions(c *gin.Context) {
	resolutions := monitorApi.IMonitorService.GetResolutions()
	c.JSON(http.StatusOK, resolutions)
}

func (monitorApi *MonitorAPI) GetRefreshRates(c *gin.Context) {
	refreshRates := monitorApi.IMonitorService.GetRefreshRates()
	c.JSON(http.StatusOK, refreshRates)
}

func (monitorApi *MonitorAPI) Create(c *gin.Context) {
	var monitorDTO dto.MonitorDTO
	err := c.BindJSON(&monitorDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	monitor := mapper.ToMonitor(monitorDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := monitorApi.IMonitorService.Create(monitor)

	if msg == "" {
		c.JSON(http.StatusOK, "Monitor added successfully.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (monitorApi *MonitorAPI) Update(c *gin.Context) {
	var monitorDTO dto.MonitorDTO
	err := c.BindJSON(&monitorDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := monitorApi.IMonitorService.Update(monitorDTO)

	if msg == "" {
		c.JSON(http.StatusOK, "Monitor updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}

func (monitorApi *MonitorAPI) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := monitorApi.IMonitorService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, "Monitor deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}