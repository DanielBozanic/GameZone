package mapper

import (
	"product/dto"
	"product/model"
	"time"
)


func ToProductPurchase(productPurchaseDTO dto.ProductPurchaseDTO) (model.ProductPurchase, error) {
	if productPurchaseDTO.PurchaseDate != "" {
		purchaseDate, error := time.Parse("2006-01-02", productPurchaseDTO.PurchaseDate)
		if error != nil {
			return model.ProductPurchase{}, error
		}
		return model.ProductPurchase {
			Id: productPurchaseDTO.Id,
			UserId: productPurchaseDTO.UserId,
			Product: ToProduct(productPurchaseDTO.Product),
			Amount: productPurchaseDTO.Amount,
			TotalPrice: productPurchaseDTO.TotalPrice,
			PurchaseDate: purchaseDate,
			DeliveryAddress: productPurchaseDTO.DeliveryAddress,
			City: productPurchaseDTO.City,
			MobilePhoneNumber: productPurchaseDTO.MobilePhoneNumber,
			TypeOfPayment: productPurchaseDTO.TypeOfPayment,
			IsPaidFor: productPurchaseDTO.IsPaidFor,
		}, nil
	} else {
		purchaseDate := time.Time{}
		return model.ProductPurchase {
			Id: productPurchaseDTO.Id,
			UserId: productPurchaseDTO.UserId,
			Product: ToProduct(productPurchaseDTO.Product),
			Amount: productPurchaseDTO.Amount,
			TotalPrice: productPurchaseDTO.TotalPrice,
			PurchaseDate: purchaseDate,
			DeliveryAddress: productPurchaseDTO.DeliveryAddress,
			City: productPurchaseDTO.City,
			MobilePhoneNumber: productPurchaseDTO.MobilePhoneNumber,
			TypeOfPayment: productPurchaseDTO.TypeOfPayment,
			IsPaidFor: productPurchaseDTO.IsPaidFor,
		}, nil
	}
}

func ToProductPurchaseDTO(productPurchase model.ProductPurchase) dto.ProductPurchaseDTO {
	return dto.ProductPurchaseDTO {
		Id: productPurchase.Id,
		UserId: productPurchase.UserId,
		Product: ToProductDTO(productPurchase.Product),
		TotalPrice: productPurchase.TotalPrice,
		Amount: productPurchase.Amount,
		PurchaseDate: productPurchase.PurchaseDate.Format("2006-01-02"),
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