package intermediate

import "fmt"

func main() {

	num := 424
	fmt.Printf("%05d\n", num)

	message := "Hello"
	fmt.Printf("|%10s|\n", message)
	fmt.Printf("|%-10s|\n", message)

	message1 := "Hello \nWorld!"
	message2 := `Hello \nWorld!`

	fmt.Println(message1)
	fmt.Println(message2)

	// sqlQuery := `SELECT * FROM users WHERE age > 30`

}
