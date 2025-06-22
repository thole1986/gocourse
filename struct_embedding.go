package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) introduce() {
	fmt.Printf("Hi, I am %s and I am %d years old.\n", p.name, p.age)
}

type Employee struct {
	employeeInfo person // embedded struct Named field
	// person              // Anonymous field
	empId  string
	salary float64
}

// This overwrite methods with the same name
func (e Employee) introduce() {
	fmt.Printf("Hi, I am %s, employee ID: %s and I earn %.2f\n", e.employeeInfo.name, e.empId, e.salary)
}

func main() {
	emp := Employee{
		employeeInfo: person{name: "John", age: 30},
		empId:        "E001", salary: 50000,
	}

	fmt.Println("Name", emp.employeeInfo.name) // Accessing the embedded struct field emp.person.name
	fmt.Println("Age", emp.employeeInfo.age)   // Same as above
	fmt.Println("Emp ID:", emp.empId)
	fmt.Println("Salary", emp.salary)

	emp.introduce() // Call the method of instance person
}
