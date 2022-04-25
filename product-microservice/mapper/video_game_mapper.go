package mapper

import (
	"product/dto"
	"product/model"

	"time"
)


func ToVideoGame(videoGameDTO dto.VideoGameDTO) (model.VideoGame, error) {
	releaseDate, error := time.Parse("2006-01-02", videoGameDTO.ReleaseDate)
	if error != nil {
		return model.VideoGame{}, error
	}
	
	return model.VideoGame {
		Product: model.Product(videoGameDTO.Product),
		Digital: videoGameDTO.Digital,
		Platform: videoGameDTO.Platform,
		Rating: videoGameDTO.Rating,
		Genre: videoGameDTO.Genre,
		ReleaseDate: releaseDate,
	}, error
}

func ToVideoGameDTO(videoGame model.VideoGame) dto.VideoGameDTO {
	return dto.VideoGameDTO {
		Product: dto.ProductDTO(videoGame.Product),
		Digital: videoGame.Digital,
		Platform: videoGame.Platform,
		Rating: videoGame.Rating,
		Genre: videoGame.Genre,
		ReleaseDate: videoGame.ReleaseDate.Format("2006-01-02"),
	}
}

func ToVideoGameDTOs(videoGames []model.VideoGame) []dto.VideoGameDTO {
	videoGameDTOs := make([]dto.VideoGameDTO, len(videoGames))

	for i, itm := range videoGames {
		videoGameDTOs[i] = ToVideoGameDTO(itm)
	}

	return videoGameDTOs
}