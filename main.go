package main

import (
	"bitbucket.org/nmuth/synacor-go/synacor"
	"bitbucket.org/nmuth/synacor-go/synacor/parser"
	"fmt"
	"io"
	"os"
)

var EXIT_HALT = "program halted"

func doSet(register synacor.Register, value uint16) {

}

func vmHalt() {
	panic(EXIT_HALT)
}

func execOpcode(op synacor.Opcode) {
	switch op {
	case synacor.Halt:
		vmHalt()
	case synacor.Set:
		register, err := parser.NextRegister()
		value, err := parser.NextValue()

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
			if r == EXIT_HALT {
				os.Exit(0)
			} else {
				fmt.Println(r)
				os.Exit(1)
			}
		}
	}()

	fileInput, err := os.Open("testbin")
	defer fileInput.Close()

	parser.SetFileInput(fileInput)

	if err != nil {
		panic(err)
	}

	for opcode, err := parser.NextOpcode(); err != io.EOF; opcode, err = parser.NextOpcode() {
		if err != nil {
			panic(err)
		}

		opName := opcodeToString(opcode)

		fmt.Printf("[DEBUG] exec %d (%s)\n", opcode, opName)

		execOpcode(opcode)
	}

	fmt.Printf("[DEBUG] reached EOF\n")
}
