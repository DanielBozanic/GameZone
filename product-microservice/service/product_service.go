package service

import (
	"errors"
	"product/model"
	"product/repository"

	"gorm.io/gorm"
)

type productService struct {
	IProductRepository repository.IProductRepository
	IProductPurchaseRepository repository.IProductPurchaseRepository
}

type IProductService interface {
	GetProductById(id int) (model.Product, error)
	SearchByName(page int, pageSize int, name string) ([]model.Product, error)
	GetNumberOfRecordsSearch(name string) int64
	DeleteProduct(id int) string
	GetMainPageProducts() []model.Product
	AddProductToMainPage(productId int) string
	RemoveProductFromMainPage(productId int) string
	GetPopularProducts() []model.Product
	GetRecommendedProducts(userId int) []model.Product
}

func NewProductService(productRepository repository.IProductRepository, productPurchaseRepository repository.IProductPurchaseRepository) IProductService {
	return &productService{IProductRepository: productRepository, IProductPurchaseRepository: productPurchaseRepository}
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

func (productService *productService) DeleteProduct(id int) string {
	product, err := productService.GetProductById(id)
	if err != nil {
		return err.Error()
	}
	_, err = productService.IProductPurchaseRepository.GetUnpaidProductPurchase(product.Id)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "This product is in use"
	}
	*product.Archived = true
	err = productService.IProductRepository.UpdateProduct(product)
	if err != nil {
		return err.Error()
	}
	return ""
}


// Main page related services
func (productService *productService) GetMainPageProducts() []model.Product {
	return productService.IProductRepository.GetMainPageProducts()
}

func (productService *productService) AddProductToMainPage(productId int) string {
	msg := ""
	product, err := productService.IProductRepository.GetProductById(productId)
	if err != nil {
		return err.Error()
	}

	_, err = productService.IProductRepository.IsProductOnMainPage(productId)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "Product is already on main page"
	}

	productsOnMain := productService.IProductRepository.GetMainPageProducts();
	if len(productsOnMain) >= 9 {
		return "The maximum limit of products on main page is 9"
	}
	
	*product.MainPage = true  
	err = productService.IProductRepository.UpdateProduct(product)
	return msg
}

func (productService *productService) RemoveProductFromMainPage(productId int) string {
	msg := ""
	product, err := productService.IProductRepository.GetProductById(productId)
	if err != nil {
		return err.Error()
	}
	
	*product.MainPage = false  
	err = productService.IProductRepository.UpdateProduct(product)
	return msg
}

func (productService *productService) GetPopularProducts() []model.Product {
	return productService.IProductRepository.GetPopularProducts();
}

func (productService *productService) GetRecommendedProducts(userId int) []model.Product {
	return productService.IProductRepository.GetRecommendedProducts(userId)
}