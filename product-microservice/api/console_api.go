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

func (consoleApi *ConsoleAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
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

	error := consoleApi.IConsoleService.Create(console)

	if error == nil {
		c.JSON(http.StatusOK, "Console stored successfully.")
	} else {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}

func (consoleApi *ConsoleAPI) Update(c *gin.Context) {
	var consoleDTO dto.ConsoleDTO
	err := c.BindJSON(&consoleDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := consoleApi.IConsoleService.Update(consoleDTO)

	if error == nil {
		c.JSON(http.StatusOK, "Console updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	} 
}

func (consoleApi *ConsoleAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
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