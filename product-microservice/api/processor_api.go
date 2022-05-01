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

	error := processorApi.IProcessorService.Create(processor)

	if error == nil {
		c.JSON(http.StatusOK, "Processor stored successfully.")
	} else {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}

func (processorApi *ProcessorAPI) Update(c *gin.Context) {
	var processorDTO dto.ProcessorDTO
	err := c.BindJSON(&processorDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := processorApi.IProcessorService.Update(processorDTO)

	if error == nil {
		c.JSON(http.StatusOK, "Processor updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
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