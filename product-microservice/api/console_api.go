package api

import (
	"product/dto"
	"product/mapper"
	"product/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type ConsoleAPI struct {
	IConsoleService service.IConsoleService
}

func NewConsoleAPI(consoleService service.IConsoleService) ConsoleAPI {
	return ConsoleAPI{IConsoleService: consoleService}
}

func (consoleApi *ConsoleAPI) GetAll(c *gin.Context) {
	consoles := consoleApi.IConsoleService.GetAll()
	c.JSON(http.StatusOK, gin.H{"consoles": mapper.ToConsoleDTOs(consoles)})
}

func (consoleApi *ConsoleAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	console, err := consoleApi.IConsoleService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"console": mapper.ToConsoleDTO(console)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (consoleApi *ConsoleAPI) GetByName(c *gin.Context) {
	console, err := consoleApi.IConsoleService.GetByName(c.Param("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"console": mapper.ToConsoleDTO(console)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (consoleApi *ConsoleAPI) Create(c *gin.Context) {
	var consoleDTO dto.ConsoleDTO
	err := c.BindJSON(&consoleDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	console := mapper.ToConsole(consoleDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := consoleApi.IConsoleService.Create(console)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Console stored successfully."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}

func (consoleApi *ConsoleAPI) Update(c *gin.Context) {
	var consoleDTO dto.ConsoleDTO
	err := c.BindJSON(&consoleDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := consoleApi.IConsoleService.Update(consoleDTO)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Console updated successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	} 
}

func (consoleApi *ConsoleAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	error := consoleApi.IConsoleService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Console deleted successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}