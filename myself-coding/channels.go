package main

import "fmt"

func main() {

	greeting := make(chan string)
	// greetingString := "Hello"

	go func() {
		// blocking because it is continously trying to receive values,
		// it is ready to receive continuous flow of data
		receiver := <-greeting
		fmt.Println(receiver)
	}()

	receiver := <-greeting
	fmt.Println(receiver)
	fmt.Println("End of program.")
}
