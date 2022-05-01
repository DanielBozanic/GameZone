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


type HeadphonesAPI struct {
	IHeadphonesService service.IHeadphonesService
}

func NewHeadphonesAPI(headphonesService service.IHeadphonesService) HeadphonesAPI {
	return HeadphonesAPI{IHeadphonesService: headphonesService}
}

func (headphonesApi *HeadphonesAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }
	
	headphoness := headphonesApi.IHeadphonesService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, gin.H{"headphoness": mapper.ToHeadphonesDTOs(headphoness)})
}

func (headphonesApi *HeadphonesAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	headphones, err := headphonesApi.IHeadphonesService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"headphones": mapper.ToHeadphonesDTO(headphones)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (headphonesApi *HeadphonesAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }

	headphones, err := headphonesApi.IHeadphonesService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"headphones": mapper.ToHeadphonesDTOs(headphones)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (headphonesApi *HeadphonesAPI) Create(c *gin.Context) {
	var headphonesDTO dto.HeadphonesDTO
	err := c.BindJSON(&headphonesDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	headphones := mapper.ToHeadphones(headphonesDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := headphonesApi.IHeadphonesService.Create(headphones)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Headphones stored successfully."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}

func (headphonesApi *HeadphonesAPI) Update(c *gin.Context) {
	var headphonesDTO dto.HeadphonesDTO
	err := c.BindJSON(&headphonesDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := headphonesApi.IHeadphonesService.Update(headphonesDTO)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Headphones updated successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	} 
}

func (headphonesApi *HeadphonesAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	error := headphonesApi.IHeadphonesService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Headphones deleted successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}