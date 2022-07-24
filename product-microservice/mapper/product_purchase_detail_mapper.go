package mapper

import (
	"product/dto"
	"product/model"
)


func ToProductPurchaseDetail(productPurchaseDetailDTO dto.ProductPurchaseDetailDTO) model.ProductPurchaseDetail {
		return model.ProductPurchaseDetail {
			Id: productPurchaseDetailDTO.Id,
			ProductPurchaseId: productPurchaseDetailDTO.ProductPurchaseId,
			ProductId: productPurchaseDetailDTO.ProductId,
			ProductName: productPurchaseDetailDTO.ProductName,
			ProductPrice: productPurchaseDetailDTO.ProductPrice,
			ProductQuantity: productPurchaseDetailDTO.ProductQuantity,
		}
}

func ToProductPurchaseDetailDTO(productPurchaseDetail model.ProductPurchaseDetail) dto.ProductPurchaseDetailDTO {
	return dto.ProductPurchaseDetailDTO {
		Id: productPurchaseDetail.Id,
		ProductId: productPurchaseDetail.ProductId,
		ProductPurchaseId: productPurchaseDetail.ProductPurchaseId,
		ProductName: productPurchaseDetail.ProductName,
		ProductPrice: productPurchaseDetail.ProductPrice,
		ProductQuantity: productPurchaseDetail.ProductQuantity,
	}
}

func ToProductPurchaseDetails(productPurchaseDetailDTOs []dto.ProductPurchaseDetailDTO) []model.ProductPurchaseDetail {
	productPurchaseDetails := make([]model.ProductPurchaseDetail, len(productPurchaseDetailDTOs))

	for i, itm := range productPurchaseDetailDTOs {
		productPurchaseDetails[i] = ToProductPurchaseDetail(itm)
	}

	return productPurchaseDetails
}

func ToProductPurchaseDetailDTOs(productPurchaseDetails []model.ProductPurchaseDetail) []dto.ProductPurchaseDetailDTO {
	productPurchaseDetailsDTOs := make([]dto.ProductPurchaseDetailDTO, len(productPurchaseDetails))

	for i, itm := range productPurchaseDetails {
		productPurchaseDetailsDTOs[i] = ToProductPurchaseDetailDTO(itm)
	}

	return productPurchaseDetailsDTOs
}