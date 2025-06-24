package intermediate

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("He~lo, Base64 Encoding")

	// Encode Base64
	encoded := base64.StdEncoding.EncodeToString(data)

	fmt.Println(encoded)

	// Decode from Base64
	// Return with bytes slices or list
	decoded, err := base64.StdEncoding.DecodeString(encoded) // Return a byte slices
	if err != nil {
		fmt.Println("Error in decoding:", err)
		return
	}
	decoded_str := string(decoded)
	fmt.Println("Decoded: ", decoded_str)

	// URL safe, avoid '/' and '+'
	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)

	fmt.Println("URL Safe encoded: ", urlSafeEncoded)
}
