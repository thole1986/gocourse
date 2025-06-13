package basics

import "fmt"

func main() {
	// For as while with break
	// sum := 0
	// for {
	// 	sum += 10
	// 	fmt.Println("Iteration:", sum)
	// 	if sum >= 50 {
	// 		break
	// 	}
	// }

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Print("Odd number: ", num)
		num++ // Increase by 1
	}
}
