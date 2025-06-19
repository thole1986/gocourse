package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/*
		String are immutable
	*/
	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!"
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	// fmt.Println("Length of message variable is", len(message))
	fmt.Println("Length of message2 variable is", len(message2))
	fmt.Println("Length of rawMessage variable is", len(rawMessage))

	fmt.Println("The first character in message var is ", message[0]) // ASCII

	greeting := "Hello "
	name := "Alice"

	// Concat 2 strings
	fmt.Println(greeting + name)
	str1 := "Apple"  // "A" has an ASCII value of 65
	str := "apple"   // "a" has an ASCII value of 97
	str2 := "banana" // "b" has an ASCII value of 98
	str3 := "app"    // "a" has an ASCII value of 97

	// Using lexicographical comparison
	// to compare strings
	fmt.Println(str1 < str2) // 65 < 98
	fmt.Println(str3 < str1) // 97 > 65
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	// String iteration.
	for _, char := range message {
		// fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%x", char)   // Get Heximal values
		fmt.Printf("%v\n", char) // Get the uint value
	}

	// Count in UTF8 characters
	// A rune is an integer value.
	fmt.Println("Rune count: ", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	// Declare rune with single quote ''
	var ch rune = 'a' // Unicode 32int and 4 bytes
	fmt.Println(ch)

	fmt.Printf("%c\n", ch)

	cstr := string(ch)
	fmt.Println(cstr)

	// %T to check any type is which type.
	fmt.Printf("Type of cstr is: (%T)\n", cstr)

	const NIHONGO = "日本語" // Japanese text
	fmt.Println(NIHONGO)

	jhello := "こんにちは" // Japanese "Hello"

	for _, runeValue := range jhello {
		fmt.Printf("This is unicode - integer format: %v\n", runeValue)
		fmt.Printf("This is actual value - string format: %c\n", runeValue)
	}

	r := '😊'
	// Unicode value 32int and 4 bytes
	fmt.Printf("%v\n", r)
	// Actual string value
	fmt.Printf("%c\n", r)
}
