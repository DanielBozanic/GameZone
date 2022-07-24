package api

import (
	"product/dto"
	"product/mapper"
	"product/middleware"
	"product/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)


type ProductPurchaseAPI struct {
	IProductPurchaseService service.IProductPurchaseService
}

func NewProductPurchaseAPI(productPurchaseService service.IProductPurchaseService) ProductPurchaseAPI {
	return ProductPurchaseAPI{IProductPurchaseService: productPurchaseService}
}

func (productPurchaseApi *ProductPurchaseAPI) GetPurchaseHistory(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	productPurchases := productPurchaseApi.IProductPurchaseService.GetPurchaseHistory(userId)
	c.JSON(http.StatusOK, mapper.ToProductPurchaseDTOs(productPurchases))
}

func (productPurchaseApi *ProductPurchaseAPI) CheckIfProductIsPaidFor(c *gin.Context) {
	productId, err := strconv.Atoi(c.Query("productId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userData := middleware.GetUserData(c)
	isPaidFor := productPurchaseApi.IProductPurchaseService.CheckIfProductIsPaidFor(productId, userData.Id)
	c.JSON(http.StatusOK, isPaidFor)
}

func (productPurchaseApi *ProductPurchaseAPI) ConfirmPurchase(c *gin.Context) {
	var productPurchaseDTO dto.ProductPurchaseDTO
	err := c.BindJSON(&productPurchaseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userData := middleware.GetUserData(c)
	error := productPurchaseApi.IProductPurchaseService.ConfirmPurchase(productPurchaseDTO, userData.Id)

	if error == nil {
		c.JSON(http.StatusOK, "Purchase successful")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}

func (productPurchaseApi *ProductPurchaseAPI) GetProductAlertByProductIdAndEmail(c *gin.Context) {
	productId, err := strconv.Atoi(c.Query("productId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	userData := middleware.GetUserData(c)
	productAlert, err := productPurchaseApi.IProductPurchaseService.GetProductAlertByProductIdAndEmail(userData.Email, productId)

	if err == nil {
		c.JSON(http.StatusOK, productAlert)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func (productPurchaseApi *ProductPurchaseAPI) AddProductAlert(c *gin.Context) {
	productId, err := strconv.Atoi(c.Query("productId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userData := middleware.GetUserData(c)
	msg := productPurchaseApi.IProductPurchaseService.AddProductAlert(userData.Email, productId);

	if msg == "" {
		c.JSON(http.StatusOK, "You will be notified via email when product is in stock.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	} 
}

func (productPurchaseApi *ProductPurchaseAPI) NotifyProductAvailability(c *gin.Context) {
	productId, err := strconv.Atoi(c.Query("productId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := productPurchaseApi.IProductPurchaseService.NotifyProductAvailability(productId);

	if err == nil {
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}