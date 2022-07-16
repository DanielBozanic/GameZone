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

func (mouseApi *MouseAPI) GetNumberOfRecords(c *gin.Context) {
	numberOfRecords := mouseApi.IMouseService.GetNumberOfRecords()
	c.JSON(http.StatusOK, numberOfRecords)
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

func (mouseApi *MouseAPI) GetNumberOfRecordsSearch(c *gin.Context) {
	numberOfRecords := mouseApi.IMouseService.GetNumberOfRecordsSearch(c.Query("name"))
	c.JSON(http.StatusOK, numberOfRecords)
}

func (mouseApi *MouseAPI) Filter(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	var filter filter.MouseFilter
	err = c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	mouse, err := mouseApi.IMouseService.Filter(page, pageSize, filter)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToMouseDTOs(mouse))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (mouseApi *MouseAPI) GetNumberOfRecordsFilter(c *gin.Context) {
	var filter filter.MouseFilter
	err := c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	numberOfRecords := mouseApi.IMouseService.GetNumberOfRecordsFilter(filter)
	c.JSON(http.StatusOK, numberOfRecords)
}

func (mouseApi *MouseAPI) GetManufacturers(c *gin.Context) {
	manufacturers := mouseApi.IMouseService.GetManufacturers()
	c.JSON(http.StatusOK, manufacturers)
}

func (mouseApi *MouseAPI) GetDPIs(c *gin.Context) {
	DPIs := mouseApi.IMouseService.GetDPIs()
	c.JSON(http.StatusOK, DPIs)
}

func (mouseApi *MouseAPI) GetConnections(c *gin.Context) {
	connections := mouseApi.IMouseService.GetConnections()
	c.JSON(http.StatusOK, connections)
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

	msg := mouseApi.IMouseService.Create(mouse)

	if msg == "" {
		c.JSON(http.StatusOK, "Mouse added successfully.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (mouseApi *MouseAPI) Update(c *gin.Context) {
	var mouseDTO dto.MouseDTO
	err := c.BindJSON(&mouseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := mouseApi.IMouseService.Update(mouseDTO)

	if msg == "" {
		c.JSON(http.StatusOK, "Mouse updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
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