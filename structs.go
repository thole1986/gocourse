package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

func main() {
	// type Person struct {
	// 	firstName string
	// 	lastName  string
	// 	age       int
	// }

	p1 := Person{
		firstName: "Tho",
		lastName:  "Le",
		age:       30,
		address: Address{
			city:    "London",
			country: "U.K",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "123456",
			cell: "55441122",
		},
	}

	p2 := Person{
		firstName: "Jane",
		age:       25,
	}

	p3 := Person{
		firstName: "Jane",
		age:       25,
	}

	// p2.address.city = "New York"
	// p2.address.country = "USA"

	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println(p1.fullName())
	fmt.Println(p1.address)
	fmt.Println(p2.address.country)
	fmt.Println(p1.cell)
	fmt.Println(p1.address.city)
	fmt.Println("Are p1 and p2 equal:", p1 == p2)
	fmt.Println("Are p2 and p3 equal:", p2 == p3)
	// Anonymous Struts
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "tholhlearning@gmail.com",
	}

	fmt.Println(user.username)
	fmt.Println("Before increment", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After increment", p1.age)
}
