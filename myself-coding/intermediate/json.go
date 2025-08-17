package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
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
			"state": "Q1"
		}
	}`

	var employeeFromJson Employee

	err = json.Unmarshal([]byte(jsonData2), &employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON", err)
		return
	}

	fmt.Println(employeeFromJson)
	fmt.Println("Tho Le's Age increased by 5 years", employeeFromJson.Age+5)
	fmt.Println(employeeFromJson.Address.City)

	listOfCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "San Yose", State: "CA"},
		{City: "San Yose", State: "CA"},
		{City: "San Yose", State: "CA"},
		{City: "San Yose", State: "CA"},
	}
	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		log.Fatalln("Error Marshalling to JSON:", err)
	}
	fmt.Println("JSON List: ", string(jsonList))

	// Handling unknow json structures.
	jsonData3 := `
		{"name": "John", "age": 30, "address": {
			"city": "New York",
			"state": "NY"
		}}
	`
	var data map[string]interface{}

	err = json.Unmarshal([]byte(jsonData3), &data)

	if err != nil {
		log.Fatalln("Error Unmarshalled JSON: ", data)
	}

	fmt.Println("Decoded/Unmarshalled JSON: ", data)
	fmt.Println("Decoded/Unmarshalled JSON: ", data["address"])
	fmt.Println("Decoded/Unmarshalled JSON: ", data["name"])
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
