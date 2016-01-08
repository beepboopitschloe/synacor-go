package main

import (
	"bitbucket.org/nmuth/synacor-go/synacor"
	"bitbucket.org/nmuth/synacor-go/synacor/parser"
	"fmt"
	"io"
	"log"
	"os"
)

var EXIT_HALT = "program halted"
var EXIT_UNRECOGNIZED_INSTR = "unrecognized instruction"
var EXIT_NOT_IMPLEMENTED = "not implemented"

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
	log.Printf("[DEBUG] setting value of register %d to %d\n", register, value)

	machine.Registers[register] = value

	log.Printf("[DEBUG] value of register %d is now %d\n", register, machine.Registers[register])
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
	case synacor.Push:
		fallthrough
	case synacor.Pop:
		fallthrough
	case synacor.Eq:
		fallthrough
	case synacor.Gt:
		fallthrough
	case synacor.Jmp:
		fallthrough
	case synacor.Jt:
		fallthrough
	case synacor.Jf:
		panic(EXIT_NOT_IMPLEMENTED)
	case synacor.Add:
		register, err := nextRegister()

		a, err := nextValue()
		b, err := nextValue()

		log.Printf("[DEBUG] storing into register %d : %d + %d = %d\n", register, a, b, (a+b)%32768)

		if err != nil {
			panic(err)
		}

		value := (a + b) % 32768

		doSet(register, value)
	case synacor.Mult:
		fallthrough
	case synacor.Mod:
		fallthrough
	case synacor.And:
		fallthrough
	case synacor.Or:
		fallthrough
	case synacor.Not:
		fallthrough
	case synacor.Rmem:
		fallthrough
	case synacor.Wmem:
		fallthrough
	case synacor.Call:
		fallthrough
	case synacor.Ret:
		panic(EXIT_NOT_IMPLEMENTED)
	case synacor.Out:
		asciiCode, err := nextValue()

		if err != nil {
			panic(err)
		} else {
			// this might be slightly faster if we wrote asciiCode to a buffer instead
			// of using fmt
			fmt.Print(string(asciiCode))
		}
	case synacor.In:
		panic(EXIT_NOT_IMPLEMENTED)
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
				log.Fatalln("user error: unrecognized instruction")
			} else if r == EXIT_NOT_IMPLEMENTED {
				log.Fatalln("system error: instruction not implemented")
			} else {
				log.Panicln(r)
			}
		}
	}()

	output, err := os.Create("output")
	defer output.Close()

	log.SetOutput(output)

	if err != nil {
		panic(err)
	}

	log.Println("[DEBUG] opening ./data/challenge.bin")

	f, err := os.Open("data/challenge.bin")
	defer f.Close()

	fileInput = f

	log.Println("[DEBUG] creating machine")

	machine = synacor.Machine{}

	if err != nil {
		panic(err)
	}

	log.Println("[DEBUG] begin reading instructions")

	for opcode, err := nextOpcode(); err != io.EOF; opcode, err = nextOpcode() {
		if err != nil {
			log.Panicln("[DEBUG] error reading instruction")
		}

		opName := opcodeToString(opcode)

		log.Printf("[DEBUG] exec %d (%s)\n", opcode, opName)

		execOpcode(opcode)
	}

	log.Printf("[DEBUG] reached EOF\n")
}
