package api

import (
	"comment-and-rating/dto"
	"comment-and-rating/mapper"
	"comment-and-rating/middleware"
	"comment-and-rating/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)


type ProductCommentAPI struct {
	IProductCommentService service.IProductCommentService
}

func NewProductCommentAPI(productCommentService service.IProductCommentService) ProductCommentAPI {
	return ProductCommentAPI{IProductCommentService: productCommentService}
}

func (productCommentApi *ProductCommentAPI) GetAll(c *gin.Context) {
	productComments := productCommentApi.IProductCommentService.GetAll()
	c.JSON(http.StatusOK, mapper.ToProductCommentDTOs(productComments))
}

func (productCommentApi *ProductCommentAPI) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	productComment, err := productCommentApi.IProductCommentService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToProductCommentDTO(productComment))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (productCommentApi *ProductCommentAPI) GetByProductId(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	productCommentDTOs := productCommentApi.IProductCommentService.GetByProductId(productId)
	c.JSON(http.StatusOK, productCommentDTOs)
}

func (productCommentApi *ProductCommentAPI) GetByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	productComments := productCommentApi.IProductCommentService.GetByUserId(userId)
	c.JSON(http.StatusOK, mapper.ToProductCommentDTOs(productComments))
}

func (productCommentApi *ProductCommentAPI) GetByProductAndUser(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	productComment, err := productCommentApi.IProductCommentService.GetByProductAndUser(productId, userId)

	if err == nil {
		c.JSON(http.StatusOK, mapper.ToProductCommentDTO(productComment))
	} else  {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (productCommentApi *ProductCommentAPI) AddComment(c *gin.Context) {
	var productCommentDTO dto.ProductCommentDTO
	err := c.BindJSON(&productCommentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	productComment := mapper.ToProductComment(productCommentDTO)
	userData := middleware.GetUserData(c)
	msg := productCommentApi.IProductCommentService.AddComment(productComment, userData.Id)

	if msg == "" {
		c.JSON(http.StatusOK, "Comment and rating added successfully.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (productCommentApi *ProductCommentAPI) EditComment(c *gin.Context) {
	var productCommentDTO dto.ProductCommentDTO
	err := c.BindJSON(&productCommentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := productCommentApi.IProductCommentService.EditComment(productCommentDTO)

	if msg == "" {
		c.JSON(http.StatusOK, "Comment and rating updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}

func (productCommentApi *ProductCommentAPI) DeleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := productCommentApi.IProductCommentService.DeleteComment(id)

	if error == nil {
		c.JSON(http.StatusOK, "Comment and rating deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}