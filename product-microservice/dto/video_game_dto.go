package dto

type VideoGameDTO struct {
	Product     ProductDTO
	Digital     bool
	Platform    string
	Rating      uint
	Genre       string
	ReleaseDate string
}