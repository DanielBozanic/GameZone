package service

import (
	"product/dto"
	"product/model"
	"product/repository"

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


func (productService *productService) AddProductToCart(productPurchaseDto dto.ProductPurchaseDTO, userData dto.UserData) string {
	var product model.ProductPurchase
	msg := ""
	productOutOfStockMsg := "Product is out of stock!"

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
			if console.Amount < productPurchaseDto.Amount {
				msg = productOutOfStockMsg
			}
			product.ProductId = console.Id
			product.Amount = productPurchaseDto.Amount
			product.UserId = userData.Id
			product.TotalPrice = console.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
			break
		case graphicsCard.Id == productPurchaseDto.ProductId:
			if graphicsCard.Amount < productPurchaseDto.Amount {
				msg = productOutOfStockMsg
			}
			product.ProductId = graphicsCard.Id
			product.Amount = productPurchaseDto.Amount
			product.UserId = userData.Id
			product.TotalPrice = graphicsCard.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
			break
		case hardDisk.Id == productPurchaseDto.ProductId:
			if hardDisk.Amount < productPurchaseDto.Amount {
				msg = productOutOfStockMsg
			}
			product.ProductId = hardDisk.Id
			product.Amount = productPurchaseDto.Amount
			product.UserId = userData.Id
			product.TotalPrice = hardDisk.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
			break
		case headphones.Id == productPurchaseDto.ProductId:
			if headphones.Amount < productPurchaseDto.Amount {
				msg = productOutOfStockMsg
			}
			product.ProductId = headphones.Id
			product.Amount = productPurchaseDto.Amount
			product.UserId = userData.Id
			product.TotalPrice = headphones.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
			break
		case keyboard.Id == productPurchaseDto.ProductId:
			if keyboard.Amount < productPurchaseDto.Amount {
				msg = productOutOfStockMsg
			}
			product.ProductId = keyboard.Id
			product.Amount = productPurchaseDto.Amount
			product.UserId = userData.Id
			product.TotalPrice = keyboard.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
			break
		case monitor.Id == productPurchaseDto.ProductId:
			if monitor.Amount < productPurchaseDto.Amount {
				msg = productOutOfStockMsg
			}
			product.ProductId = monitor.Id
			product.Amount = productPurchaseDto.Amount
			product.UserId = userData.Id
			product.TotalPrice = monitor.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
			break
		case motherboard.Id == productPurchaseDto.ProductId:
			if motherboard.Amount < productPurchaseDto.Amount {
				msg = productOutOfStockMsg
			}
			product.ProductId = motherboard.Id
			product.Amount = productPurchaseDto.Amount
			product.UserId = userData.Id
			product.TotalPrice = motherboard.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
			break
		case mouse.Id == productPurchaseDto.ProductId:
			if mouse.Amount < productPurchaseDto.Amount {
				msg = productOutOfStockMsg
			}
			product.ProductId = mouse.Id
			product.Amount = productPurchaseDto.Amount
			product.UserId = userData.Id
			product.TotalPrice = mouse.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
			break
		case psu.Id == productPurchaseDto.ProductId:
			if psu.Amount < productPurchaseDto.Amount {
				msg = productOutOfStockMsg
			}
			product.ProductId = psu.Id
			product.Amount = productPurchaseDto.Amount
			product.UserId = userData.Id
			product.TotalPrice = psu.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
			break
		case processor.Id == productPurchaseDto.ProductId:
			if processor.Amount < productPurchaseDto.Amount {
				msg = productOutOfStockMsg
			}
			product.ProductId = processor.Id
			product.Amount = productPurchaseDto.Amount
			product.UserId = userData.Id
			product.TotalPrice = processor.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
			break
		case ram.Id == productPurchaseDto.ProductId:
			if ram.Amount < productPurchaseDto.Amount {
				msg = productOutOfStockMsg
			}
			product.ProductId = ram.Id
			product.Amount = productPurchaseDto.Amount
			product.UserId = userData.Id
			product.TotalPrice = ram.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
			break
		case ssd.Id == productPurchaseDto.ProductId:
			if ssd.Amount < productPurchaseDto.Amount {
				msg = productOutOfStockMsg
			}
			product.ProductId = ssd.Id
			product.Amount = productPurchaseDto.Amount
			product.UserId = userData.Id
			product.TotalPrice = ssd.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
			break
		case videoGame.Id == productPurchaseDto.ProductId:
			if videoGame.Amount < productPurchaseDto.Amount && !videoGame.Digital{
				msg = productOutOfStockMsg
			}
			product.ProductId = videoGame.Id
			product.Amount = productPurchaseDto.Amount
			product.UserId = userData.Id
			product.TotalPrice = videoGame.Price.Mul(decimal.NewFromInt(int64(productPurchaseDto.Amount)))
			break
		default:
			msg = "Product not found!"
	}

	if msg == "" {
		productService.IProductRepository.AddPurchase(product);
	}
	return msg
}