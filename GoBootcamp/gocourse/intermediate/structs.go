package intermediate

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

	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "London",
			country: "U.K.",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "465456454",
			cell: "45456464544",
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
	fmt.Println("Are p3 and p2 equal:", p3 == p2)

	// Anonymous Structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "pseudoemail@example.org",
	}

	fmt.Println(user.username)
	fmt.Println("Before increment", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After increment", p1.age)

}
