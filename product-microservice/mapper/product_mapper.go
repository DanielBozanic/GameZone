package mapper

import (
	"product/dto"
	"product/model"
)

func ToProduct(productDTO dto.ProductDTO) (model.Product) {
	return model.Product {
		Id: productDTO.Id,
		Name: productDTO.Name,
		Description: productDTO.Description,
		Manufacturer: productDTO.Manufacturer,
		Price: productDTO.Price,
		Quantity: productDTO.Quantity,
		Type: productDTO.Type,
		Image: model.File(productDTO.Image),
		MainPage: productDTO.MainPage,
		Archived: productDTO.Archived,
	}
}

func ToProductDTO(product model.Product) dto.ProductDTO {
	return dto.ProductDTO {
		Id: product.Id,
		Name: product.Name,
		Description: product.Description,
		Manufacturer: product.Manufacturer,
		Price: product.Price,
		Quantity: product.Quantity,
		Type: product.Type,
		Image: dto.FileDTO(product.Image),
		MainPage: product.MainPage,
		Archived: product.Archived,
	}
}

func ToProductDTOs(products []model.Product) []dto.ProductDTO {
	productsDTOs := make([]dto.ProductDTO, len(products))

	for i, itm := range products {
		productsDTOs[i] = ToProductDTO(itm)
	}

	return productsDTOs
}