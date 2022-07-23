package mapper

import (
	"product/dto"
	"product/model"
)


func ToVideoGame(videoGameDTO dto.VideoGameDTO) (model.VideoGame) {
	return model.VideoGame {
		Product: ToProduct(videoGameDTO.Product),
		Digital: videoGameDTO.Digital,
		Platform: videoGameDTO.Platform,
		Rating: videoGameDTO.Rating,
		Genre: videoGameDTO.Genre,
		ReleaseDate: videoGameDTO.ReleaseDate,
	}
}

func ToVideoGameDTO(videoGame model.VideoGame) dto.VideoGameDTO {
	return dto.VideoGameDTO {
		Product: ToProductDTO(videoGame.Product),
		Digital: videoGame.Digital,
		Platform: videoGame.Platform,
		Rating: videoGame.Rating,
		Genre: videoGame.Genre,
		ReleaseDate: videoGame.ReleaseDate,
	}
}

func ToVideoGameDTOs(videoGames []model.VideoGame) []dto.VideoGameDTO {
	videoGameDTOs := make([]dto.VideoGameDTO, len(videoGames))

	for i, itm := range videoGames {
		videoGameDTOs[i] = ToVideoGameDTO(itm)
	}

	return videoGameDTOs
}