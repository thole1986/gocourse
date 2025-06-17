package basics

import (
	"fmt"
)

/*
******
	Notes:
	// Private function must be start with lowercase.
	// Publish function must be start with uppercase.
******
*/

func main() {
	// sum := add(1, 2)
	// fmt.Println(sum)

	// // This below is anonymous functions
	// greet := func() {
	// 	fmt.Println("Hello Anonymous Function")
	// }

	// greet() // greet is assigned to a function.

	// operation := add

	// result := operation(3, 5)

	// fmt.Println(result)

	// Passing a function as an argument.
	result := applyOperation(5, 3, add)
	fmt.Println("5 + 3 = ", result)

	// Returning and using a function.
	multiplyBy2 := createMultiplier(2) // -> Return a function
	fmt.Println("6 * 2 = ", multiplyBy2(6))
}

func add(a, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function that return a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
