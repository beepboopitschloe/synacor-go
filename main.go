package main

import (
	"bitbucket.org/nmuth/synacor-go/synacor"
	"bitbucket.org/nmuth/synacor-go/synacor/parser"
	"fmt"
	"io"
	"os"
)

var EXIT_HALT = "program halted"
var EXIT_UNRECOGNIZED_INSTR = "unrecognized instruction"

var fileInput io.Reader
var machine synacor.Machine

func nextOpcode() (synacor.Opcode, error) {
	return parser.NextOpcode(fileInput)
}

func nextRegister() (uint16, error) {
	return parser.NextRegister(fileInput)
}

func nextValue() (uint16, error) {
	return parser.NextValue(fileInput, machine)
}

func doSet(register uint16, value uint16) {
	fmt.Printf("[DEBUG] setting value of register %d to %d\n", register, value)

	machine.Registers[register] = value

	fmt.Printf("[DEBUG] value of register %d is now %d\n", register, machine.Registers[register])
}

func vmHalt() {
	panic(EXIT_HALT)
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
	case synacor.Add:
		register, err := nextRegister()

		a, err := nextValue()
		b, err := nextValue()

		fmt.Printf("[DEBUG] storing into register %d : %d + %d = %d\n",
			register, a, b, (a+b)%32768)

		if err != nil {
			panic(err)
		}

		value := (a + b) % 32768

		doSet(register, value)
	case synacor.Out:
		asciiCode, err := nextValue()

		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%c", asciiCode)
		}
	case synacor.Noop:
		// do nothing
	default:
		panic(EXIT_UNRECOGNIZED_INSTR)
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
			} else if r == EXIT_UNRECOGNIZED_INSTR {
				fmt.Println("user error: unrecognized instruction")
				os.Exit(1)
			} else {
				fmt.Println(r)
				os.Exit(2)
			}
		}
	}()

	fmt.Println("[DEBUG] opening ./testbin")

	f, err := os.Open("testbin")
	defer f.Close()

	fileInput = f

	fmt.Println("[DEBUG] creating machine")

	machine = synacor.Machine{}

	if err != nil {
		panic(err)
	}

	fmt.Println("[DEBUG] begin reading instructions")
	for opcode, err := nextOpcode(); err != io.EOF; opcode, err = nextOpcode() {
		if err != nil {
			fmt.Println("[DEBUG] error reading instruction")
			panic(err)
		}

		opName := opcodeToString(opcode)

		fmt.Printf("[DEBUG] exec %d (%s)\n", opcode, opName)

		execOpcode(opcode)
	}

	fmt.Printf("[DEBUG] reached EOF\n")
}
