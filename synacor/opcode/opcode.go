package opcode

type Opcode uint16

var OPCODE_STRINGS = [...]string{
	"halt",
	"set",
	"push",
	"pop",
	"eq",
	"gt",
	"jmp",
	"jt",
	"jf",
	"add",
	"mult",
	"mod",
	"and",
	"or",
	"not",
	"rmem",
	"wmem",
	"call",
	"ret",
	"out",
	"in",
	"noop",
}

var LEN_OPCODE_STRINGS = len(OPCODE_STRINGS)

const (
	Halt = iota
	Set
	Push
	Pop
	Eq
	Gt
	Jmp
	Jt
	Jf
	Add
	Mult
	Mod
	And
	Or
	Not
	Rmem
	Wmem
	Call
	Ret
	Out
	In
	Noop
)

func OpcodeToString(op Opcode) string {
	if op < Opcode(LEN_OPCODE_STRINGS) {
		return OPCODE_STRINGS[op]
	} else {
		return "unrecognized"
	}
}
