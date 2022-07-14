package service

import (
	"errors"
	"product/dto"
	"product/model"
	"product/repository"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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
	GetProductById(id uuid.UUID) (model.Product, error)
	AddProductToCart(productPurchaseDTO dto.ProductPurchaseDTO, userData dto.UserData) (string, error)
	GetCurrentCart(userId int) []model.ProductPurchase
	CartContainsOnlyDigitalItems(userId int) bool
	GetPurchaseHistory(userId int) []model.ProductPurchase
	SearchByName(page int, pageSize int, name string) ([]model.Product, error)
	GetNumberOfRecordsSearch(name string) int64
	UpdatePurchase(productPurchaseDto dto.ProductPurchaseDTO) error
	RemoveProductFromCart(productPurchaseId uuid.UUID) error
	ConfirmPurchase(productPurchaseDto dto.ProductPurchaseDTO, userId int) error
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

func (productService *productService) GetProductById(id uuid.UUID) (model.Product, error) {
	return productService.IProductRepository.GetProductById(id);
}

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

func (productService *productService) SearchByName(page int, pageSize int, name string) ([]model.Product, error) {
	return productService.IProductRepository.SearchByName(page, pageSize, name)
}

func (productService *productService) GetNumberOfRecordsSearch(name string) int64 {
	return productService.IProductRepository.GetNumberOfRecordsSearch(name)
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
		if videoGame.Product.Amount < productPurchaseDTO.Amount && !videoGame.Digital {
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
			productPurchase.Id = uuid.New()
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

func (productService *productService) RemoveProductFromCart(productPurchaseId uuid.UUID) error {
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
			if !videoGame.Digital {
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