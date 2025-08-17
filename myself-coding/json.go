package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName    string  `json:"first_name"` // return json field first_name
	Age          int     `json:"age,omitempty"`
	EmailAddress string  `json:"email,omitempty"`
	Address      Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	// person := Person{FirstName: "Tho", Age: 39}

	// Optinal use omitempty
	person := Person{FirstName: "Tho"}

	// Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON: ", err)
		return
	}

	fmt.Println(string(jsonData))

	person1 := Person{
		FirstName:    "Khoi",
		Age:          14,
		EmailAddress: "tholh@gmail.com",
		Address: Address{
			City:  "HCMC",
			State: "Hoc Mon",
		},
	}

	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON: ", err)
		return
	}

	fmt.Println(string(jsonData1))

	jsonData2 := `{
		"full_name": "Tho Le",
		"emp_id": "0009",
		"age": 39,
		"address": {
			"city": "HCMC",
			"state": "Q1",
		}
	}`

	var employeeFromJson Employee

	json.Unmarshal([]byte(jsonData2), employeeFromJson)
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      string  `json:"age"`
	Address  Address `json:"address"`
}
