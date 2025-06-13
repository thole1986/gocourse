package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	fmt.Println("SOURCE: ", source)
	random := rand.New(source)

	// Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// fmt.Println("THE TARGET: ", target)

	// Welcome message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100")

	fmt.Println("Can you guess what it is?")

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess) // Store the address of variable

		// Check if the guess if correct.
		if guess == target {
			fmt.Println("Congralations! You guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a high number!")
		} else {
			fmt.Println("Too high! Try guessing a low number!")
		}
	}
}
