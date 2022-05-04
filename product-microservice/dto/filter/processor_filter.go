package filter

type ProcessorFilter struct {
	Manufacturers []string
	Types         []string
	Sockets       []string
	NumberOfCores []uint
	Threads       []uint
}