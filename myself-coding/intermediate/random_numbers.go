package intermediate

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create random number between 0 and 101.
	// fmt.Println(rand.Intn(100) + 1) // lower limit 1

	val := rand.New(rand.NewSource(time.Now().Unix()))

	// fmt.Println(rand.Intn(6) + 5)
	fmt.Println(val.Intn(101))
	// Between 0.0 and 1.0
	fmt.Println(rand.Float64())

	for {
		// Show the menu
		fmt.Println("Welcom to the Dice Game!")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Println("Enter your choice (1 or 2): ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invali choice, pls enter 1 or 2.")
			continue
		}
		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

		// Get max of 6 value numbers
		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		// Show the results
		fmt.Printf("You rolled a %d and a %d. \n", die1, die2)
		fmt.Println("Total:", die1+die2)

		// Ask for user wants to roll again
		fmt.Print("Do you to roll again? (y/n): ")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invali input, assuming no.")
		}
		if rollAgain == "n" {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}
	}
}
