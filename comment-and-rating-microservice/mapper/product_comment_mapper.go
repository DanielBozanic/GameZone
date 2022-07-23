package mapper

import (
	"comment-and-rating/dto"
	"comment-and-rating/model"
)


func ToProductComment(productCommentDTO dto.ProductCommentDTO) (model.ProductComment) {
	return model.ProductComment {
		Id: productCommentDTO.Id,
		ProductName: productCommentDTO.ProductName,
		Username: productCommentDTO.Username,
		Comment: productCommentDTO.Comment,
		Rating: productCommentDTO.Rating,
		Archived: productCommentDTO.Archived,
		DateTime: productCommentDTO.DateTime,
	}
}

func ToProductCommentDTO(productComment model.ProductComment) dto.ProductCommentDTO {
	return dto.ProductCommentDTO {
		Id: productComment.Id,
		ProductName: productComment.ProductName,
		Username: productComment.Username,
		Comment: productComment.Comment,
		Rating: productComment.Rating,
		Archived: productComment.Archived,
		DateTime: productComment.DateTime,
	}
}

func ToProductCommentDTOs(productComments []model.ProductComment) []dto.ProductCommentDTO {
	productCommentDTOs := make([]dto.ProductCommentDTO, len(productComments))

	for i, itm := range productComments {
		productCommentDTOs[i] = ToProductCommentDTO(itm)
	}

	return productCommentDTOs
}