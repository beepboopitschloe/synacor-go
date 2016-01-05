package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func readUint16(r io.Reader) (result uint16, err error) {
	var lo, hi byte
	var bytes []byte

	err = binary.Read(r, binary.LittleEndian, &lo)

	if err != nil {
		err = binary.Read(r, binary.LittleEndian, &hi)
	}

	bytes = append(bytes, lo)
	bytes = append(bytes, hi)

	result = binary.LittleEndian.Uint16(bytes)

	return result, err
}

func main() {
	f, err := os.Open("data/challenge.bin")

	if err != nil {
		panic(err)
	}

	for i := 0; i < 4; i++ {
		var num uint16

		num, err = readUint16(f)

		if err == nil {
			fmt.Println(num)
		} else {
			panic(err)
		}
	}

	defer f.Close()
}
