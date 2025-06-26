package intermediate

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {
	// Create a new template with name example
	// tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")
	// if err != nil {
	// 	panic(err)
	// }

	// template.Must() => will handle error for us.
	// tmpl := template.Must(
	// 	template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n"),
	// )
	// Define data for the welcom message template.
	// data := map[string]interface{}{
	// 	"name": "John",
	// }

	// Extract the map name field to get the value
	// and output to the console.
	// err := tmpl.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	//
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	name, _ := reader.ReadString('\n') // this is bytes so use ''
	name = strings.TrimSpace(name)     // Clear whitespaces.

	// Define named tempates for different types of
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! We're glad you joined.",
		"notification": "{{.notification_name}}, you have a new notification: {{.ntf}}",
		"error":        "Oops! An error occurred: {{.em}}",
	}

	// Parse and store templates.
	parsedTemplates := make(map[string]*template.Template)

	// Creat multi templates by using for loop
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		// Show menu
		fmt.Println("\nMenu: ")
		fmt.Println("1. John")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option: ")

		// Convert to string when input into
		// console.
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		var data map[string]interface{}
		var tmpl *template.Template

		// Using swith to make more clear instead of using "if else if"
		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]interface{}{
				"name": name,
			}
		case "2":
			fmt.Println("Enter your notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{
				"notification_name": name,
				"ntf":               notification,
			}
		case "3":
			fmt.Println("Enter yr error message: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{
				"notification_name": name,
				"em":                errorMessage,
			}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
			continue
		}

		// Render and print the template to the console.
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing template: ", err)
		}
	}
}
