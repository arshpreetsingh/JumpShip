package main

import (
	"os"
	"parking_lot/ishell"
)

func main() {
	// Process the input file commands only
	if len(os.Args) > 1 && os.Args[1] != "" {
		ishell.ProcessFile(os.Args[1])
		return
	}
	// Start interactive shell
	ishell.Start()
}
