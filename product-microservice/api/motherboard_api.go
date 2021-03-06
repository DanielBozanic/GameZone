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

func (motherboardApi *MotherboardAPI) GetNumberOfRecords(c *gin.Context) {
	numberOfRecords := motherboardApi.IMotherboardService.GetNumberOfRecords()
	c.JSON(http.StatusOK, numberOfRecords)
}

func (motherboardApi *MotherboardAPI) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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

func (motherboardApi *MotherboardAPI) GetNumberOfRecordsSearch(c *gin.Context) {
	numberOfRecords := motherboardApi.IMotherboardService.GetNumberOfRecordsSearch(c.Query("name"))
	c.JSON(http.StatusOK, numberOfRecords)
}

func (motherboardApi *MotherboardAPI) Filter(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	var filter filter.MotherboardFilter
	err = c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	motherboards, err := motherboardApi.IMotherboardService.Filter(page, pageSize, filter)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToMotherboardDTOs(motherboards))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (motherboardApi *MotherboardAPI) GetNumberOfRecordsFilter(c *gin.Context) {
	var filter filter.MotherboardFilter
	err := c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	numberOfRecords := motherboardApi.IMotherboardService.GetNumberOfRecordsFilter(filter)
	c.JSON(http.StatusOK, numberOfRecords)
}

func (motherboardApi *MotherboardAPI) GetManufacturers(c *gin.Context) {
	manufacturers := motherboardApi.IMotherboardService.GetManufacturers()
	c.JSON(http.StatusOK, manufacturers)
}

func (motherboardApi *MotherboardAPI) GetProcessorTypes(c *gin.Context) {
	processorTypes := motherboardApi.IMotherboardService.GetProcessorTypes()
	c.JSON(http.StatusOK, processorTypes)
}

func (motherboardApi *MotherboardAPI) GetSockets(c *gin.Context) {
	sockets := motherboardApi.IMotherboardService.GetSockets()
	c.JSON(http.StatusOK, sockets)
}

func (motherboardApi *MotherboardAPI) GetFormFactors(c *gin.Context) {
	formFactors := motherboardApi.IMotherboardService.GetFormFactors()
	c.JSON(http.StatusOK, formFactors)
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

	msg := motherboardApi.IMotherboardService.Create(motherboard)

	if msg == "" {
		c.JSON(http.StatusOK, "Motherboard added successfully.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (motherboardApi *MotherboardAPI) Update(c *gin.Context) {
	var motherboardDTO dto.MotherboardDTO
	err := c.BindJSON(&motherboardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := motherboardApi.IMotherboardService.Update(motherboardDTO)

	if msg == "" {
		c.JSON(http.StatusOK, "Motherboard updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}

func (motherboardApi *MotherboardAPI) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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