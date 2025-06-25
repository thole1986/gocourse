package intermediate

import (
	// "crypto/internal/fips140/sha256"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password123"

	// Return a list slices of bytes
	// hash := sha256.Sum256([]byte(password))
	// hash512 := sha512.Sum512([]byte(password))
	// fmt.Println(password)
	// fmt.Println(hash)
	// fmt.Println(hash512)

	// print SHA-256 characters in strings
	// fmt.Printf("SHA-256 Hash hex val: %x\n", hash)

	// print SHA-512 characters in strings
	// fmt.Printf("SHA-512 Hash hex val: %x\n", hash512)

	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt: ", err)
		return
	}

	fmt.Println("GENERATED SALTED: ", salt)

	// Hash the password with salt
	signUpHash := hashPassword("password124", salt)
	fmt.Println("HASH PASSWORD: ", signUpHash)

	// Store the salt and password in database, just println
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt String FROM byte of slices: ", saltStr)

	hashOriginalPassword := sha256.Sum256([]byte(password))

	fmt.Println("Hash of just the password string without salt:", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))

	// Verify password
	// Retrieve the saltStr and decode it.
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt:", err)
		return
	}

	loginHash := hashPassword(password, decodedSalt)

	// Compare the stored signUpHash with loginHash
	if signUpHash == loginHash {
		fmt.Println(("Password is correct. You are logged in."))
	} else {
		fmt.Println("Login failed. Please check user credentials.")
	}
}

func generateSalt() ([]byte, error) {
	// Make a bytes slice with 16 bytes
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Function to hash password
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...) // Destructor slices of password
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
