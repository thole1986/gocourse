package main

import "fmt"

func main() {
	// Switch statement in golang -> switch case default
	// switch expression {
	// case value1:
	// // Code to be excuted if expression equal value1
	// case value2:
	// 	// Code to be excuted if expression equal value1
	// default:
	// 	// Code executed if does not match any value.
	// }

	// fruit := "apple"
	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an apple.")
	// case "banana":
	// 	fmt.Println("It's a banana.")
	// default:
	// 	fmt.Println("Unknow Fruit!")
	// }

	// day := "Monday"
	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday.")
	// case "Sunday":
	// 	fmt.Println("It's a weekend.")
	// default:
	// 	fmt.Println("Invalid day.")
	// }

	// number := 15
	// switch {
	// case number < 10:
	// 	fmt.Println("Number is less than 10.")
	// case number >= 10 && number < 20:
	// 	fmt.Println("Number is between 10 and 19.")
	// default:
	// 	fmt.Println("Number is 20 or more.")
	// }

	// num := 2
	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough // This use will check for the next case below
	// case num == 2:
	// 	fmt.Println("Number is 2")
	// default:
	// 	fmt.Println("Not two")
	// }

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

// x is any data types
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer.")
	case float64:
		fmt.Println("It's an float.")
	case string:
		fmt.Println("It's an string.")
	default:
		fmt.Println("Unknow Type!")
	}
}
