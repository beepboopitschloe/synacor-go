package main

import (
	"bitbucket.org/nmuth/synacor-go/synacor/machine"
	"log"
	"os"
)

func main() {
	output, err := os.Create("output")
	defer output.Close()

	if err != nil {
		panic(err)
	}

	log.SetOutput(output)

	log.Println("[DEBUG] creating activeMachine")

	m := machine.NewMachine()

	filename := os.Args[1]

	m.LoadFile(filename)

	m.Execute()
}
