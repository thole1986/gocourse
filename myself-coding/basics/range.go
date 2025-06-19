package basics

import "fmt"

func main() {
	message := "Hello World"

	for i, v := range message {
		// The value is the unicode value
		// fmt.Println(i, v)
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}
}
