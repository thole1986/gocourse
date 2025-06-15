package advanced

import (
	"fmt"
	"reflect"
)

// ============== WORKING WITH METHODS
type Greeter struct{}

func (g Greeter) Greet(fname, lname string) string {
	return "Hello " + fname + " " + lname
}

func main() {
	g := Greeter{}
	t := reflect.TypeOf(g)
	v := reflect.ValueOf(g)
	var method reflect.Method

	fmt.Println("Type:", t)
	for i := range t.NumMethod() {
		method = t.Method(i)
		fmt.Printf("Method %d: %s\n", i, method.Name)
	}

	m := v.MethodByName(method.Name)
	results := m.Call([]reflect.Value{reflect.ValueOf("Alice"), reflect.ValueOf("Doe")})
	fmt.Println("Greet result:", results[0].String())
	// []string{"Alice"}
	// []type{type("hjdfghj"), type{"jdhfgkj"}}

}

// ======== WORKING WITH STRUCTS and FIELDS
// type person struct {
// 	Name string
// 	age  int
// }

// func main() {
// 	p := person{Name: "Alice", age: 30}
// 	v := reflect.ValueOf(p)

// 	for i := range v.NumField() {
// 		fmt.Printf("Field %d: %v\n", i, v.Field(i))
// 	}

// 	v1 := reflect.ValueOf(&p).Elem()

// 	nameField := v1.FieldByName("Name")
// 	if nameField.CanSet() {
// 		nameField.SetString("Jane")
// 	} else {
// 		fmt.Println("Cannot set")
// 	}

// 	fmt.Println("Modified Person:", p)
// }

// func main() {

// 	x := 42
// 	v := reflect.ValueOf(x)
// 	t := v.Type()

// 	fmt.Println("Value:", v)
// 	fmt.Println("Type:", t)
// 	fmt.Println("Kind:", t.Kind())
// 	fmt.Println("Is Int:", t.Kind() == reflect.Int)
// 	fmt.Println("Is String:", t.Kind() == reflect.String)
// 	fmt.Println("Is Zero:", v.IsZero())

// 	y := 10
// 	v1 := reflect.ValueOf(&y).Elem()
// 	v2 := reflect.ValueOf(&y)
// 	fmt.Println("V2 Type:", v2.Type())

// 	fmt.Println("Original value:", v1.Int())

// 	v1.SetInt(18)
// 	fmt.Println("Modified value:", v1.Int())

// 	var itf interface{} = "Hello"
// 	v3 := reflect.ValueOf(itf)

// 	fmt.Println("V3 Type:", v3.Type())
// 	if v3.Kind() == reflect.String {
// 		fmt.Println("String value:", v3.String())
// 	}

// }
