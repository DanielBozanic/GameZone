package mapper

import (
	"news/dto"
	"news/model"
)


func ToNewsArticle(newsArticleDTO dto.NewsArticleDTO) (model.NewsArticle) {
	return model.NewsArticle {
		Id: newsArticleDTO.Id,
		UnpublishedTitle: newsArticleDTO.UnpublishedTitle,
		UnpublishedDescription: newsArticleDTO.UnpublishedDescription,
		UnpublishedContent: newsArticleDTO.UnpublishedContent,
		PublishedTitle: newsArticleDTO.PublishedTitle,
		PublishedDescription: newsArticleDTO.PublishedDescription,
		PublishedContent: newsArticleDTO.PublishedContent,
		DateTime: newsArticleDTO.DateTime,
		IsSent: newsArticleDTO.IsSent,
		Archived: newsArticleDTO.Archived,
	}
}

func ToNewsArticleDTO(newsArticle model.NewsArticle) dto.NewsArticleDTO {
	return dto.NewsArticleDTO {
		Id: newsArticle.Id,
		UnpublishedTitle: newsArticle.UnpublishedTitle,
		UnpublishedDescription: newsArticle.UnpublishedDescription,
		UnpublishedContent: newsArticle.UnpublishedContent,
		PublishedTitle: newsArticle.PublishedTitle,
		PublishedDescription: newsArticle.PublishedDescription,
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