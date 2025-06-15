package intermediate

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	Email   string   `xml:"-"`
	Address Address  `xml:"address"`
	// Email   string   `xml:"email"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {

	// person := Person{Name: "John", Age: 30, City: "London", Email: "email@exampleemail.com"}
	person := Person{Name: "John", Email: "email@exampleemail.com", Address: Address{City: "Oakland", State: "CA"}}

	// xmlData, err := xml.Marshal(person)
	// if err != nil {
	// 	log.Fatalln("Error Marshalling data into XML:", err)
	// }
	// fmt.Println(string(xmlData))

	xmlData1, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalln("Error Marshalling data into XML:", err)
	}
	fmt.Println(string(xmlData1))

	// xmlRaw := `<person><name>Jane</name><age>25</age></person>`
	xmlRaw := `<person><name>John</name><age>25</age><address><city>San Francisco</city><state>CA</state></address></person>`

	var personxml Person

	err = xml.Unmarshal([]byte(xmlRaw), &personxml)
	if err != nil {
		log.Fatalln("Error Unmarshalling XML:", err)
	}

	fmt.Println(personxml)
	fmt.Println("Local String", personxml.XMLName.Local)
	fmt.Println("Namespace", personxml.XMLName.Space)

	book := Book{
		ISBN:       "584-75-437-587-34-32",
		Title:      "Go Bootcamp",
		Author:     "Ashish",
		Pseudo:     "Pseudo",
		PseudoAttr: "Pseudo Attribute",
	}

	xmlDataAttr, err := xml.MarshalIndent(book, "a", "  ")
	if err != nil {
		log.Fatalln("Error Marshalling data:", err)
	}

	fmt.Println(string(xmlDataAttr))

}

type Book struct {
	XMLName    xml.Name `xml:"book"`
	ISBN       string   `xml:"isbn,attr"`
	Title      string   `xml:"title,attr"`
	Author     string   `xml:"author,attr"`
	Pseudo     string   `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudoattr,attr"`
}

// <book isbn="jhghdf889e7yr89" color="blue">
