package basics

import "fmt"

func main() {
	// var arrayName [size]elementType

	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}

	fmt.Println(fruits)

	fmt.Println("Third element: ", fruits[2])

	// originalArray := [3]int{1, 2, 3}
	// coppiedArray := originalArray
	// coppiedArray[0] = 100

	// fmt.Println("Original array: ", originalArray)
	// // Copy a new array -> update not change orginal arrays
	// fmt.Println("Coppied array: ", coppiedArray)

	// Use normal way to loop.
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index: ", i, ":", numbers[i])
	}

	// Use short way.
	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// Loop and ignore index.
	// underscore is blank identifier, used to store unused values.
	for _, v := range numbers {
		fmt.Printf("Value: %d\n", v)
	}

	// Get only one value from return function
	// Not got second -> use "_" to ignore.
	a, _ := someFunction()
	fmt.Println(a)
	// fmt.Println(b)
	b := 2
	_ = b

	fmt.Println("The length of numbers array is", len(numbers))

	// Compare arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Println("Array1 is equal to Array2: ", array1 == array2)

	// Multi dimension arrays.
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	originalArray := [3]int{1, 2, 3}
	// Create a pointer
	var coppiedArray *[3]int
	// A pointer to memory address in orginal arrays.
	coppiedArray = &originalArray
	coppiedArray[0] = 100

	fmt.Println("Original array: ", originalArray)
	// Copy a new array -> update not change orginal arrays
	fmt.Println("Coppied array: ", coppiedArray)
}

func someFunction() (int, int) {
	return 1, 2
}
