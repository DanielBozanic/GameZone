package model

type Type int

const (
	CONSOLE Type = iota + 1
	GRAPHICS_CARD
	HARD_DISK_DRIVE
	HEADPHONES
	KEYBOARD
	MONITOR
	MOTHERBOARD
	MOUSE
	POWER_SUPPLY_UNIT
	PROCESSOR
	RAM
	SOLID_STATE_DRIVE
	VIDEO_GAME
)

func (t Type) String() string {
	return [...]string{
		"Console",
		"Graphics card",
		"Hard disk drive",
		"Headphones",
		"Keyboard",
		"Monitor",
		"Motherboard",
		"Mouse",
		"Power supply unit",
		"Processor",
		"RAM",
		"Solid state drive",
		"Video game"}[t-1]
}