package basics

import "fmt"

/* This will be run first automatically in Go */
func init() {
	fmt.Println("Initializing pacakage 1 ...")
}

func init() {
	fmt.Println("Initializing pacakage 2 ...")
}

func init() {
	fmt.Println("Initializing pacakage 3 ...")
}

func main() {
	fmt.Println("Inside the main function!")
}
