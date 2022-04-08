package api

import (
	"product/dto"
	"product/mapper"
	"product/service"

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
	processors := processorApi.IProcessorService.GetAll()
	c.JSON(http.StatusOK, gin.H{"processors": mapper.ToProcessorDTOs(processors)})
}

func (processorApi *ProcessorAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	processor, err := processorApi.IProcessorService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"processor": mapper.ToProcessorDTO(processor)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (processorApi *ProcessorAPI) GetByName(c *gin.Context) {
	processor, err := processorApi.IProcessorService.GetByName(c.Param("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"processor": mapper.ToProcessorDTO(processor)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (processorApi *ProcessorAPI) Create(c *gin.Context) {
	var processorDTO dto.ProcessorDTO
	err := c.BindJSON(&processorDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	processor := mapper.ToProcessor(processorDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := processorApi.IProcessorService.Create(processor)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Processor stored successfully."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}

func (processorApi *ProcessorAPI) Update(c *gin.Context) {
	var processorDTO dto.ProcessorDTO
	err := c.BindJSON(&processorDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := processorApi.IProcessorService.Update(processorDTO)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Processor updated successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	} 
}

func (processorApi *ProcessorAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	error := processorApi.IProcessorService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Processor deleted successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}