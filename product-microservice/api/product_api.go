package api

import (
	"product/dto"
	"product/middleware"
	"product/service"

	"net/http"

	"github.com/gin-gonic/gin"
)


type ProductAPI struct {
	IProductService service.IProductService
}

func NewProductAPI(productService service.IProductService) ProductAPI {
	return ProductAPI{IProductService: productService}
}

func (productApi *ProductAPI) AddProductToCart(c *gin.Context) {
	var productPurchaseDTO dto.ProductPurchaseDTO
	err := c.BindJSON(&productPurchaseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userData := middleware.GetUserData(c)
	msg := productApi.IProductService.AddProductToCart(productPurchaseDTO, userData);

	if msg == "" {
		c.JSON(http.StatusOK, gin.H{"msg": "Product added to cart."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
	}
}