package intermediate

import "fmt"

func main() {
	num := 424
	fmt.Printf("%05d\n", num)

	message := "Hello fsfsfdsffsdfdsf"
	// Fix the width of output with
	// 10 leaing spaces characters on the right.
	fmt.Printf("|%10s|\n", message)
	fmt.Printf("|%-10s|\n", message)

	// String inter
	message1 := "Hello \nWorld"
	message2 := `Hello \nWorld` // Cannot add a new line
	fmt.Println(message1)
	fmt.Println(message2)

	sqlQuery := `SELECT * FROM users WHERE age > 30`
}
