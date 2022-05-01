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


type MouseAPI struct {
	IMouseService service.IMouseService
}

func NewMouseAPI(mouseService service.IMouseService) MouseAPI {
	return MouseAPI{IMouseService: mouseService}
}

func (mouseApi *MouseAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	mouses := mouseApi.IMouseService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToMouseDTOs(mouses))
}

func (mouseApi *MouseAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	mouse, err := mouseApi.IMouseService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToMouseDTO(mouse))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (mouseApi *MouseAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	mouse, err := mouseApi.IMouseService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToMouseDTOs(mouse))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (mouseApi *MouseAPI) Create(c *gin.Context) {
	var mouseDTO dto.MouseDTO
	err := c.BindJSON(&mouseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	mouse := mapper.ToMouse(mouseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := mouseApi.IMouseService.Create(mouse)

	if error == nil {
		c.JSON(http.StatusOK, "Mouse stored successfully.")
	} else {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}

func (mouseApi *MouseAPI) Update(c *gin.Context) {
	var mouseDTO dto.MouseDTO
	err := c.BindJSON(&mouseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := mouseApi.IMouseService.Update(mouseDTO)

	if error == nil {
		c.JSON(http.StatusOK, "Mouse updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	} 
}

func (mouseApi *MouseAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := mouseApi.IMouseService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, "Mouse deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}