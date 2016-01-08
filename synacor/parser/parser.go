package parser

import (
	"bitbucket.org/nmuth/synacor-go/synacor/opcode"
	"encoding/binary"
	"errors"
	"io"
)

// Registers 0-7 are denoted by numbers in 32768..32775.
var REGISTER_START = uint16(32768)
var REGISTER_END = uint16(32775)

var ERROR_REGISTER_OUT_OF_BOUNDS = "register out of bounds"

// Read an unsigned 16-bit integer from the file input.
func NextCodepoint(r io.Reader) (result uint16, err error) {
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

// Read a register address from the file input.
func NextOpcode(fileInput io.Reader) (result opcode.Opcode, err error) {
	num, err := NextCodepoint(fileInput)

	if err == nil {
		result = opcode.Opcode(num)
	}

	return result, err
}

// Tell whether the given number is a register address.
func isRegister(num uint16) (result bool) {
	return (num >= REGISTER_START) && (num <= REGISTER_END)
}

func numToRegister(num uint16) (result uint16, err error) {
	result = num - REGISTER_START

	if result > 7 {
		err = errors.New(ERROR_REGISTER_OUT_OF_BOUNDS)
	}

	return result, err
}

// Read a register address from the file input.
func NextRegister(fileInput io.Reader) (result uint16, err error) {
	num, err := NextCodepoint(fileInput)

	if isRegister(num) {
		result, err = numToRegister(num)
	} else {
		err = errors.New(ERROR_REGISTER_OUT_OF_BOUNDS)
	}

	return result, err
}

// Get the next value from the file input. If the value is an integer literal,
// return the literal. If the value is a register, return the current value of
// the register.
func NextValue(fileInput io.Reader, registers [8]uint16) (result uint16, err error) {
	num, err := NextCodepoint(fileInput)

	if err == nil && isRegister(num) {
		register, err := numToRegister(num)

		if err != nil {
			panic(err)
		} else {
			result = registers[register]
		}
	} else {
		result = num
	}

	return result, err
}
