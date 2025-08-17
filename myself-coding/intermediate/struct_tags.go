package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	// Age       int    `json:"age,omitempty"`
	Age int `json:"-"` // Remove field in response by using "-"
}

func main() {
	// person := Person{FirstName: "Tho", LastName: "", Age: 50}
	// person := Person{FirstName: "Tho", Age: 0}
	person := Person{FirstName: "Tho", Age: 39}
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatalln("Error marshalling struct: ", err)
	}
	fmt.Println(string(jsonData))
}
