package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"product/dto"
	"product/model"
	"product/repository"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type productService struct {
	IConsoleRepository repository.IConsoleRepository
	IGraphicsCardRepository repository.IGraphicsCardRepository
	IHardDiskDriveRepository repository.IHardDiskDriveRepository
	IHeadphonesRepository repository.IHeadphonesRepository
	IKeyboardRepository repository.IKeyboardRepository
	IMonitorRepository repository.IMonitorRepository
	IMotherboardRepository repository.IMotherboardRepository
	IMouseRepository repository.IMouseRepository
	IPowerSupplyUnitRepository repository.IPowerSupplyUnitRepository
	IProcessorRepository repository.IProcessorRepository
	IRamRepository repository.IRamRepository
	ISolidStateDriveRepository repository.ISolidStateDriveRepository
	IVideoGameRepository repository.IVideoGameRepository
	IProductRepository repository.IProductRepository
}

type IProductService interface {
	GetProductById(id int) (model.Product, error)
	SearchByName(page int, pageSize int, name string) ([]model.Product, error)
	GetNumberOfRecordsSearch(name string) int64
	DeleteProduct(id int) error
	AddProductToCart(productPurchaseDTO dto.ProductPurchaseDTO, userData dto.UserData) (string, error)
	GetCurrentCart(userId int) []model.ProductPurchase
	CartContainsOnlyDigitalItems(userId int) bool
	GetPurchaseHistory(userId int) []model.ProductPurchase
	UpdatePurchase(productPurchaseDto dto.ProductPurchaseDTO) error
	RemoveProductFromCart(productPurchaseid int) error
	ConfirmPurchase(productPurchaseDto dto.ProductPurchaseDTO, userId int) error
	GetProductAlertByProductIdAndEmail(email string, productId int) (model.ProductAlert, error)
	AddProductAlert(userEmail string, productId int) string
	NotifyProductAvailability(productId int) (interface{}, error)
}

func NewProductService(
	consoleRepository repository.IConsoleRepository, 
	graphicsCardRepository repository.IGraphicsCardRepository,
	hardDiskDriveRepository repository.IHardDiskDriveRepository,
	headphonesRepository repository.IHeadphonesRepository,
	keyboardRepository repository.IKeyboardRepository,
	monitorRepository repository.IMonitorRepository,
	motherboardRepository repository.IMotherboardRepository,
	mouseRepository repository.IMouseRepository,
	powerSupplyUnitRepository repository.IPowerSupplyUnitRepository,
	processorRepository repository.IProcessorRepository,
	ramRepository repository.IRamRepository,
	solidStateDriveRepository repository.ISolidStateDriveRepository,
	videoGameRepository repository.IVideoGameRepository,
	productRepository repository.IProductRepository) IProductService {
	return &productService{
		IConsoleRepository: consoleRepository,
		IGraphicsCardRepository: graphicsCardRepository,
		IHardDiskDriveRepository: hardDiskDriveRepository,
		IHeadphonesRepository: headphonesRepository,
		IKeyboardRepository: keyboardRepository,
		IMonitorRepository: monitorRepository,
		IMotherboardRepository: motherboardRepository,
		IMouseRepository: mouseRepository,
		IPowerSupplyUnitRepository: powerSupplyUnitRepository,
		IProcessorRepository: processorRepository,
		IRamRepository: ramRepository,
		ISolidStateDriveRepository: solidStateDriveRepository,
		IVideoGameRepository: videoGameRepository,
		IProductRepository: productRepository,
	}
}


// General product related services 
func (productService *productService) GetProductById(id int) (model.Product, error) {
	return productService.IProductRepository.GetProductById(id);
}

func (productService *productService) SearchByName(page int, pageSize int, name string) ([]model.Product, error) {
	return productService.IProductRepository.SearchByName(page, pageSize, name)
}

func (productService *productService) GetNumberOfRecordsSearch(name string) int64 {
	return productService.IProductRepository.GetNumberOfRecordsSearch(name)
}

func (productService *productService) DeleteProduct(id int) error {
	product, err := productService.GetProductById(id)
	if err != nil {
		return err
	}
	*product.Archived = true
	return productService.IProductRepository.UpdateProduct(product)
}


// Product purchasing related services
func (productService *productService) GetCurrentCart(userId int) []model.ProductPurchase {
	return productService.IProductRepository.GetCurrentCart(userId)
}

func (productService *productService) CartContainsOnlyDigitalItems(userId int) bool {
	currentCart := productService.IProductRepository.GetCurrentCart(userId)
	digitalItemsInCart := productService.IProductRepository.GetAllDigitalItemsFromCart(userId)
	if len(currentCart) == len(digitalItemsInCart) {
		return true
	} else {
		return false
	}
}

func (productService *productService) GetPurchaseHistory(userId int) []model.ProductPurchase {
	return productService.IProductRepository.GetPurchaseHistory(userId)
}

func (productService *productService) AddProductToCart(productPurchaseDTO dto.ProductPurchaseDTO, userData dto.UserData) (string, error) {
	msg := ""
	productOutOfStockMsg := "Product is out of stock!"
	var productPurchase model.ProductPurchase
	product, err := productService.IProductRepository.GetProductById(productPurchaseDTO.Product.Id)

	if err != nil {
		return msg, err
	}

	if product.Type == model.VIDEO_GAME {
		videoGame, _ := productService.IVideoGameRepository.GetById(productPurchaseDTO.Product.Id)
		if videoGame.Product.Amount < productPurchaseDTO.Amount && !*videoGame.Digital {
			err = errors.New(productOutOfStockMsg)
		}
	} else {
		if product.Amount < productPurchaseDTO.Amount {
			err = errors.New(productOutOfStockMsg)
		}
	}

	if msg == "" {
		productInCart, err := productService.IProductRepository.GetProductPurchaseFromCart(product.Name, userData.Id)
		if err == nil {
			productInCart.Amount = productPurchaseDTO.Amount
			productInCart.TotalPrice = productInCart.Product.Price.Mul(decimal.NewFromInt(int64(productInCart.Amount)))
			productService.IProductRepository.UpdatePurchase(productInCart)
			msg = "Cart updated."
		} else {
			productPurchase.Product = product
			productPurchase.TotalPrice = productPurchase.Product.Price.Mul(decimal.NewFromInt(int64(productPurchaseDTO.Amount)))
			productPurchase.Amount = productPurchaseDTO.Amount
			productPurchase.UserId = userData.Id 
			productService.IProductRepository.AddPurchase(productPurchase);
			msg = "Product added to cart."
		}
	}
	return msg, nil
}

func (productService *productService) UpdatePurchase(productPurchaseDto dto.ProductPurchaseDTO) error {
	productPurchase, err := productService.IProductRepository.GetProductPurchaseById(productPurchaseDto.Id)
	if err != nil {
		return err
	}

	productPurchase.Amount = productPurchaseDto.Amount
	productPurchase.TotalPrice = productPurchase.Product.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))

	return productService.IProductRepository.UpdatePurchase(productPurchase)
}

func (productService *productService) RemoveProductFromCart(productPurchaseId int) error {
	product, err := productService.IProductRepository.GetProductPurchaseById(productPurchaseId)
	if err != nil {
		return err
	}
	return productService.IProductRepository.RemoveProductFromCart(product)
}

func (productService *productService) ConfirmPurchase(productPurchaseDto dto.ProductPurchaseDTO, userId int) error {
	productPurchases := productService.IProductRepository.GetCurrentCart(userId)
	for _, purchase := range productPurchases {
		purchase.PurchaseDate = time.Now().UTC()
		purchase.DeliveryAddress = productPurchaseDto.DeliveryAddress
		purchase.City = productPurchaseDto.City
		purchase.MobilePhoneNumber = productPurchaseDto.MobilePhoneNumber
		purchase.TypeOfPayment = productPurchaseDto.TypeOfPayment
		err := productService.IProductRepository.UpdatePurchase(purchase)
		if err != nil {
			return err
		}
		product, err := productService.IProductRepository.GetProductById(purchase.Product.Id)
		if err != nil {
			return err
		}

		if product.Type == model.VIDEO_GAME {
			videoGame, _ := productService.IVideoGameRepository.GetById(purchase.Product.Id)
			if !*videoGame.Digital {
				product.Amount -= purchase.Amount
				productService.IProductRepository.UpdateProduct(product)
			}
		} else {
			product.Amount -= purchase.Amount
			productService.IProductRepository.UpdateProduct(product)
		}
	}
	return nil
}


// Product alert related services
func (productService *productService) GetProductAlertByProductIdAndEmail(email string, productId int) (model.ProductAlert, error) {
	return productService.IProductRepository.GetProductAlertByProductIdAndEmail(email, productId)
}

func (productService *productService) AddProductAlert(userEmail string, productId int) string {
	_, err := productService.IProductRepository.GetProductAlertByProductIdAndEmail(userEmail, productId)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "You are already registered to receive emails for this product."
	}

	product, err := productService.IProductRepository.GetProductById(productId)
	if err != nil {
		return err.Error()
	}

	var productAlert model.ProductAlert
	productAlert.Product = product
	productAlert.ProductId = product.Id
	productAlert.UserEmail = userEmail
	productService.IProductRepository.AddProductAlert(productAlert)

	return ""
}

func (productService *productService) NotifyProductAvailability(productId int) (interface{}, error) {
	recipients := productService.IProductRepository.GetUserEmailsByProductId(productId)
	product, _ := productService.IProductRepository.GetProductById(productId)
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

	for _, recipient := range recipients {
		productService.IProductRepository.RemoveProductAlertByEmailAndProductId(recipient, productId)
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