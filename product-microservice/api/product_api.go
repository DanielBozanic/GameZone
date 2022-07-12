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

func (productApi *ProductAPI) GetProductById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
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

func (productApi *ProductAPI) AddProductToCart(c *gin.Context) {
	var productPurchaseDTO dto.ProductPurchaseDTO
	err := c.BindJSON(&productPurchaseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userData := middleware.GetUserData(c)
	msg, err := productApi.IProductService.AddProductToCart(productPurchaseDTO, userData);

	if msg == "" {
		c.JSON(http.StatusOK, "Product added to cart.")
	} else if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (productApi *ProductAPI) GetCurrentCart(c *gin.Context) {
	userData := middleware.GetUserData(c)
	productPurchases := productApi.IProductService.GetCurrentCart(userData.Id)
	c.JSON(http.StatusOK, mapper.ToProductPurchaseDTOs(productPurchases))
}

func (productApi *ProductAPI) GetPurchaseHistory(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	productPurchases := productApi.IProductService.GetPurchaseHistory(userId)
	c.JSON(http.StatusOK, mapper.ToProductPurchaseDTOs(productPurchases))
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

func (productApi *ProductAPI) UpdatePurchase(c *gin.Context) {
	var productPurchaseDTO dto.ProductPurchaseDTO
	err := c.BindJSON(&productPurchaseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := productApi.IProductService.UpdatePurchase(productPurchaseDTO)

	if error == nil {
		c.JSON(http.StatusOK, "Cart updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	} 
}

func (productApi *ProductAPI) RemoveProductFromCart(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := productApi.IProductService.RemoveProductFromCart(id)

	if error == nil {
		c.JSON(http.StatusOK, "Product removed from cart.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}

func (productApi *ProductAPI) ConfirmPurchase(c *gin.Context) {
	var productPurchaseDTO dto.ProductPurchaseDTO
	err := c.BindJSON(&productPurchaseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userData := middleware.GetUserData(c)
	error := productApi.IProductService.ConfirmPurchase(productPurchaseDTO, userData.Id)

	if error == nil {
		c.Status(200)
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}