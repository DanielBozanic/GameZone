package model

type Reason int

const (
	OFFENSIVE_LANGUAGE Reason = iota + 1
	INAPPROPRIATE_USERNAME
	HARRASMENT
	SPAM
	MISINFORMATION
)

func (r Reason) String() string {
	return [...]string{
		"Offensive language",
		"Inappropriate language",
		"Harrasment",
		"Spam",
		"Misinformation"}[r-1]
}