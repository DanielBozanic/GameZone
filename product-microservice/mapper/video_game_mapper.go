package mapper

import (
	"product/dto"
	"product/model"
)


func ToVideoGame(videoGameDTO dto.VideoGameDTO) (model.VideoGame) {


	
	return model.VideoGame {
		Product: model.Product(videoGameDTO.Product),
		Digital: videoGameDTO.Digital,
		Platform: videoGameDTO.Platform,
		Rating: videoGameDTO.Rating,
		Genre: videoGameDTO.Genre,
		ReleaseDate: videoGameDTO.ReleaseDate,
	}
}

func ToVideoGameDTO(videoGame model.VideoGame) dto.VideoGameDTO {
	return dto.VideoGameDTO {
		Product: dto.ProductDTO(videoGame.Product),
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