package main

import (
	"fmt"
	"os"
)

func main() {
	// The full path to execute.
	fmt.Println("Command: ", os.Args[0])

	// Get the argument input in console
	fmt.Println("Argument 1: ", os.Args[1])
}
