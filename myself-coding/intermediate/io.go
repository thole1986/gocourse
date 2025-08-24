package intermediate

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)

	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}

	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}
}

func closeResource(c io.Closer) {
	err := c.Close() // Inferface
	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer // In stack
	buf.WriteString("Hello Buffer!")

	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello")
	r2 := strings.NewReader("World!")

	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer) // in heap

	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}

	fmt.Println(buf.String())

}

func pipeExample() {
	pr, pw := io.Pipe()

	go func() {
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func writeToFile(filepath string, data string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening/creating file.", err)
	}

	defer closeResource(file)

	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error opening/creating file: ", err)
	}

	// Type(value)
	// writer := io.Writer(file)           // Writer interface
	// _, err = writer.Write([]byte(data)) // Write method of above interaface
	// if err != nil {
	// 	log.Fatalln("Error opening/creating file.", err)
	// }
}

func main() {
	fmt.Println("=== read from Reader")
	readFromReader(strings.NewReader("Hello Reader!"))

	fmt.Println("=== write from Writer")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello Writer")
	fmt.Println(writer.String()) // Convert to string

	fmt.Println("=== Buffer Example === ")
	bufferExample()

	fmt.Println("=== Multi Reader Example === ")
	multiReaderExample()

	fmt.Println("=== Pipe Example === ")
	pipeExample()

	filepath := "io.txt"
	writeToFile(filepath, "Hello File!")

	resource := &MyResource{name: "TestResource"}
	closeResource(resource)
}

type MyResource struct {
	name string
}

func (m MyResource) Close() error {
	fmt.Println("Closing resource: ", m.name)
	return nil
}
