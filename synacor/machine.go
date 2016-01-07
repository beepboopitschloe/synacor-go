package synacor

type Register uint16

const (
	A = iota
	B
	C
	D
	E
	F
	G
	H
)

type Machine struct {
	registers [8]Register
	instrPtr  uint16
}
