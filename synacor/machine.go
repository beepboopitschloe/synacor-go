package synacor

const INITIAL_STACK_CAPACITY = 128

type Machine struct {
	Registers [8]uint16
	InstrPtr  uint16
	Memory    [32768]uint16
	stack     []uint16
}

func NewMachine() Machine {
	m := Machine{}

	m.stack = make([]uint16, 0, INITIAL_STACK_CAPACITY)

	return m
}

func (m *Machine) PushStack(val uint16) {
	m.stack = append(m.stack, val)
}

func (m *Machine) PopStack() (result uint16) {
	result = m.stack[len(m.stack)-1]

	m.stack = m.stack[0 : len(m.stack)-1]

	return result
}
