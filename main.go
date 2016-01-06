package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type opcode uint16

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

func readUint16(r io.Reader) (result uint16, err error) {
	var lo, hi byte
	var bytes []byte

	err = binary.Read(r, binary.LittleEndian, &lo)

	if err == nil {
		err = binary.Read(r, binary.LittleEndian, &hi)
	}

	bytes = append(bytes, lo)
	bytes = append(bytes, hi)

	result = binary.LittleEndian.Uint16(bytes)

	return result, err
}

func opcodeToString(op uint16) string {
	if op < uint16(LEN_OPCODE_STRINGS) {
		return OPCODE_STRINGS[op]
	} else {
		return "unrecognized"
	}
}

func main() {
	f, err := os.Open("data/challenge.bin")

	if err != nil {
		panic(err)
	}

	for i := 0; i < 16; i++ {
		var num uint16

		num, err = readUint16(f)

		op := opcodeToString(num)

		if err == nil {
			fmt.Printf("%d\t\t=> %s\n", num, op)
		} else {
			panic(err)
		}
	}

	defer f.Close()
}
