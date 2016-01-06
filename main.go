package main

import (
	"bitbucket.org/nmuth/synacor-go/synacor"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

var fileInput io.Reader

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

func nextOpcode() (result synacor.Opcode, err error) {
	num, err := readUint16(fileInput)

	if err == nil {
		result = synacor.Opcode(num)
	}

	return result, err
}

func nextRegister() (result synacor.Register, err error) {
	return result, err
}

func nextValue() (result uint16, err error) {
	return result, err
}

func doSet(register synacor.Register, value uint16) {

}

func vmHalt() {
	panic("program halted")
}

func execOpcode(op synacor.Opcode) {
	switch op {
	case synacor.Halt:
		vmHalt()
	case synacor.Set:
		register, err := nextRegister()
		value, err := nextValue()

		if err != nil {
			panic(err)
		} else {
			doSet(register, value)
		}
	case synacor.Noop:
		// do nothing
	default:
		panic("unrecognized instruction")
	}
}

func opcodeToString(op synacor.Opcode) string {
	if op < synacor.Opcode(synacor.LEN_OPCODE_STRINGS) {
		return synacor.OPCODE_STRINGS[op]
	} else {
		return "unrecognized"
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "program halted" {
				os.Exit(0)
			} else {
				fmt.Println(r)
				os.Exit(1)
			}
		}
	}()

	fileInput, err := os.Open("testbin")

	if err != nil {
		panic(err)
	}

	for opcode, err := nextOpcode(); err != io.EOF; opcode, err = nextOpcode() {
		if err != nil {
			panic(err)
		}

		execOpcode(opcode)

		opName := opcodeToString(opcode)

		fmt.Printf("%d\t\t=> %s\n", opcode, opName)
	}

	defer fileInput.Close()
}
