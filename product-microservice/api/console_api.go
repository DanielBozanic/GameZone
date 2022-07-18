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


type ConsoleAPI struct {
	IConsoleService service.IConsoleService
}

func NewConsoleAPI(consoleService service.IConsoleService) ConsoleAPI {
	return ConsoleAPI{IConsoleService: consoleService}
}

func (consoleApi *ConsoleAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	consoles := consoleApi.IConsoleService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToConsoleDTOs(consoles))
}

func (consoleApi *ConsoleAPI) GetNumberOfRecords(c *gin.Context) {
	numberOfRecords := consoleApi.IConsoleService.GetNumberOfRecords()
	c.JSON(http.StatusOK, numberOfRecords)
}

func (consoleApi *ConsoleAPI) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	console, err := consoleApi.IConsoleService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToConsoleDTO(console))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (consoleApi *ConsoleAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	consoles, err := consoleApi.IConsoleService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToConsoleDTOs(consoles))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (consoleApi *ConsoleAPI) GetNumberOfRecordsSearch(c *gin.Context) {
	numberOfRecords := consoleApi.IConsoleService.GetNumberOfRecordsSearch(c.Query("name"))
	c.JSON(http.StatusOK, numberOfRecords)
}

func (consoleApi *ConsoleAPI) Filter(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	var filter filter.ConsoleFilter
	err = c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	consoles, err := consoleApi.IConsoleService.Filter(page, pageSize, filter)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToConsoleDTOs(consoles))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (consoleApi *ConsoleAPI) GetNumberOfRecordsFilter(c *gin.Context) {
	var filter filter.ConsoleFilter
	err := c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	numberOfRecords := consoleApi.IConsoleService.GetNumberOfRecordsFilter(filter)
	c.JSON(http.StatusOK, numberOfRecords)
}

func (consoleApi *ConsoleAPI) GetPlatforms(c *gin.Context) {
	platforms := consoleApi.IConsoleService.GetPlatforms()
	c.JSON(http.StatusOK, platforms)
}

func (consoleApi *ConsoleAPI) Create(c *gin.Context) {
	var consoleDTO dto.ConsoleDTO
	err := c.BindJSON(&consoleDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	console := mapper.ToConsole(consoleDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := consoleApi.IConsoleService.Create(console)

	if msg == "" {
		c.JSON(http.StatusOK, "Console added successfully.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (consoleApi *ConsoleAPI) Update(c *gin.Context) {
	var consoleDTO dto.ConsoleDTO
	err := c.BindJSON(&consoleDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := consoleApi.IConsoleService.Update(consoleDTO)

	if msg == "" {
		c.JSON(http.StatusOK, "Console updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}

func (consoleApi *ConsoleAPI) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := consoleApi.IConsoleService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, "Console deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}