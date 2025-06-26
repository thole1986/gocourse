package mai

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file.", file)
	}

	defer file.Close()

	// Write data to file.
	data := []byte("Hello World!\n\n\n")
	_, err = file.Write(data)

	if err != nil {
		fmt.Println("Error Writing to file: ", err)
		return
	}

	fmt.Println("Data has been written to file successfully!")

	file, err = os.Create("writeString.txt")

	if err != nil {
		fmt.Println("Error Writing to file: ", err)
		return
	}

	defer file.Close()
	file.WriteString("Hello Go!\n\n\n")

	if err != nil {
		fmt.Println("Error wrinting to file:", err)
		return
	}

	fmt.Println("Writting to writeString.txt successfully!")
}
