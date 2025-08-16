package intermediate

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User env var: ", user)
	fmt.Println("Home env var: ", home)

	err := os.Setenv("FRUIT", "APPLE")
	if err != nil {
		fmt.Println("Error setting enviroment variable: ", err)
	}
	fmt.Println("Enviroment variable FRUIT set successfully.")
	fmt.Println("FRUIT env var: ", os.Getenv("FRUIT"))

	// Unset env
	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error unsetting enviroment variable: ", err)
		return
	}
	fmt.Println("Unset env var done on key FRUIT")
	fmt.Println("FRUIT env var: ", os.Getenv("FRUIT"))

	fmt.Println("---------------------------------------------------")

	str := "a=b=c=d=e"

	/*Note: */
	// n = 1 return "a=b=c=d"
	// n = 2 return "a" and "b=c=d"
	// n = 3 return "a" and "b" and "c=d"
	// n = 4 return "a" and "b" and "c" and "d"

	fmt.Println(strings.SplitN(str, "=", -1)) // => [a b c d e]
	fmt.Println(strings.SplitN(str, "=", 0))  // => []
	fmt.Println(strings.SplitN(str, "=", 1))  // Split into 1 substring => [a=b=c=d=e]
	fmt.Println(strings.SplitN(str, "=", 2))  // Split into 2 substring => [a b=c=d=e]
	fmt.Println(strings.SplitN(str, "=", 3))  // => [a b c=d=e]
	fmt.Println(strings.SplitN(str, "=", 4))  // => [a b c d=e]
	fmt.Println(strings.SplitN(str, "=", 5))  // => [a b c d e]

	// for _, e := range os.Environ() {
	// 	kvpair := strings.SplitN(e, "=", 2)
	// 	fmt.Println(kvpair[0])
	// }
}
