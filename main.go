package main

import (
	"bitbucket.org/nmuth/synacor-go/synacor"
	"bitbucket.org/nmuth/synacor-go/synacor/machine"
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
var activeMachine machine.Machine

func nextOpcode() (synacor.Opcode, error) {
	return parser.NextOpcode(fileInput)
}

func nextRegister() (uint16, error) {
	return parser.NextRegister(fileInput)
}

func nextValue() (uint16, error) {
	return parser.NextValue(fileInput, activeMachine.Registers)
}

func doSet(register uint16, value uint16) {
	log.Printf("[DEBUG] setting value of register %d to %d\n", register, value)

	activeMachine.Registers[register] = value

	log.Printf("[DEBUG] value of register %d is now %d\n", register, activeMachine.Registers[register])
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
		val, err := nextValue()

		if err != nil {
			panic(err)
		} else {
			activeMachine.PushStack(val)
		}
	case synacor.Pop:
		register, err := nextRegister()

		if err != nil {
			panic(err)
		} else {
			doSet(register, activeMachine.PopStack())
		}
	case synacor.Eq:
		register, err := nextRegister()
		a, err := nextValue()
		b, err := nextValue()

		if err != nil {
			panic(err)
		} else if a == b {
			doSet(register, 1)
		} else {
			doSet(register, 0)
		}
	case synacor.Gt:
		register, err := nextRegister()
		a, err := nextValue()
		b, err := nextValue()

		if err != nil {
			panic(err)
		} else if a > b {
			doSet(register, 1)
		} else {
			doSet(register, 0)
		}
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
		register, err := nextRegister()
		a, err := nextValue()
		b, err := nextValue()

		if err != nil {
			panic(err)
		} else {
			doSet(register, (a*b)%32768)
		}
	case synacor.Mod:
		register, err := nextRegister()
		a, err := nextValue()
		b, err := nextValue()

		if err != nil {
			panic(err)
		} else {
			doSet(register, a%b)
		}
	case synacor.And:
		register, err := nextRegister()
		a, err := nextValue()
		b, err := nextValue()

		if err != nil {
			panic(err)
		} else {
			doSet(register, a&b)
		}
	case synacor.Or:
		register, err := nextRegister()
		a, err := nextValue()
		b, err := nextValue()

		if err != nil {
			panic(err)
		} else {
			doSet(register, a|b)
		}
	case synacor.Not:
		register, err := nextRegister()
		a, err := nextValue()

		if err != nil {
			panic(err)
		} else {
			doSet(register, a^0xff)
		}
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
			// @TODO it would be better if we wrote asciiCode to a buffer instead of using fmt
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

func execFile(filename string) {
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
}

func main() {
	output, err := os.Create("output")
	defer output.Close()

	if err != nil {
		panic(err)
	}

	log.SetOutput(output)

	log.Println("[DEBUG] creating activeMachine")

	activeMachine = machine.NewMachine()

	filename := os.Args[1]

	activeMachine.LoadFile(filename)

	log.Println("[DEBUG] memory is now:", activeMachine.Memory)
}
