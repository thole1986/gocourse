package main

import "fmt"

func main() {}

func adder() func() int {
	i := 0
	fmt.Println("Previous value of i: ", i)
	return func() int {
		i++
	}
}
