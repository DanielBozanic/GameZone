package service

import (
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
	AddProductToCart(productPurchaseDto dto.ProductPurchaseDTO, userData dto.UserData) (string, error)
	GetCurrentCart(userId int) []model.ProductPurchase
	GetPurchaseHistory(userId int) []model.ProductPurchase
	UpdatePurchase(productPurchaseDto dto.ProductPurchaseDTO) error
	RemoveProductFromCart(productPurchaseId uuid.UUID) error
	ConfirmPurchase(userId int) error
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

func (productService *productService) GetCurrentCart(userId int) []model.ProductPurchase {
	return productService.IProductRepository.GetCurrentCart(userId)
}

func (productService *productService) GetPurchaseHistory(userId int) []model.ProductPurchase {
	return productService.IProductRepository.GetPurchaseHistory(userId)
}

func (productService *productService) AddProductToCart(productPurchaseDto dto.ProductPurchaseDTO, userData dto.UserData) (string, error) {
	msg := ""
	productOutOfStockMsg := "Product is out of stock!"
	var productPurchase model.ProductPurchase
	product, err := productService.IProductRepository.GetProductById(productPurchaseDto.ProductId)

	if err != nil {
		return msg, err
	}

	if product.Type == model.VIDEO_GAME {
		videoGame, _ := productService.IVideoGameRepository.GetById(productPurchaseDto.ProductId)
		if videoGame.Product.Amount < productPurchaseDto.Amount && !videoGame.Digital {
			msg = productOutOfStockMsg
		}
	} else {
		if product.Amount < productPurchaseDto.Amount {
			msg = productOutOfStockMsg
		}
	}

	if msg == "" {
		productService.IProductRepository.AddPurchase(productPurchase);
	}
	return msg, nil
}

func (productService *productService) UpdatePurchase(productPurchaseDto dto.ProductPurchaseDTO) error {
	productPurchase, err := productService.IProductRepository.GetProductPurchaseById(productPurchaseDto.Id)
	if err != nil {
		return err
	}

	if productPurchaseDto.Amount == 0 {
		return productService.IProductRepository.RemoveProductFromCart(productPurchase)
	}

	productPurchase.Amount = productPurchaseDto.Amount
	productPrice := productPurchase.TotalPrice.Div(decimal.NewFromInt(int64(productPurchase.Amount)))
	productPurchase.TotalPrice = productPrice.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))

	return productService.IProductRepository.UpdatePurchase(productPurchase)
}

func (productService *productService) RemoveProductFromCart(productPurchaseId uuid.UUID) error {
	product, err := productService.IProductRepository.GetProductPurchaseById(productPurchaseId)
	if err != nil {
		return err
	}
	return productService.IProductRepository.RemoveProductFromCart(product)
}

func (productService *productService) ConfirmPurchase(userId int) error {
	productPurchases := productService.IProductRepository.GetCurrentCart(userId)
	for _, purchase := range productPurchases {
		purchase.PurchaseDate = time.Now().UTC()
		err := productService.IProductRepository.UpdatePurchase(purchase)
		if err != nil {
			return err
		}
	}
	return nil
}