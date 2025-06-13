package basics

import "fmt"

var middleName = "Cane"

func main() {
	// var age int
	// var name string = "John"
	// var name1 = "Jane"

	// count := 10
	// lastName := "Smith"
	middleName := "Mayor"

	fmt.Println(middleName)

	// Default values
	// Numeric Types: 0
	// Boolean Types: False
	// String Type: ""
	// Pointers, slices, maps, functions, and structs: nil

	// ---- Scope.
	// fmt.Println(firstName)
}

func printName() {
	firstName := "Michale"
	fmt.Println(firstName)
}
