package machine

import (
	"bitbucket.org/nmuth/synacor-go/synacor/parser"
	"io"
	"log"
	"os"
)

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
