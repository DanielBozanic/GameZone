package mapper

import (
	"news/dto"
	"news/model"
)


func ToNewsComment(newsCommentDTO dto.NewsCommentDTO) (model.NewsComment) {
	return model.NewsComment {
		Id: newsCommentDTO.Id,
		NewsArticle: ToNewsArticle(newsCommentDTO.NewsArticle),
		Comment: newsCommentDTO.Comment,
		UserId: newsCommentDTO.UserId,
		Username: newsCommentDTO.Username,
		DateTime: newsCommentDTO.DateTime,
		Archived: newsCommentDTO.Archived,
	}
}

func ToNewsCommentDTO(newsComment model.NewsComment) dto.NewsCommentDTO {
	return dto.NewsCommentDTO {
		Id: newsComment.Id,
		NewsArticle: ToNewsArticleDTO(newsComment.NewsArticle),
		Comment: newsComment.Comment,
		UserId: newsComment.UserId,
		Username: newsComment.Username,
		DateTime: newsComment.DateTime,
		Archived: newsComment.Archived,
	}
}

func ToNewsCommentDTOs(newsComments []model.NewsComment) []dto.NewsCommentDTO {
	newsCommentsDTOs := make([]dto.NewsCommentDTO, len(newsComments))

	for i, itm := range newsComments {
		newsCommentsDTOs[i] = ToNewsCommentDTO(itm)
	}

	return newsCommentsDTOs
}