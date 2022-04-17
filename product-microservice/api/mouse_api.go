package api

import (
	"product/dto"
	"product/mapper"
	"product/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type MouseAPI struct {
	IMouseService service.IMouseService
}

func NewMouseAPI(mouseService service.IMouseService) MouseAPI {
	return MouseAPI{IMouseService: mouseService}
}

func (mouseApi *MouseAPI) GetAll(c *gin.Context) {
	mouses := mouseApi.IMouseService.GetAll()
	c.JSON(http.StatusOK, gin.H{"mouses": mapper.ToMouseDTOs(mouses)})
}

func (mouseApi *MouseAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	mouse, err := mouseApi.IMouseService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"mouse": mapper.ToMouseDTO(mouse)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (mouseApi *MouseAPI) GetByName(c *gin.Context) {
	mouse, err := mouseApi.IMouseService.GetByName(c.Param("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"mouse": mapper.ToMouseDTO(mouse)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (mouseApi *MouseAPI) Create(c *gin.Context) {
	var mouseDTO dto.MouseDTO
	err := c.BindJSON(&mouseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mouse := mapper.ToMouse(mouseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := mouseApi.IMouseService.Create(mouse)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "mouse stored successfully."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}

func (mouseApi *MouseAPI) Update(c *gin.Context) {
	var mouseDTO dto.MouseDTO
	err := c.BindJSON(&mouseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := mouseApi.IMouseService.Update(mouseDTO)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Mouse updated successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	} 
}

func (mouseApi *MouseAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	error := mouseApi.IMouseService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Mouse deleted successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}