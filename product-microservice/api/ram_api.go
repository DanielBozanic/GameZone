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


type RamAPI struct {
	IRamService service.IRamService
}

func NewRamAPI(ramService service.IRamService) RamAPI {
	return RamAPI{IRamService: ramService}
}

func (ramApi *RamAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }
	
	rams := ramApi.IRamService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, gin.H{"rams": mapper.ToRamDTOs(rams)})
}

func (ramApi *RamAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ram, err := ramApi.IRamService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"ram": mapper.ToRamDTO(ram)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (ramApi *RamAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }

	rams, err := ramApi.IRamService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"rams": mapper.ToRamDTOs(rams)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (ramApi *RamAPI) Create(c *gin.Context) {
	var ramDTO dto.RamDTO
	err := c.BindJSON(&ramDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ram := mapper.ToRam(ramDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := ramApi.IRamService.Create(ram)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "RAM stored successfully."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}

func (ramApi *RamAPI) Update(c *gin.Context) {
	var ramDTO dto.RamDTO
	err := c.BindJSON(&ramDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := ramApi.IRamService.Update(ramDTO)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "RAM updated successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	} 
}

func (ramApi *RamAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	error := ramApi.IRamService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "RAM deleted successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}