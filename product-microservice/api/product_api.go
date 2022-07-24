package api

import (
	"product/mapper"
	"product/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)


type ProductAPI struct {
	IProductService service.IProductService
}

func NewProductAPI(productService service.IProductService) ProductAPI {
	return ProductAPI{IProductService: productService}
}

func (productApi *ProductAPI) GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	product, err := productApi.IProductService.GetProductById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToProductDTO(product))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (productApi *ProductAPI) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	msg := productApi.IProductService.DeleteProduct(id)

	if msg == "" {
		c.JSON(http.StatusOK, "Product deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (productApi *ProductAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	products, err := productApi.IProductService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToProductDTOs(products))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (productApi *ProductAPI) GetNumberOfRecordsSearch(c *gin.Context) {
	numberOfRecords := productApi.IProductService.GetNumberOfRecordsSearch(c.Query("name"))
	c.JSON(http.StatusOK, numberOfRecords)
}