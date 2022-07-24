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