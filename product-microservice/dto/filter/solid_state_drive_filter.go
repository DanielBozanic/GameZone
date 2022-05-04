package filter

type SolidStateDriveFilter struct {
	Manufacturers       []string
	Capacities          []string
	MaxSequentialReads  []string
	MaxSequentialWrites []string
	Forms               []string
}