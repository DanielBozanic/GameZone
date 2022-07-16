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


type ProcessorAPI struct {
	IProcessorService service.IProcessorService
}

func NewProcessorAPI(processorService service.IProcessorService) ProcessorAPI {
	return ProcessorAPI{IProcessorService: processorService}
}

func (processorApi *ProcessorAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	processors := processorApi.IProcessorService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToProcessorDTOs(processors))
}

func (processorApi *ProcessorAPI) GetNumberOfRecords(c *gin.Context) {
	numberOfRecords := processorApi.IProcessorService.GetNumberOfRecords()
	c.JSON(http.StatusOK, numberOfRecords)
}

func (processorApi *ProcessorAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	processor, err := processorApi.IProcessorService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToProcessorDTO(processor))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (processorApi *ProcessorAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	processors, err := processorApi.IProcessorService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToProcessorDTOs(processors))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (processorApi *ProcessorAPI) GetNumberOfRecordsSearch(c *gin.Context) {
	numberOfRecords := processorApi.IProcessorService.GetNumberOfRecordsSearch(c.Query("name"))
	c.JSON(http.StatusOK, numberOfRecords)
}

func (processorApi *ProcessorAPI) Filter(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	var filter filter.ProcessorFilter
	err = c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	processors, err := processorApi.IProcessorService.Filter(page, pageSize, filter)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToProcessorDTOs(processors))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (processorApi *ProcessorAPI) GetNumberOfRecordsFilter(c *gin.Context) {
	var filter filter.ProcessorFilter
	err := c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	numberOfRecords := processorApi.IProcessorService.GetNumberOfRecordsFilter(filter)
	c.JSON(http.StatusOK, numberOfRecords)
}

func (processorApi *ProcessorAPI) GetManufacturers(c *gin.Context) {
	manufacturers := processorApi.IProcessorService.GetManufacturers()
	c.JSON(http.StatusOK, manufacturers)
}

func (processorApi *ProcessorAPI) GetTypes(c *gin.Context) {
	types := processorApi.IProcessorService.GetTypes()
	c.JSON(http.StatusOK, types)
}

func (processorApi *ProcessorAPI) GetSockets(c *gin.Context) {
	sockets := processorApi.IProcessorService.GetSockets()
	c.JSON(http.StatusOK, sockets)
}

func (processorApi *ProcessorAPI) GetNumberOfCores(c *gin.Context) {
	numberOfCores := processorApi.IProcessorService.GetNumberOfCores()
	c.JSON(http.StatusOK, numberOfCores)
}

func (processorApi *ProcessorAPI) GetThreads(c *gin.Context) {
	threads := processorApi.IProcessorService.GetThreads()
	c.JSON(http.StatusOK, threads)
}

func (processorApi *ProcessorAPI) Create(c *gin.Context) {
	var processorDTO dto.ProcessorDTO
	err := c.BindJSON(&processorDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	processor := mapper.ToProcessor(processorDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := processorApi.IProcessorService.Create(processor)

	if msg == "" {
		c.JSON(http.StatusOK, "Processor added successfully.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (processorApi *ProcessorAPI) Update(c *gin.Context) {
	var processorDTO dto.ProcessorDTO
	err := c.BindJSON(&processorDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := processorApi.IProcessorService.Update(processorDTO)

	if msg == "" {
		c.JSON(http.StatusOK, "Processor updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}

func (processorApi *ProcessorAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := processorApi.IProcessorService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, "Processor deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}