package machine

import (
	"bitbucket.org/nmuth/synacor-go/synacor/opcode"
	"bitbucket.org/nmuth/synacor-go/synacor/parser"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// Registers 0-7 are denoted by numbers in 32768..32775.
const REGISTER_START = uint16(32768)
const REGISTER_END = uint16(32775)

const ERROR_REGISTER_OUT_OF_BOUNDS = "register out of bounds"

const EXIT_HALT = "EXIT_HALT"
const EXIT_UNRECOGNIZED_INSTR = "EXIT_UNRECOGNIZED_INSTR"
const EXIT_NOT_IMPLEMENTED = "EXIT_NOT_IMPLEMENTED"

const INITIAL_STACK_CAPACITY = 128

type Machine struct {
	Registers [8]uint16
	InstrPtr  uint16
	Memory    [32768]uint16
	stack     []uint16
}

// create a new virtual machine
func NewMachine() Machine {
	m := Machine{}

	m.stack = make([]uint16, 0, INITIAL_STACK_CAPACITY)

	return m
}

// push a value onto the stack
func (m *Machine) PushStack(val uint16) {
	m.stack = append(m.stack, val)
}

// pop a value off the stack
func (m *Machine) PopStack() (result uint16) {
	result = m.stack[len(m.stack)-1]

	m.stack = m.stack[0 : len(m.stack)-1]

	return result
}

// Tell whether the given number is a register address.
func isRegister(num uint16) (result bool) {
	return (num >= REGISTER_START) && (num <= REGISTER_END)
}

// Convert a number to a register index.
func numToRegister(num uint16) (result uint16, err error) {
	result = num - REGISTER_START

	if result > 7 {
		err = errors.New(ERROR_REGISTER_OUT_OF_BOUNDS)
	}

	return result, err
}

// read a register index from memory and increment the instruction pointer
func (m *Machine) NextRegister() (result uint16) {
	num := m.Memory[m.InstrPtr]

	m.InstrPtr += 1

	result, err := numToRegister(num)

	if err != nil {
		log.Println("[ERROR] tried to use %d as a register index\n", result)
		panic(err)
	}

	return result
}

// read a value from memory and increment the instruction pointer
func (m *Machine) NextValue() (result uint16) {
	log.Printf("reading from memory, instrptr is %d\n", m.InstrPtr)
	result = m.Memory[m.InstrPtr]

	m.InstrPtr += 1

	if isRegister(result) {
		register, err := numToRegister(result)
		result = m.Registers[register]

		if err != nil {
			log.Println("[ERROR] tried to use %d as a register index\n", result)
			panic(err)
		}

		log.Printf("reading value %d of register %d\n", result, register)
	}

	log.Printf("done reading from memory, instrptr is %d\n", m.InstrPtr)

	return result
}

// execute memory, starting from address 0
func (m *Machine) Execute() {
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

	m.InstrPtr = 0

	for {
		m.ExecuteNextInstruction()
	}
}

func (m *Machine) ExecuteNextInstruction() {
	op := opcode.Opcode(m.Memory[m.InstrPtr])
	m.InstrPtr += 1

	log.Printf("[DEBUG] executing %s (%d)\n", opcode.OpcodeToString(op), op)

	switch op {
	case opcode.Halt:
		panic(EXIT_HALT)
	case opcode.Set:
		m.DoSet()
	case opcode.Push:
		m.DoPush()
	case opcode.Pop:
		m.DoPop()
	case opcode.Eq:
		m.DoEq()
	case opcode.Gt:
		m.DoGt()
	case opcode.Jmp:
		m.DoJmp()
	case opcode.Jt:
		m.DoJt()
	case opcode.Jf:
		m.DoJf()
	case opcode.Add:
		m.DoAdd()
	case opcode.Mult:
		m.DoMult()
	case opcode.Mod:
		m.DoMod()
	case opcode.And:
		m.DoAnd()
	case opcode.Or:
		m.DoOr()
	case opcode.Not:
		m.DoNot()
	case opcode.Rmem:
		m.DoRmem()
	case opcode.Wmem:
		m.DoWmem()
	case opcode.Call:
		m.DoCall()
	case opcode.Ret:
		m.DoRet()
	case opcode.Out:
		m.DoOut()
	case opcode.In:
		m.DoIn()
	case opcode.Noop:
		// nothing
	default:
		panic(EXIT_UNRECOGNIZED_INSTR)
	}
}

func (m *Machine) DoSet() {
	reg := m.NextRegister()
	val := m.NextValue()

	log.Printf("[DEBUG] setting register %d to value %d\n", reg, val)

	m.Registers[reg] = val
}

func (m *Machine) DoPush() {
	val := m.NextValue()

	m.PushStack(val)
}

func (m *Machine) DoPop() {
	reg := m.NextRegister()

	m.Registers[reg] = m.PopStack()
}

func (m *Machine) DoEq() {
	reg := m.NextRegister()
	a := m.NextValue()
	b := m.NextValue()

	if a == b {
		m.Registers[reg] = 1
	} else {
		m.Registers[reg] = 0
	}
}

func (m *Machine) DoGt() {
	reg := m.NextRegister()
	a := m.NextValue()
	b := m.NextValue()

	if a > b {
		m.Registers[reg] = 1
	} else {
		m.Registers[reg] = 0
	}
}

func (m *Machine) DoJmp() {
	panic(EXIT_NOT_IMPLEMENTED)
}

func (m *Machine) DoJt() {
	panic(EXIT_NOT_IMPLEMENTED)
}

func (m *Machine) DoJf() {
	panic(EXIT_NOT_IMPLEMENTED)
}

func (m *Machine) DoAdd() {
	reg := m.NextRegister()
	a := m.NextValue()
	b := m.NextValue()

	value := (a + b) % 32768

	log.Printf("[DEBUG] storing (%d + %d = %d) into register %d\n", a, b, value, reg)

	m.Registers[reg] = value
}

func (m *Machine) DoMult() {
	reg := m.NextRegister()
	a := m.NextValue()
	b := m.NextValue()

	value := (a * b) % 32768

	m.Registers[reg] = value
}

func (m *Machine) DoMod() {
	reg := m.NextRegister()
	a := m.NextValue()
	b := m.NextValue()

	value := (a % b) % 32768

	m.Registers[reg] = value
}

func (m *Machine) DoAnd() {
	reg := m.NextRegister()
	a := m.NextValue()
	b := m.NextValue()

	value := (a & b) % 32768

	m.Registers[reg] = value
}

func (m *Machine) DoOr() {
	reg := m.NextRegister()
	a := m.NextValue()
	b := m.NextValue()

	value := (a | b) % 32768

	m.Registers[reg] = value
}

func (m *Machine) DoNot() {
	reg := m.NextRegister()
	a := m.NextValue()

	value := (a ^ 0xff) % 32768

	m.Registers[reg] = value
}

func (m *Machine) DoRmem() {
	panic(EXIT_NOT_IMPLEMENTED)
}

func (m *Machine) DoWmem() {
	panic(EXIT_NOT_IMPLEMENTED)
}

func (m *Machine) DoCall() {
	panic(EXIT_NOT_IMPLEMENTED)
}

func (m *Machine) DoRet() {
	panic(EXIT_NOT_IMPLEMENTED)
}

func (m *Machine) DoOut() {
	asciiCode := m.NextValue()

	// @TODO this should be a method on Machine
	// @TODO this should use a buffer
	fmt.Print(string(asciiCode))
}

func (m *Machine) DoIn() {
	panic(EXIT_NOT_IMPLEMENTED)
}

// load an executable binary into memory
func (m *Machine) LoadFile(filename string) {
	log.Println("[DEBUG] opening", filename)

	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	log.Println("[DEBUG] begin reading file")

	memoryIndex := 0

	for {
		codepoint, err := parser.NextCodepoint(f)

		if err != nil && err != io.EOF {
			log.Panicln("[DEBUG] error reading file")
		}

		m.Memory[memoryIndex] = codepoint

		memoryIndex += 1

		if err == io.EOF {
			log.Println("[DEBUG] reached EOF")
			break
		}
	}
}
