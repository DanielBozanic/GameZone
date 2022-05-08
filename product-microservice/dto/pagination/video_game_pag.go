package pagination

import "product/dto"

type VideoGamesPag struct {
	VideoGames []dto.VideoGameDTO
	PageCount int
}