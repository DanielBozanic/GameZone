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
			ProductId: productPurchaseDTO.ProductId,
			ProductName: productPurchaseDTO.ProductName,
			ProductPrice: productPurchaseDTO.ProductPrice,
			Amount: productPurchaseDTO.Amount,
			TotalPrice: productPurchaseDTO.TotalPrice,
			PurchaseDate: purchaseDate,
			DeliveryAddress: productPurchaseDTO.DeliveryAddress,
			TypeOfPayment: productPurchaseDTO.TypeOfPayment,
		}, nil
	} else {
		purchaseDate := time.Time{}
		return model.ProductPurchase {
			Id: productPurchaseDTO.Id,
			UserId: productPurchaseDTO.UserId,
			ProductId: productPurchaseDTO.ProductId,
			ProductName: productPurchaseDTO.ProductName,
			ProductPrice: productPurchaseDTO.ProductPrice,
			Amount: productPurchaseDTO.Amount,
			TotalPrice: productPurchaseDTO.TotalPrice,
			PurchaseDate: purchaseDate,
			DeliveryAddress: productPurchaseDTO.DeliveryAddress,
			TypeOfPayment: productPurchaseDTO.TypeOfPayment,
		}, nil
	}
}

func ToProductPurchaseDTO(productPurchase model.ProductPurchase) dto.ProductPurchaseDTO {
	return dto.ProductPurchaseDTO {
		Id: productPurchase.Id,
		UserId: productPurchase.UserId,
		ProductId: productPurchase.ProductId,
		ProductName: productPurchase.ProductName,
		ProductPrice: productPurchase.ProductPrice,
		TotalPrice: productPurchase.TotalPrice,
		Amount: productPurchase.Amount,
		PurchaseDate: productPurchase.PurchaseDate.Format("2006-01-02"),
		DeliveryAddress: productPurchase.DeliveryAddress,
		TypeOfPayment: productPurchase.TypeOfPayment,
	}
}

func ToProductPurchaseDTOs(productPurchases []model.ProductPurchase) []dto.ProductPurchaseDTO {
	productPurchaseDTOs := make([]dto.ProductPurchaseDTO, len(productPurchases))

	for i, itm := range productPurchases {
		productPurchaseDTOs[i] = ToProductPurchaseDTO(itm)
	}

	return productPurchaseDTOs
}