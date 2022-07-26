package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"product/dto"
	"product/mapper"
	"product/model"
	"product/repository"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type productPurchaseService struct {
	IVideoGameRepository repository.IVideoGameRepository
	IProductRepository repository.IProductRepository
	IProductPurchaseRepository repository.IProductPurchaseRepository
}

type IProductPurchaseService interface {
	GetPurchaseHistory(userId int) []model.ProductPurchase
	CheckIfProductIsPaidFor(productId int, userId int) bool
	ConfirmPurchase(productPurchaseDto dto.ProductPurchaseDTO, userId int) error
	GetProductAlertByProductIdAndUserId(userId int, productId int) (model.ProductAlert, error)
	AddProductAlert(userId int, productId int) string
	NotifyProductAvailability(productId int) (interface{}, error)
}

func NewProductPurchaseService(
	videoGameRepository repository.IVideoGameRepository,
	productRepository repository.IProductRepository, 
	productPurchaseRepository repository.IProductPurchaseRepository) IProductPurchaseService {
	return &productPurchaseService{
		IVideoGameRepository: videoGameRepository,
		IProductRepository: productRepository,
		IProductPurchaseRepository: productPurchaseRepository,
	}
}

// Product purchasing related services
func (productPurchaseService *productPurchaseService) GetPurchaseHistory(userId int) []model.ProductPurchase {
	return productPurchaseService.IProductPurchaseRepository.GetPurchaseHistory(userId)
}

func (productPurchaseService *productPurchaseService) CheckIfProductIsPaidFor(productId int, userId int) bool {
	_, err := productPurchaseService.IProductPurchaseRepository.GetPaidProductPurchase(productId, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func (productPurchaseService *productPurchaseService) ConfirmPurchase(productPurchaseDto dto.ProductPurchaseDTO, userId int) error {
	productPurchase := mapper.ToProductPurchase(productPurchaseDto)
	productPurchase.PurchaseDate = time.Now()
	productPurchase.UserId = userId
	return productPurchaseService.IProductPurchaseRepository.AddPurchase(productPurchase)
}


// Product alert related services
func (productPurchaseService *productPurchaseService) GetProductAlertByProductIdAndUserId(userId int, productId int) (model.ProductAlert, error) {
	return productPurchaseService.IProductPurchaseRepository.GetProductAlertByProductIdAndUserId(userId, productId)
}

func (productPurchaseService *productPurchaseService) AddProductAlert(userId int, productId int) string {
	_, err := productPurchaseService.IProductPurchaseRepository.GetProductAlertByProductIdAndUserId(userId, productId)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "You are already registered to receive emails for this product."
	}

	product, err := productPurchaseService.IProductRepository.GetProductById(productId)
	if err != nil {
		return err.Error()
	}

	var productAlert model.ProductAlert
	productAlert.Product = product
	productAlert.ProductId = product.Id
	productAlert.UserId = userId
	productPurchaseService.IProductPurchaseRepository.AddProductAlert(productAlert)

	return ""
}

func (productPurchaseService *productPurchaseService) NotifyProductAvailability(productId int) (interface{}, error) {
	userIds := productPurchaseService.IProductPurchaseRepository.GetUserIdsByProductId(productId)
	product, _ := productPurchaseService.IProductRepository.GetProductById(productId)
	recipients := []string{}
	for _, userId := range userIds {
		req, err := http.NewRequest("GET", "http://localhost:5000/api/users/getById?userId=" +  strconv.Itoa(userId), nil)
		client := &http.Client{}
		resp, err := client.Do(req)

		var target map[string]interface{}
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&target)
		email := target["user"].(map[string]interface{})["email"].(string)
		recipients = append(recipients, email)
	}
	data := map[string]interface{}{
		"subject": "Product in stock" ,
		"recipients": recipients,
		"content": map[string]interface{}{
			"template": "product_in_stock",
			"params": map[string]interface{}{
				"productName": product.Name,
				"productRoute": getProductRoute(product),
			},
		},
	}
	jsonData, _ := json.Marshal(data)
	
	req, err := http.NewRequest("POST", "http://localhost:5001/api/email/sendEmail", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	var target interface{}
	if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(target)

	for _, userId := range userIds {
		productPurchaseService.IProductPurchaseRepository.RemoveProductAlertByUserIdAndProductId(userId, productId)
	}
	return target, nil
}

func getProductRoute(product model.Product) string {
	switch {
		case model.CONSOLE == product.Type:
			return "/consoles/" + strconv.Itoa(product.Id);
		case model.GRAPHICS_CARD == product.Type:
			return "/graphicsCards/" + strconv.Itoa(product.Id);
		case model.HARD_DISK_DRIVE == product.Type:
			return "/hdds/" + strconv.Itoa(product.Id);
		case model.HEADPHONES == product.Type:
			return "/headphones/" + strconv.Itoa(product.Id);
		case model.KEYBOARD == product.Type:
			return "/keyboards/" + strconv.Itoa(product.Id);
		case model.MONITOR == product.Type:
			return "/monitors/" + strconv.Itoa(product.Id);
		case model.MOTHERBOARD == product.Type:
			return "/motherboards/" + strconv.Itoa(product.Id);
		case model.MOUSE == product.Type:
			return "/mice/" + strconv.Itoa(product.Id);
		case model.POWER_SUPPLY_UNIT == product.Type:
			return "/psus/" + strconv.Itoa(product.Id);
		case model.PROCESSOR == product.Type:
			return "/processors/" + strconv.Itoa(product.Id);
		case model.RAM == product.Type:
			return "/rams/" + strconv.Itoa(product.Id);
		case model.SOLID_STATE_DRIVE == product.Type:
			return "/ssds/" + strconv.Itoa(product.Id);
		case model.VIDEO_GAME == product.Type:
			return "/videoGames/" + strconv.Itoa(product.Id);
		default:
			return ""
	}
};