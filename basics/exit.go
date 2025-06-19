package basics

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("Deferred statement")

	fmt.Println("Starting the main function")

	// Exit with status code of 1
	// Using Exit function when need the programe stop
	// immediately when cannot control the errors
	os.Exit(1)

	// This will never be executed.
	fmt.Println("End of main function")
}
