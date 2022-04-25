package api

import (
	"product/dto"
	"product/mapper"
	"product/middleware"
	"product/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	msg, err := productApi.IProductService.AddProductToCart(productPurchaseDTO, userData);

	if msg == "" {
		c.JSON(http.StatusOK, gin.H{"msg": "Product added to cart."})
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
	}
}

func (productApi *ProductAPI) GetCurrentCart(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }
	productPurchases := productApi.IProductService.GetCurrentCart(userId)
	c.JSON(http.StatusOK, gin.H{"cart": mapper.ToProductPurchaseDTOs(productPurchases)})
}

func (productApi *ProductAPI) GetPurchaseHistory(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }
	productPurchases := productApi.IProductService.GetPurchaseHistory(userId)
	c.JSON(http.StatusOK, gin.H{"purchase_history": mapper.ToProductPurchaseDTOs(productPurchases)})
}

func (productApi *ProductAPI) UpdatePurchase(c *gin.Context) {
	var productPurchaseDTO dto.ProductPurchaseDTO
	err := c.BindJSON(&productPurchaseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := productApi.IProductService.UpdatePurchase(productPurchaseDTO)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Cart updated successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	} 
}

func (productApi *ProductAPI) RemoveProductFromCart(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := productApi.IProductService.RemoveProductFromCart(id)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Product removed from cart."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}

func (productApi *ProductAPI) ConfirmPurchase(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }

	error := productApi.IProductService.ConfirmPurchase(userId)

	if error == nil {
		c.Status(200)
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}