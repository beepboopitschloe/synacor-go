package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func readUint16(r io.Reader, b byte) (result uint16, err error) {
	var lo, hi byte

	err := binary.Read(r, binary.LittleEndian, &lo)
	err := binary.Read(r, binary.LittleEndian, &hi)

	if err != nil {
		return 0, err
	}
}

func main() {
	fmt.Println("Hello world")

	f, err := os.Open("data/challenge.bin")

	if err != nil {
		panic(err)
	}

	for i := 0; i < 16; i++ {
		var b byte

		err = binary.Read(f, binary.LittleEndian, &b)

		fmt.Println(b)
	}

	defer f.Close()
}
