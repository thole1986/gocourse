package intermediate

import (
	"fmt"
	"strings"
)

func main() {
	// str := "Hello Go!"
	// fmt.Println(len(str))

	// str1 := "Hello"
	// str2 := "World"

	// result := str1 + " " + str2

	// fmt.Println(result)

	// fmt.Println(str[0])
	// fmt.Println(str[1:7])

	// // String conversion
	// num := 18
	// str3 := strconv.Itoa(num) // Convert number to string
	// fmt.Println(len(str3))

	// // strings splitting
	// fruits := "apple, orange, banana"
	// fruits1 := "apple-orange-banana"
	// parts := strings.Split(fruits, ",")
	// parts1 := strings.Split(fruits1, "-")
	// fmt.Println(parts)
	// fmt.Println(parts1)

	// countries := []string{"Gerpany", "France", "Italy"}
	// joined := strings.Join(countries, ", ")

	// fmt.Println(joined)

	// fmt.Println(strings.Contains(str, "Go"))
	// replaced := strings.Replace(str, "Go", "Universe", 1)
	// fmt.Println(replaced)

	// strspace := " Hello Everyone! "
	// fmt.Println(strspace)
	// fmt.Println(strings.TrimSpace(strspace))

	// fmt.Println(strings.ToLower(strspace))
	// fmt.Println(strings.ToUpper(strspace))

	// fmt.Println(strings.Repeat("foo ", 3))

	// fmt.Println(strings.Count("Hello", "l"))

	// fmt.Println(strings.HasPrefix("Hello", "he"))
	// fmt.Println(strings.HasSuffix("Hello", "la"))

	// // Advanced technics using regular patterns
	// str5 := "Hel1lo, 123 Go 11!"
	// re := regexp.MustCompile(`\d+`)
	// matched := re.FindAllString(str5, -1) // Find all string into paragraphs and return arrays
	// fmt.Println(matched)

	// str6 := "Hello, 世界"
	// fmt.Println(utf8.RuneCountInString(str6))

	// STRING BUILDER -> Optimize memorory
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	// Convert builder to a string
	result := builder.String() // Get the final strings to store in result
	fmt.Println(result)

	// Using Writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you!")
	result = builder.String()
	fmt.Println(result)

	// Reset the builder
	builder.Reset()

	// Start new sequences strings.
	builder.WriteString("Starting fresh!")
	result = builder.String()
	fmt.Println(result)
}
