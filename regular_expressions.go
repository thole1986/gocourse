package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("He said, \n\"I am great\"")
	fmt.Println(`He said, "I am great"`)

	// Compile a regex pattern to match email address.
	re := regexp.MustCompile()
}
