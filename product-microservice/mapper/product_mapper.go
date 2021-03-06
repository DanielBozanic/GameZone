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
		Amount: productDTO.Amount,
		Type: productDTO.Type,
		Image: model.File(productDTO.Image),
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
		Amount: product.Amount,
		Type: product.Type,
		Image: dto.FileDTO(product.Image),
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