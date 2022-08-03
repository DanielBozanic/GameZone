package api

import (
	"news/dto"
	"news/mapper"
	"news/middleware"
	"news/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)


type NewsCommentAPI struct {
	INewsCommentService service.INewsCommentService
}

func NewNewsCommentAPI(newsCommentService service.INewsCommentService) NewsCommentAPI {
	return NewsCommentAPI{INewsCommentService: newsCommentService}
}

func (newsCommentApi *NewsCommentAPI) GetAll(c *gin.Context) {
	newsComments := newsCommentApi.INewsCommentService.GetAll()
	c.JSON(http.StatusOK, mapper.ToNewsCommentDTOs(newsComments))
}

func (newsCommentApi *NewsCommentAPI) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	newsComment, err := newsCommentApi.INewsCommentService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToNewsCommentDTO(newsComment))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (newsCommentApi *NewsCommentAPI) GetByNewsArticle(c *gin.Context) {
	newsArticleId, err := strconv.Atoi(c.Param("newsArticleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newsComments := newsCommentApi.INewsCommentService.GetByNewsArticle(newsArticleId)
	c.JSON(http.StatusOK, mapper.ToNewsCommentDTOs(newsComments))
}

func (newsCommentApi *NewsCommentAPI) GetByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newsComments := newsCommentApi.INewsCommentService.GetByUserId(userId)
	c.JSON(http.StatusOK, mapper.ToNewsCommentDTOs(newsComments))
}

func (newsCommentApi *NewsCommentAPI) AddNewsComment(c *gin.Context) {
	var newsCommentDTO dto.NewsCommentDTO
	err := c.BindJSON(&newsCommentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newsComment := mapper.ToNewsComment(newsCommentDTO)
	userData := middleware.GetUserData(c)
	msg := newsCommentApi.INewsCommentService.AddNewsComment(newsComment, userData)

	if msg == "" {
		c.JSON(http.StatusOK, "Article comment added.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (newsCommentApi *NewsCommentAPI) EditNewsCommment(c *gin.Context) {
	var newsCommentDTO dto.NewsCommentDTO
	err := c.BindJSON(&newsCommentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := newsCommentApi.INewsCommentService.EditNewsCommment(newsCommentDTO)

	if msg == "" {
		c.JSON(http.StatusOK, "Article comment updated.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}

func (newsCommentApi *NewsCommentAPI) DeleteNewsComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := newsCommentApi.INewsCommentService.DeleteNewsComment(id)

	if error == nil {
		c.JSON(http.StatusOK, "Article comment deleted.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}