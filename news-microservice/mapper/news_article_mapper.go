package mapper

import (
	"news/dto"
	"news/model"
)


func ToNewsArticle(newsArticleDTO dto.NewsArticleDTO) (model.NewsArticle) {
	return model.NewsArticle {
		Id: newsArticleDTO.Id,
		Title: newsArticleDTO.Title,
		Description: newsArticleDTO.Description,
		UnpublishedContent: newsArticleDTO.UnpublishedContent,
		PublishedContent: newsArticleDTO.PublishedContent,
		DateTime: newsArticleDTO.DateTime,
		IsSent: newsArticleDTO.IsSent,
		Archived: newsArticleDTO.Archived,
	}
}

func ToNewsArticleDTO(newsArticle model.NewsArticle) dto.NewsArticleDTO {
	return dto.NewsArticleDTO {
		Id: newsArticle.Id,
		Title: newsArticle.Title,
		Description: newsArticle.Description,
		UnpublishedContent: newsArticle.UnpublishedContent,
		PublishedContent: newsArticle.PublishedContent,
		DateTime: newsArticle.DateTime,
		IsSent: newsArticle.IsSent,
		Archived: newsArticle.Archived,
	}
}

func ToNewsArticleDTOs(newsArticles []model.NewsArticle) []dto.NewsArticleDTO {
	newsArticleDTOs := make([]dto.NewsArticleDTO, len(newsArticles))

	for i, itm := range newsArticles {
		newsArticleDTOs[i] = ToNewsArticleDTO(itm)
	}

	return newsArticleDTOs
}