package api

import (
	"news/dto"
	"news/mapper"
	"news/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type NewsArticleAPI struct {
	INewsArticleService service.INewsArticleService
}

func NewNewsArticleAPI(newsArticleService service.INewsArticleService) NewsArticleAPI {
	return NewsArticleAPI{INewsArticleService: newsArticleService}
}

func (newsArticleApi *NewsArticleAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	newsArticles := newsArticleApi.INewsArticleService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToNewsArticleDTOs(newsArticles))
}

func (newsArticleApi *NewsArticleAPI) GetNumberOfRecords(c *gin.Context) {
	numberOfRecords := newsArticleApi.INewsArticleService.GetNumberOfRecords()
	c.JSON(http.StatusOK, numberOfRecords)
}

func (newsArticleApi *NewsArticleAPI) GetPublishedArticles(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	newsArticles := newsArticleApi.INewsArticleService.GetPublishedArticles(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToNewsArticleDTOs(newsArticles))
}

func (newsArticleApi *NewsArticleAPI) GetNumberOfRecordsPublishedArticles(c *gin.Context) {
	numberOfRecords := newsArticleApi.INewsArticleService.GetNumberOfRecordsPublishedArticles()
	c.JSON(http.StatusOK, numberOfRecords)
}

func (newsArticleApi *NewsArticleAPI) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	newsArticle, err := newsArticleApi.INewsArticleService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToNewsArticleDTO(newsArticle))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (newsArticleApi *NewsArticleAPI) AddNewsArticle(c *gin.Context) {
	var newsArticleDTO dto.NewsArticleDTO
	err := c.BindJSON(&newsArticleDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newsArticle := mapper.ToNewsArticle(newsArticleDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := newsArticleApi.INewsArticleService.AddNewsArticle(newsArticle)

	if msg == "" {
		c.JSON(http.StatusOK, "News article added successfully.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (newsArticleApi *NewsArticleAPI) EditNewsArticle(c *gin.Context) {
	var newsArticleDTO dto.NewsArticleDTO
	err := c.BindJSON(&newsArticleDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := newsArticleApi.INewsArticleService.EditNewsArticle(newsArticleDTO)

	if msg == "" {
		c.JSON(http.StatusOK, "News article updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}

func (newsArticleApi *NewsArticleAPI) DeleteNewsArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := newsArticleApi.INewsArticleService.DeleteNewsArticle(id)

	if error == nil {
		c.JSON(http.StatusOK, "News article deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}

func (newsArticleApi *NewsArticleAPI) PublishNewsArticle(c *gin.Context) {
	var newsArticleDTO dto.NewsArticleDTO
	err := c.BindJSON(&newsArticleDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := newsArticleApi.INewsArticleService.PublishNewsArticle(newsArticleDTO)

	if msg == "" {
		c.JSON(http.StatusOK, "News article published.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}