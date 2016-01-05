package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type opcode uint16

func readUint16(r io.Reader) (result uint16, err error) {
	var lo, hi byte
	var bytes []byte

	err = binary.Read(r, binary.LittleEndian, &lo)

	if err != nil {
		err = binary.Read(r, binary.LittleEndian, &hi)
	}

	fmt.Printf("--> lo is %b\n", lo)
	fmt.Printf("--> hi is %b\n", hi)

	bytes = append(bytes, lo)
	bytes = append(bytes, hi)

	result = binary.LittleEndian.Uint16(bytes)

	return result, err
}

func opcodeToString(op uint16) string {
	switch op {
	case 21:
		return "noop"
	}
	return "unrecognized"
}

func main() {
	f, err := os.Open("data/challenge.bin")

	if err != nil {
		panic(err)
	}

	for i := 0; i < 4; i++ {
		var num uint16

		num, err = readUint16(f)

		op := opcodeToString(num)

		if err == nil {
			fmt.Printf("%d\t\t=> %s\n", num, op)
		} else {
			panic(err)
		}
	}

	defer f.Close()
}
