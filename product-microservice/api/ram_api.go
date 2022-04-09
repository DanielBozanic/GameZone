package api

import (
	"product/dto"
	"product/mapper"
	"product/service"

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
	rams := ramApi.IRamService.GetAll()
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

func (ramApi *RamAPI) GetByName(c *gin.Context) {
	ram, err := ramApi.IRamService.GetByName(c.Param("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"ram": mapper.ToRamDTO(ram)})
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