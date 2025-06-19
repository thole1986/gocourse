package basics

import "fmt"

func main() {
	process()
	fmt.Println("Returned from Process")
}

func process() {
	defer func() {
		// recover() works like catch any result from panic func
		// and return. It not prevent the process

		// if r := recover(); r != nil {
		r := recover()
		if r != nil {
			// recover() catch the error from
			// panic func and return in final
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong!")
	// fmt.Println("End process!")
}
