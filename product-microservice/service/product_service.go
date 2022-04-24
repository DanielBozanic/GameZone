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
	AddProductToCart(productPurchaseDto dto.ProductPurchaseDTO, userData dto.UserData) string
	GetCurrentCart(userId int) []dto.ProductPurchaseDTO
	UpdatePurchase(productPurchaseDto dto.ProductPurchaseDTO) error
	RemoveProductFromCart(productPurchaseId uuid.UUID) error
	ConfirmPurchase(userId int)
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

func (productService *productService) GetCurrentCart(userId int) []dto.ProductPurchaseDTO {
	var productPurchaseDto dto.ProductPurchaseDTO
	var productPurchaseDtos []dto.ProductPurchaseDTO
	productPurchases := productService.IProductRepository.GetCurrentCart(userId)

	for _, p := range productPurchases {
		productPurchaseDto.Amount = p.Amount
		productPurchaseDto.Id = p.Id
		productPurchaseDto.ProductId = p.ProductId
		productPurchaseDto.ProductName = p.ProductName
		productPurchaseDto.TotalPrice = p.TotalPrice
		productPurchaseDtos = append(productPurchaseDtos, productPurchaseDto)
	}

	return productPurchaseDtos
}

func (productService *productService) GetPurchaseHistory(userId int) []dto.ProductPurchaseDTO {
	var productPurchaseDto dto.ProductPurchaseDTO
	var productPurchaseDtos []dto.ProductPurchaseDTO
	productPurchases := productService.IProductRepository.GetPurchaseHistory(userId)

	for _, p := range productPurchases {
		productPurchaseDto.Amount = p.Amount
		productPurchaseDto.Id = p.Id
		productPurchaseDto.ProductId = p.ProductId
		productPurchaseDto.ProductName = p.ProductName
		productPurchaseDto.TotalPrice = p.TotalPrice
		productPurchaseDto.PurchaseDate = p.PurchaseDate.Format("2006-01-02")
		productPurchaseDtos = append(productPurchaseDtos, productPurchaseDto)
	}

	return productPurchaseDtos
}

func (productService *productService) AddProductToCart(productPurchaseDto dto.ProductPurchaseDTO, userData dto.UserData) string {
	var product model.ProductPurchase
	msg := ""

	console, _ := productService.IConsoleRepository.GetById(productPurchaseDto.ProductId)
	graphicsCard, _ := productService.IGraphicsCardRepository.GetById(productPurchaseDto.ProductId)
	hardDisk, _ := productService.IHardDiskDriveRepository.GetById(productPurchaseDto.ProductId)
	headphones, _ := productService.IHeadphonesRepository.GetById(productPurchaseDto.ProductId)
	keyboard, _ := productService.IKeyboardRepository.GetById(productPurchaseDto.ProductId)
	monitor, _ := productService.IMonitorRepository.GetById(productPurchaseDto.ProductId)
	motherboard, _ := productService.IMotherboardRepository.GetById(productPurchaseDto.ProductId)
	mouse, _ := productService.IMouseRepository.GetById(productPurchaseDto.ProductId)
	psu, _ := productService.IPowerSupplyUnitRepository.GetById(productPurchaseDto.ProductId)
	processor, _ := productService.IProcessorRepository.GetById(productPurchaseDto.ProductId)
	ram, _ := productService.IRamRepository.GetById(productPurchaseDto.ProductId)
	ssd, _ := productService.ISolidStateDriveRepository.GetById(productPurchaseDto.ProductId)
	videoGame, _ := productService.IVideoGameRepository.GetById(productPurchaseDto.ProductId)


	switch {
		case console.Id == productPurchaseDto.ProductId:
			product, msg = setProduct(console.Amount, console.Id, console.Price, console.Name, userData, productPurchaseDto)
			break
		case graphicsCard.Id == productPurchaseDto.ProductId:
			product, msg = setProduct(graphicsCard.Amount, graphicsCard.Id, graphicsCard.Price, graphicsCard.Name, userData, productPurchaseDto)
			break
		case hardDisk.Id == productPurchaseDto.ProductId:
			product, msg = setProduct(hardDisk.Amount, hardDisk.Id, hardDisk.Price, hardDisk.Name, userData, productPurchaseDto)
			break
		case headphones.Id == productPurchaseDto.ProductId:
			product, msg = setProduct(headphones.Amount, headphones.Id, headphones.Price, headphones.Name, userData, productPurchaseDto)
			break
		case keyboard.Id == productPurchaseDto.ProductId:
			product, msg = setProduct(keyboard.Amount, keyboard.Id, keyboard.Price, keyboard.Name, userData, productPurchaseDto)
			break
		case monitor.Id == productPurchaseDto.ProductId:
			product, msg = setProduct(monitor.Amount, monitor.Id, monitor.Price, monitor.Name, userData, productPurchaseDto)
			break
		case motherboard.Id == productPurchaseDto.ProductId:
			product, msg = setProduct(motherboard.Amount, motherboard.Id, motherboard.Price, motherboard.Name, userData, productPurchaseDto)
			break
		case mouse.Id == productPurchaseDto.ProductId:
			product, msg = setProduct(mouse.Amount, mouse.Id, mouse.Price, mouse.Name, userData, productPurchaseDto)
			break
		case psu.Id == productPurchaseDto.ProductId:
			product, msg = setProduct(psu.Amount, psu.Id, psu.Price, psu.Name, userData, productPurchaseDto)
			break
		case processor.Id == productPurchaseDto.ProductId:
			product, msg = setProduct(processor.Amount, processor.Id, processor.Price, processor.Name, userData, productPurchaseDto)
			break
		case ram.Id == productPurchaseDto.ProductId:
			product, msg = setProduct(ram.Amount, ram.Id, ram.Price, ram.Name, userData, productPurchaseDto)
			break
		case ssd.Id == productPurchaseDto.ProductId:
			product, msg = setProduct(ssd.Amount, ssd.Id, ssd.Price, ssd.Name, userData, productPurchaseDto)
			break
		case videoGame.Id == productPurchaseDto.ProductId:
			product, msg = setVideoGame(videoGame, userData, productPurchaseDto)
			break
		default:
			msg = "Product not found!"
	}

	if msg == "" {
		productService.IProductRepository.AddPurchase(product);
	}
	return msg
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

func (productService *productService) ConfirmPurchase(userId int) {
	productPurchases := productService.IProductRepository.GetCurrentCart(userId)
	for _, purchase := range productPurchases {
		purchase.PurchaseDate = time.Now().UTC()
		productService.IProductRepository.UpdatePurchase(purchase)
	}
}

func setProduct(productAmount uint, productId uuid.UUID, productPrice decimal.Decimal, name string,
	userData dto.UserData, productPurchaseDto dto.ProductPurchaseDTO) (model.ProductPurchase, string) {
	var product model.ProductPurchase
	msg := ""
	productOutOfStockMsg := "Product is out of stock!"
	if productAmount < productPurchaseDto.Amount {
		msg = productOutOfStockMsg
	}
	product.ProductName = name
	product.ProductId = productId
	product.Amount = productPurchaseDto.Amount
	product.UserId = userData.Id
	product.TotalPrice = productPrice.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
	return product, msg
}

func setVideoGame(videoGame model.VideoGame, userData dto.UserData, productPurchaseDto dto.ProductPurchaseDTO) (model.ProductPurchase, string) {
	var product model.ProductPurchase
	msg := ""
	productOutOfStockMsg := "Product is out of stock!"
	if videoGame.Amount < productPurchaseDto.Amount && !videoGame.Digital{
		msg = productOutOfStockMsg
	}
	product.ProductName = videoGame.Name
	product.ProductId = videoGame.Id
	product.Amount = productPurchaseDto.Amount
	product.UserId = userData.Id
	product.TotalPrice = videoGame.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
	return product, msg
}