package dto

type SolidStateDriveDTO struct {
	Product            ProductDTO
	Capacity           string
	Interface          string
	MaxSequentialRead  string
	MaxSequentialWrite string
	Form               string
	Dimensions         string
}