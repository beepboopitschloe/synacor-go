package parser

import (
	"bitbucket.org/nmuth/synacor-go/synacor"
	"encoding/binary"
	"io"
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

func NextOpcode() (result synacor.Opcode, err error) {
	num, err := readUint16(fileInput)

	if err == nil {
		result = synacor.Opcode(num)
	}

	return result, err
}

func NextRegister() (result synacor.Register, err error) {
	return result, err
}

func NextValue() (result uint16, err error) {
	return result, err
}

func SetFileInput(input io.Reader) {
	fileInput = input
}
