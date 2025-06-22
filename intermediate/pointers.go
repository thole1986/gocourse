package intermediate

import "fmt"

func main() {

	var ptr *int
	var a int = 10

	// Pointer to memory addr of a
	ptr = &a // referencing

	fmt.Println(a)
	fmt.Println(ptr)
	// Get memory addr of viriable
	// fmt.Println(ptr)

	// Using * to get actual value
	// fmt.Println(*ptr) // dereferencing a point

	// if ptr == nil {
	// 	fmt.Println("Pointer is nill")
	// }

	modifyValue(ptr)
	fmt.Println(a)
}

// Add param to function accept a pointer
func modifyValue(ptr *int) {
	// Dereferencing ptr or get actual value
	// of ptr param
	*ptr++
}
