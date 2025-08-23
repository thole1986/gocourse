package main

import "fmt"

func main() {
	var a int = 32
	b := int32(a)
	c := float64(b)

	fmt.Println(c)
}
