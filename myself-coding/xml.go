package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"` // The root element of XML struct
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	City    string   `xml:"city"`
	Email   string   `xml:"-"` // Ignore the field
	Address Address  `xml:"address"`
	// Email   string   `xml:"email"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {
	// person := Person{Name: "Tho", Age: 39, City: "HCMC", Email: "test@example.com"}
	person := Person{Name: "Tho", Email: "test@example.com", Address: Address{City: "Bien Hoa", State: "Dong Nai"}}
	// xmlData, err := xml.Marshal(person)
	// if err != nil {
	// 	log.Fatalln("Error marshalling data into XML: ", err)

	// }
	// fmt.Println(string(xmlData))

	// Convert struct Go data into XML data
	xmlData1, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalln("Error marshalling data into XML: ", err)

	}

	fmt.Println(string(xmlData1))

	// xmlRaw := `<person><name>Tho</name><age>39</age></person>`

	// Raw XML data
	xmlRaw := `<person><name>Tho</name><age>39</age><address><city>"Nha Trang"</city><state>Khanh Hoa</state></address></person>`
	var personxml Person
	// Convert XML data into struct Golang data
	err = xml.Unmarshal([]byte(xmlRaw), &personxml)

	if err != nil {
		log.Fatalln("Error Unmarshalling XML: ", err)
	}

	fmt.Println(personxml)
	fmt.Println("Local String: ", personxml.XMLName.Local)
	fmt.Println("Namespace", personxml.XMLName.Space)

	book := Book{
		ISBN:       "1239455654654",
		Title:      "Go Bootcamp",
		Author:     "Ashish",
		Pseudo:     "Pseudo",
		PseudoAttr: "Pseudo Attr",
	}
	xmlDataAttr, err := xml.MarshalIndent(book, "", " ")
	if err != nil {
		log.Fatalln("Error marshalling data: ", err)
	}

	fmt.Println(string(xmlDataAttr))
	// -><book isbn="1239455654654" title="Go Bootcamp" author="Ashish"></book>
}

type Book struct {
	XMLName    xml.Name `xml:"book"`
	ISBN       string   `xml:"isbn,attr"`
	Title      string   `xml:"title,attr"`
	Author     string   `xml:"author,attr"`
	Pseudo     string   `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudo,attr"`
}
