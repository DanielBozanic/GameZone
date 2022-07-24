package mapper

import (
	"product/dto"
	"product/model"
)


func ToProductPurchase(productPurchaseDTO dto.ProductPurchaseDTO) model.ProductPurchase {
		return model.ProductPurchase {
			Id: productPurchaseDTO.Id,
			UserId: productPurchaseDTO.UserId,
			ProductPurchaseDetail: ToProductPurchaseDetails(productPurchaseDTO.ProductPurchaseDetail),
			TotalPrice: productPurchaseDTO.TotalPrice,
			PurchaseDate: productPurchaseDTO.PurchaseDate,
			DeliveryAddress: productPurchaseDTO.DeliveryAddress,
			City: productPurchaseDTO.City,
			MobilePhoneNumber: productPurchaseDTO.MobilePhoneNumber,
			TypeOfPayment: productPurchaseDTO.TypeOfPayment,
			IsPaidFor: productPurchaseDTO.IsPaidFor,
		}
}

func ToProductPurchaseDTO(productPurchase model.ProductPurchase) dto.ProductPurchaseDTO {
	return dto.ProductPurchaseDTO {
		Id: productPurchase.Id,
		UserId: productPurchase.UserId,
		ProductPurchaseDetail: ToProductPurchaseDetailDTOs(productPurchase.ProductPurchaseDetail),
		TotalPrice: productPurchase.TotalPrice,
		PurchaseDate: productPurchase.PurchaseDate,
		DeliveryAddress: productPurchase.DeliveryAddress,
		City: productPurchase.City,
		MobilePhoneNumber: productPurchase.MobilePhoneNumber,
		TypeOfPayment: productPurchase.TypeOfPayment,
		IsPaidFor: productPurchase.IsPaidFor,
	}
}

func ToProductPurchaseDTOs(productPurchases []model.ProductPurchase) []dto.ProductPurchaseDTO {
	productPurchaseDTOs := make([]dto.ProductPurchaseDTO, len(productPurchases))

	for i, itm := range productPurchases {
		productPurchaseDTOs[i] = ToProductPurchaseDTO(itm)
	}

	return productPurchaseDTOs
}