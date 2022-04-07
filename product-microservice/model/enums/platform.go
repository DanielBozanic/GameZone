package enums

type Platform int

const (
	PS5                 Platform = 1
	PS4                 Platform = 2
	XBOX_SERIES_X_AND_S Platform = 3
	XBOXONE             Platform = 4
	NINTENDO_SWITCH     Platform = 5
	PC                  Platform = 6
)

func (platform Platform) String() string {
	switch platform {
	case PS5:
		return "Playstation 5"
	case PS4:
		return "Playstation 4"
	case XBOX_SERIES_X_AND_S:
		return "Xbox Series X/S"
	case XBOXONE:
		return "Xbox One"
	case NINTENDO_SWITCH:
		return "Nintendo Switch"
	case PC:
		return "PC"
	default:
		return ""
	}
}