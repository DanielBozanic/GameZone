package model

type Rating int

const (
	ONE_STAR Rating = iota + 1
	TWO_STARS
	THREE_STARS
	FOUR_STARS
	FIVE_STARS
)
