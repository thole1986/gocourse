package intermediate

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	num, err := strconv.Atoi(numStr) // Convert string to int
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}
	fmt.Println("Parsed Integer: ", num+1)

	numistr, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}
	fmt.Println("Parsed Integer: ", numistr)

	floatstr := "3.14"
	floatval, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}

	fmt.Printf("Parse float: %.2f\n", floatval)

	binaryStr := "1010" // 0 + 2 + 0 + 8 = 10
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error parsing binary valye:", err)
		return
	}
	fmt.Println("Parsed binary to decimal: ", decimal)

	hexStr := "FF"
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error parsing binary value: ", err)
		return
	}
	fmt.Println("Parsed hex to decimal: ", hex)

	invalidnum := "456abc"
	invalidparse, err := strconv.Atoi(invalidnum)
	if err != nil {
		fmt.Println("Error parsing value: ", err)
		return
	}
	fmt.Println("Parsed invalid number: ", invalidparse)
}
