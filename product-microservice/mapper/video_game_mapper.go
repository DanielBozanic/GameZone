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
		Name: videoGameDTO.Name, 
		Price: videoGameDTO.Price, 
		Digital: videoGameDTO.Digital,
		Platform: videoGameDTO.Platform,
		Publisher: videoGameDTO.Publisher,
		Rating: videoGameDTO.Rating,
		Genre: videoGameDTO.Genre,
		ReleaseDate: releaseDate,
		Amount: videoGameDTO.Amount,
	}, error
}

func ToVideoGameDTO(videoGame model.VideoGame) dto.VideoGameDTO {
	return dto.VideoGameDTO {
		Id: videoGame.Id, 
		Name: videoGame.Name,
		Price: videoGame.Price, 
		Digital: videoGame.Digital,
		Platform: videoGame.Platform,
		Publisher: videoGame.Publisher,
		Rating: videoGame.Rating,
		Genre: videoGame.Genre,
		ReleaseDate: videoGame.ReleaseDate.Format("2006-01-02"),
		Amount: videoGame.Amount,
	}
}

func ToVideoGameDTOs(videoGames []model.VideoGame) []dto.VideoGameDTO {
	videoGameDTOs := make([]dto.VideoGameDTO, len(videoGames))

	for i, itm := range videoGames {
		videoGameDTOs[i] = ToVideoGameDTO(itm)
	}

	return videoGameDTOs
}