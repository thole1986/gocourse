package main

import "fmt"

func main() {

	message := "Hello World"

	for i, v := range message {
		// fmt.Println(i, v)
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}

}
